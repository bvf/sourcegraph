package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	log15 "gopkg.in/inconshreveable/log15.v2"

	opentracing "github.com/opentracing/opentracing-go"
	otlog "github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"sourcegraph.com/sourcegraph/sourcegraph/cmd/frontend/internal/app/envvar"
	"sourcegraph.com/sourcegraph/sourcegraph/cmd/frontend/internal/db"
	"sourcegraph.com/sourcegraph/sourcegraph/cmd/frontend/internal/pkg/types"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/api"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/conf"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/env"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/gitserver"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/inventory"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/rcache"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/repoupdater"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/repoupdater/protocol"
)

// ErrRepoSeeOther indicates that the repo does not exist on this server but might exist on an external Sourcegraph
// server.
type ErrRepoSeeOther struct {
	// RedirectURL is the base URL for the repository at an external location.
	RedirectURL string
}

func (e ErrRepoSeeOther) Error() string {
	return fmt.Sprintf("repo not found at this location, but might exist at %s", e.RedirectURL)
}

var Repos = &repos{}

type repos struct{}

func (s *repos) Get(ctx context.Context, repo api.RepoID) (_ *types.Repo, err error) {
	if Mocks.Repos.Get != nil {
		return Mocks.Repos.Get(ctx, repo)
	}

	ctx, done := trace(ctx, "Repos", "Get", repo, &err)
	defer done()

	return db.Repos.Get(ctx, repo)
}

// GetByURI retrieves the repository with the given URI. If the URI refers to a repository on a known external
// service (such as a code host) that is not yet present in the database, it will automatically look up the
// repository externally and add it to the database before returning it.
func (s *repos) GetByURI(ctx context.Context, uri api.RepoURI) (_ *types.Repo, err error) {
	if Mocks.Repos.GetByURI != nil {
		return Mocks.Repos.GetByURI(ctx, uri)
	}

	ctx, done := trace(ctx, "Repos", "GetByURI", uri, &err)
	defer done()

	repo, err := db.Repos.GetByURI(ctx, uri)
	if err != nil && envvar.SourcegraphDotComMode() {
		// Automatically add repositories on Sourcegraph.com.
		if err := s.Add(ctx, uri); err != nil {
			return nil, err
		}
		return db.Repos.GetByURI(ctx, uri)
	} else if err != nil {
		if !conf.GetTODO().DisablePublicRepoRedirects && strings.HasPrefix(strings.ToLower(string(uri)), "github.com/") {
			return nil, ErrRepoSeeOther{RedirectURL: fmt.Sprintf("https://sourcegraph.com/%s", uri)}
		}
		return nil, err
	}

	return repo, nil
}

// Add adds the repository with the given URI. The URI is mapped to a repository by consulting the
// repo-updater, which contains information about all configured code hosts and the URIs that they
// handle.
func (s *repos) Add(ctx context.Context, uri api.RepoURI) (err error) {
	if Mocks.Repos.Add != nil {
		return Mocks.Repos.Add(uri)
	}

	ctx, done := trace(ctx, "Repos", "Add", uri, &err)
	defer done()

	// Avoid hitting the repoupdater (and incurring a hit against our GitHub/etc. API rate
	// limit) for repositories that don't exist or private repositories that people attempt to
	// access.
	if gitserverRepo := quickGitserverRepoInfo(uri); gitserverRepo != nil {
		if err := gitserver.DefaultClient.IsRepoCloneable(ctx, *gitserverRepo); err != nil {
			return err
		}
	}

	// Try to look up and add the repo.
	result, err := repoupdater.DefaultClient.RepoLookup(ctx, protocol.RepoLookupArgs{Repo: uri})
	if err != nil {
		return err
	}
	if result.Repo != nil {
		// Allow anonymous users on Sourcegraph.com to enable repositories just by visiting them, but
		// everywhere else, require server admins to explicitly enable repositories.
		enableAutoAddedRepos := envvar.SourcegraphDotComMode()
		if err := s.TryInsertNew(ctx, api.InsertRepoOp{
			URI:          result.Repo.URI,
			Description:  result.Repo.Description,
			Fork:         result.Repo.Fork,
			Enabled:      enableAutoAddedRepos,
			ExternalRepo: result.Repo.ExternalRepo,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *repos) TryInsertNew(ctx context.Context, op api.InsertRepoOp) error {
	return db.Repos.TryInsertNew(ctx, op)
}

func (s *repos) TryInsertNewBatch(ctx context.Context, repos []api.InsertRepoOp) error {
	return db.Repos.TryInsertNewBatch(ctx, repos)
}

func (s *repos) List(ctx context.Context, opt db.ReposListOptions) (repos []*types.Repo, err error) {
	if Mocks.Repos.List != nil {
		return Mocks.Repos.List(ctx, opt)
	}

	ctx, done := trace(ctx, "Repos", "List", opt, &err)
	defer func() {
		if err == nil {
			span := opentracing.SpanFromContext(ctx)
			span.LogFields(otlog.Int("result.len", len(repos)))
		}
		done()
	}()

	return db.Repos.List(ctx, opt)
}

var inventoryCache = rcache.New("inv")

func (s *repos) GetInventory(ctx context.Context, repo *types.Repo, commitID api.CommitID) (res *inventory.Inventory, err error) {
	if Mocks.Repos.GetInventory != nil {
		return Mocks.Repos.GetInventory(ctx, repo, commitID)
	}

	ctx, done := trace(ctx, "Repos", "GetInventory", map[string]interface{}{"repo": repo.URI, "commitID": commitID}, &err)
	defer done()

	// Cap GetInventory operation to some reasonable time.
	ctx, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()

	if !isAbsCommitID(commitID) {
		return nil, errors.Errorf("non-absolute CommitID for Repos.GetInventory: %v", commitID)
	}

	// Try cache first
	cacheKey := fmt.Sprintf("%s:%s", repo.URI, commitID)
	if b, ok := inventoryCache.Get(cacheKey); ok {
		var inv inventory.Inventory
		if err := json.Unmarshal(b, &inv); err == nil {
			return &inv, nil
		}
		log15.Warn("Repos.GetInventory failed to unmarshal cached JSON inventory", "repo", repo.URI, "commitID", commitID, "err", err)
	}

	// Not found in the cache, so compute it.
	inv, err := s.GetInventoryUncached(ctx, repo, commitID)
	if err != nil {
		return nil, err
	}

	// Store inventory in cache.
	b, err := json.Marshal(inv)
	if err != nil {
		return nil, err
	}
	inventoryCache.Set(cacheKey, b)

	return inv, nil
}

func (s *repos) GetInventoryUncached(ctx context.Context, repo *types.Repo, commitID api.CommitID) (*inventory.Inventory, error) {
	if Mocks.Repos.GetInventoryUncached != nil {
		return Mocks.Repos.GetInventoryUncached(ctx, repo, commitID)
	}

	vcsrepo := Repos.CachedVCS(repo)
	files, err := vcsrepo.ReadDir(ctx, commitID, "", true)
	if err != nil {
		return nil, err
	}
	return inventory.Get(ctx, files)
}

var indexerAddr = env.Get("SRC_INDEXER", "indexer:3179", "The address of the indexer service.")

func (s *repos) RefreshIndex(ctx context.Context, repo *types.Repo) (err error) {
	if Mocks.Repos.RefreshIndex != nil {
		return Mocks.Repos.RefreshIndex(ctx, repo)
	}

	if !repo.Enabled {
		return nil
	}

	go func() {
		resp, err := http.Get("http://" + indexerAddr + "/refresh?repo=" + string(repo.URI))
		if err != nil {
			log15.Error("RefreshIndex failed", "error", err)
			return
		}
		resp.Body.Close()
	}()

	return nil
}
