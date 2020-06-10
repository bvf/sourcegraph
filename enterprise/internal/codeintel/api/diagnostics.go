package api

import (
	"context"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/bundles/client"
	bundles "github.com/sourcegraph/sourcegraph/internal/codeintel/bundles/client"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/db"
)

type ResolvedDiagnostic struct {
	Dump       db.Dump
	Diagnostic bundles.Diagnostic
}

// Diagnostics returns the diagnostics for documents with the given path prefix.
func (api *codeIntelAPI) Diagnostics(ctx context.Context, prefix string, uploadID, limit, offset int) ([]ResolvedDiagnostic, int, error) {
	dump, exists, err := api.db.GetDumpByID(ctx, uploadID)
	if err != nil {
		return nil, 0, errors.Wrap(err, "db.GetDumpByID")
	}
	if !exists {
		return nil, 0, ErrMissingDump
	}

	pathInBundle := strings.TrimPrefix(prefix, dump.Root)
	bundleClient := api.bundleManagerClient.BundleClient(dump.ID)

	diagnostics, totalCount, err := bundleClient.Diagnostics(ctx, pathInBundle, offset, limit)
	if err != nil {
		if err == client.ErrNotFound {
			log15.Warn("Bundle does not exist")
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "bundleClient.Diagnostics")
	}

	return resolveDiagnosticsWithDump(dump, diagnostics), totalCount, nil
}

func resolveDiagnosticsWithDump(dump db.Dump, diagnostics []bundles.Diagnostic) []ResolvedDiagnostic {
	var resolvedDiagnostics []ResolvedDiagnostic
	for _, diagnostic := range diagnostics {
		diagnostic.Path = dump.Root + diagnostic.Path
		resolvedDiagnostics = append(resolvedDiagnostics, ResolvedDiagnostic{
			Dump:       dump,
			Diagnostic: diagnostic,
		})
	}

	return resolvedDiagnostics
}