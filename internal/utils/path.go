package utils

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

func ResolvePath(relativePath string) (string, bool) {
	root := ResolveProjectRoot()

	if root == "" {
		return "", false
	}

	relativePath = filepath.Clean(relativePath)

	if filepath.IsAbs(relativePath) {
		relativePath = relativePath[1:]
	}

	path := filepath.Join(root, relativePath)

	_, err := os.Stat(path)

	if err == nil {
		return path, true
	}

	if errors.Is(err, os.ErrNotExist) {
		return path, false
	}

	slog.Error(
		"failed to access path",
		"path", path,
		"error", err,
	)

	return path, false
}

func ResolveProjectRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		slog.Error("failed to determine project root: runtime.Caller failed")
		return ""
	}

	dir := filepath.Dir(filename)

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}

		parent := filepath.Dir(dir)

		if parent == dir {
			slog.Error("failed to find project root: go.mod not found")
			return ""
		}

		dir = parent
	}
}
