package internal

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	RootPath = filepath.Join(filepath.Dir(b), "../")
)
