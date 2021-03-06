// Code generated by github.com/efritz/go-mockgen 0.1.0; DO NOT EDIT.

package scheduler

import (
	"context"
	store "github.com/sourcegraph/sourcegraph/enterprise/internal/codeintel/store"
	"regexp"
	"sync"
)

// MockGitserverClient is a mock implementation of the gitserverClient
// interface (from the package
// github.com/sourcegraph/sourcegraph/enterprise/cmd/precise-code-intel-indexer/internal/scheduler)
// used for unit testing.
type MockGitserverClient struct {
	// FileExistsFunc is an instance of a mock function object controlling
	// the behavior of the method FileExists.
	FileExistsFunc *GitserverClientFileExistsFunc
	// HeadFunc is an instance of a mock function object controlling the
	// behavior of the method Head.
	HeadFunc *GitserverClientHeadFunc
	// ListFilesFunc is an instance of a mock function object controlling
	// the behavior of the method ListFiles.
	ListFilesFunc *GitserverClientListFilesFunc
	// RawContentsFunc is an instance of a mock function object controlling
	// the behavior of the method RawContents.
	RawContentsFunc *GitserverClientRawContentsFunc
}

// NewMockGitserverClient creates a new mock of the gitserverClient
// interface. All methods return zero values for all results, unless
// overwritten.
func NewMockGitserverClient() *MockGitserverClient {
	return &MockGitserverClient{
		FileExistsFunc: &GitserverClientFileExistsFunc{
			defaultHook: func(context.Context, store.Store, int, string, string) (bool, error) {
				return false, nil
			},
		},
		HeadFunc: &GitserverClientHeadFunc{
			defaultHook: func(context.Context, store.Store, int) (string, error) {
				return "", nil
			},
		},
		ListFilesFunc: &GitserverClientListFilesFunc{
			defaultHook: func(context.Context, store.Store, int, string, *regexp.Regexp) ([]string, error) {
				return nil, nil
			},
		},
		RawContentsFunc: &GitserverClientRawContentsFunc{
			defaultHook: func(context.Context, store.Store, int, string, string) ([]byte, error) {
				return nil, nil
			},
		},
	}
}

// surrogateMockGitserverClient is a copy of the gitserverClient interface
// (from the package
// github.com/sourcegraph/sourcegraph/enterprise/cmd/precise-code-intel-indexer/internal/scheduler).
// It is redefined here as it is unexported in the source packge.
type surrogateMockGitserverClient interface {
	FileExists(context.Context, store.Store, int, string, string) (bool, error)
	Head(context.Context, store.Store, int) (string, error)
	ListFiles(context.Context, store.Store, int, string, *regexp.Regexp) ([]string, error)
	RawContents(context.Context, store.Store, int, string, string) ([]byte, error)
}

// NewMockGitserverClientFrom creates a new mock of the MockGitserverClient
// interface. All methods delegate to the given implementation, unless
// overwritten.
func NewMockGitserverClientFrom(i surrogateMockGitserverClient) *MockGitserverClient {
	return &MockGitserverClient{
		FileExistsFunc: &GitserverClientFileExistsFunc{
			defaultHook: i.FileExists,
		},
		HeadFunc: &GitserverClientHeadFunc{
			defaultHook: i.Head,
		},
		ListFilesFunc: &GitserverClientListFilesFunc{
			defaultHook: i.ListFiles,
		},
		RawContentsFunc: &GitserverClientRawContentsFunc{
			defaultHook: i.RawContents,
		},
	}
}

// GitserverClientFileExistsFunc describes the behavior when the FileExists
// method of the parent MockGitserverClient instance is invoked.
type GitserverClientFileExistsFunc struct {
	defaultHook func(context.Context, store.Store, int, string, string) (bool, error)
	hooks       []func(context.Context, store.Store, int, string, string) (bool, error)
	history     []GitserverClientFileExistsFuncCall
	mutex       sync.Mutex
}

// FileExists delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockGitserverClient) FileExists(v0 context.Context, v1 store.Store, v2 int, v3 string, v4 string) (bool, error) {
	r0, r1 := m.FileExistsFunc.nextHook()(v0, v1, v2, v3, v4)
	m.FileExistsFunc.appendCall(GitserverClientFileExistsFuncCall{v0, v1, v2, v3, v4, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the FileExists method of
// the parent MockGitserverClient instance is invoked and the hook queue is
// empty.
func (f *GitserverClientFileExistsFunc) SetDefaultHook(hook func(context.Context, store.Store, int, string, string) (bool, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// FileExists method of the parent MockGitserverClient instance inovkes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *GitserverClientFileExistsFunc) PushHook(hook func(context.Context, store.Store, int, string, string) (bool, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *GitserverClientFileExistsFunc) SetDefaultReturn(r0 bool, r1 error) {
	f.SetDefaultHook(func(context.Context, store.Store, int, string, string) (bool, error) {
		return r0, r1
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *GitserverClientFileExistsFunc) PushReturn(r0 bool, r1 error) {
	f.PushHook(func(context.Context, store.Store, int, string, string) (bool, error) {
		return r0, r1
	})
}

func (f *GitserverClientFileExistsFunc) nextHook() func(context.Context, store.Store, int, string, string) (bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitserverClientFileExistsFunc) appendCall(r0 GitserverClientFileExistsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitserverClientFileExistsFuncCall objects
// describing the invocations of this function.
func (f *GitserverClientFileExistsFunc) History() []GitserverClientFileExistsFuncCall {
	f.mutex.Lock()
	history := make([]GitserverClientFileExistsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitserverClientFileExistsFuncCall is an object that describes an
// invocation of method FileExists on an instance of MockGitserverClient.
type GitserverClientFileExistsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 store.Store
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 int
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 string
	// Arg4 is the value of the 5th argument passed to this method
	// invocation.
	Arg4 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 bool
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitserverClientFileExistsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3, c.Arg4}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitserverClientFileExistsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// GitserverClientHeadFunc describes the behavior when the Head method of
// the parent MockGitserverClient instance is invoked.
type GitserverClientHeadFunc struct {
	defaultHook func(context.Context, store.Store, int) (string, error)
	hooks       []func(context.Context, store.Store, int) (string, error)
	history     []GitserverClientHeadFuncCall
	mutex       sync.Mutex
}

// Head delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockGitserverClient) Head(v0 context.Context, v1 store.Store, v2 int) (string, error) {
	r0, r1 := m.HeadFunc.nextHook()(v0, v1, v2)
	m.HeadFunc.appendCall(GitserverClientHeadFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Head method of the
// parent MockGitserverClient instance is invoked and the hook queue is
// empty.
func (f *GitserverClientHeadFunc) SetDefaultHook(hook func(context.Context, store.Store, int) (string, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Head method of the parent MockGitserverClient instance inovkes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *GitserverClientHeadFunc) PushHook(hook func(context.Context, store.Store, int) (string, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *GitserverClientHeadFunc) SetDefaultReturn(r0 string, r1 error) {
	f.SetDefaultHook(func(context.Context, store.Store, int) (string, error) {
		return r0, r1
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *GitserverClientHeadFunc) PushReturn(r0 string, r1 error) {
	f.PushHook(func(context.Context, store.Store, int) (string, error) {
		return r0, r1
	})
}

func (f *GitserverClientHeadFunc) nextHook() func(context.Context, store.Store, int) (string, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitserverClientHeadFunc) appendCall(r0 GitserverClientHeadFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitserverClientHeadFuncCall objects
// describing the invocations of this function.
func (f *GitserverClientHeadFunc) History() []GitserverClientHeadFuncCall {
	f.mutex.Lock()
	history := make([]GitserverClientHeadFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitserverClientHeadFuncCall is an object that describes an invocation of
// method Head on an instance of MockGitserverClient.
type GitserverClientHeadFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 store.Store
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 string
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitserverClientHeadFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitserverClientHeadFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// GitserverClientListFilesFunc describes the behavior when the ListFiles
// method of the parent MockGitserverClient instance is invoked.
type GitserverClientListFilesFunc struct {
	defaultHook func(context.Context, store.Store, int, string, *regexp.Regexp) ([]string, error)
	hooks       []func(context.Context, store.Store, int, string, *regexp.Regexp) ([]string, error)
	history     []GitserverClientListFilesFuncCall
	mutex       sync.Mutex
}

// ListFiles delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockGitserverClient) ListFiles(v0 context.Context, v1 store.Store, v2 int, v3 string, v4 *regexp.Regexp) ([]string, error) {
	r0, r1 := m.ListFilesFunc.nextHook()(v0, v1, v2, v3, v4)
	m.ListFilesFunc.appendCall(GitserverClientListFilesFuncCall{v0, v1, v2, v3, v4, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the ListFiles method of
// the parent MockGitserverClient instance is invoked and the hook queue is
// empty.
func (f *GitserverClientListFilesFunc) SetDefaultHook(hook func(context.Context, store.Store, int, string, *regexp.Regexp) ([]string, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// ListFiles method of the parent MockGitserverClient instance inovkes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *GitserverClientListFilesFunc) PushHook(hook func(context.Context, store.Store, int, string, *regexp.Regexp) ([]string, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *GitserverClientListFilesFunc) SetDefaultReturn(r0 []string, r1 error) {
	f.SetDefaultHook(func(context.Context, store.Store, int, string, *regexp.Regexp) ([]string, error) {
		return r0, r1
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *GitserverClientListFilesFunc) PushReturn(r0 []string, r1 error) {
	f.PushHook(func(context.Context, store.Store, int, string, *regexp.Regexp) ([]string, error) {
		return r0, r1
	})
}

func (f *GitserverClientListFilesFunc) nextHook() func(context.Context, store.Store, int, string, *regexp.Regexp) ([]string, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitserverClientListFilesFunc) appendCall(r0 GitserverClientListFilesFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitserverClientListFilesFuncCall objects
// describing the invocations of this function.
func (f *GitserverClientListFilesFunc) History() []GitserverClientListFilesFuncCall {
	f.mutex.Lock()
	history := make([]GitserverClientListFilesFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitserverClientListFilesFuncCall is an object that describes an
// invocation of method ListFiles on an instance of MockGitserverClient.
type GitserverClientListFilesFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 store.Store
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 int
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 string
	// Arg4 is the value of the 5th argument passed to this method
	// invocation.
	Arg4 *regexp.Regexp
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []string
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitserverClientListFilesFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3, c.Arg4}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitserverClientListFilesFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// GitserverClientRawContentsFunc describes the behavior when the
// RawContents method of the parent MockGitserverClient instance is invoked.
type GitserverClientRawContentsFunc struct {
	defaultHook func(context.Context, store.Store, int, string, string) ([]byte, error)
	hooks       []func(context.Context, store.Store, int, string, string) ([]byte, error)
	history     []GitserverClientRawContentsFuncCall
	mutex       sync.Mutex
}

// RawContents delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockGitserverClient) RawContents(v0 context.Context, v1 store.Store, v2 int, v3 string, v4 string) ([]byte, error) {
	r0, r1 := m.RawContentsFunc.nextHook()(v0, v1, v2, v3, v4)
	m.RawContentsFunc.appendCall(GitserverClientRawContentsFuncCall{v0, v1, v2, v3, v4, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the RawContents method
// of the parent MockGitserverClient instance is invoked and the hook queue
// is empty.
func (f *GitserverClientRawContentsFunc) SetDefaultHook(hook func(context.Context, store.Store, int, string, string) ([]byte, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// RawContents method of the parent MockGitserverClient instance inovkes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *GitserverClientRawContentsFunc) PushHook(hook func(context.Context, store.Store, int, string, string) ([]byte, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *GitserverClientRawContentsFunc) SetDefaultReturn(r0 []byte, r1 error) {
	f.SetDefaultHook(func(context.Context, store.Store, int, string, string) ([]byte, error) {
		return r0, r1
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *GitserverClientRawContentsFunc) PushReturn(r0 []byte, r1 error) {
	f.PushHook(func(context.Context, store.Store, int, string, string) ([]byte, error) {
		return r0, r1
	})
}

func (f *GitserverClientRawContentsFunc) nextHook() func(context.Context, store.Store, int, string, string) ([]byte, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitserverClientRawContentsFunc) appendCall(r0 GitserverClientRawContentsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitserverClientRawContentsFuncCall objects
// describing the invocations of this function.
func (f *GitserverClientRawContentsFunc) History() []GitserverClientRawContentsFuncCall {
	f.mutex.Lock()
	history := make([]GitserverClientRawContentsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitserverClientRawContentsFuncCall is an object that describes an
// invocation of method RawContents on an instance of MockGitserverClient.
type GitserverClientRawContentsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 store.Store
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 int
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 string
	// Arg4 is the value of the 5th argument passed to this method
	// invocation.
	Arg4 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []byte
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitserverClientRawContentsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3, c.Arg4}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitserverClientRawContentsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}
