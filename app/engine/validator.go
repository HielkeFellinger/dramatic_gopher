package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getValidatedGameFile(baseDir, userPath string) (string, error) {
	// Resolve symlinks and get absolute paths
	checkedAbsBase, err := filepath.EvalSymlinks(baseDir)
	if err != nil {
		return "", fmt.Errorf("failed to resolve symlinks in base dir: %w", err)
	}
	checkedAbsPath, err := filepath.EvalSymlinks(filepath.Join(baseDir, userPath))
	if err != nil {
		return "", fmt.Errorf("failed to resolve symlinks in path: %w", err)
	}

	// Ensure path is inside base
	rel, err := filepath.Rel(checkedAbsBase, checkedAbsPath)
	if err != nil {
		return "", fmt.Errorf("failed to compute relative path: %w", err)
	}
	if strings.HasPrefix(rel, "..") {
		return "", fmt.Errorf("path %q escapes base directory %q", userPath, baseDir)
	}

	// Require .yaml or .yml (case-insensitive)
	if ext := filepath.Ext(checkedAbsPath); !(strings.EqualFold(ext, ".yaml") || strings.EqualFold(ext, ".yml")) {
		return "", fmt.Errorf("invalid extension for file %q: must be .yaml or .yml", checkedAbsPath)
	}

	// Ensure it exists, is accessible and is a regular file
	info, statErr := os.Stat(checkedAbsPath)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			return "", fmt.Errorf("file does not exist: %s", checkedAbsPath)
		}
		return "", fmt.Errorf("error checking file: %w", statErr)
	}
	if info.IsDir() {
		return "", fmt.Errorf("expected a file, got a directory: %s", checkedAbsPath)
	}

	return checkedAbsPath, nil
}
