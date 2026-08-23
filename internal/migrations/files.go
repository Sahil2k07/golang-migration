package migrations

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Sahil2k07/golang-migration/internal/utils"
)

type migrationFile struct {
	FileName string
	FilePath string
	FileHash string
	SQL      string
}

func getMigrationFiles(root string, failOnModified bool, histories []migrationHistory) ([]migrationFile, error) {
	var files []migrationFile

	path, exists := utils.ResolvePath(root)
	if !exists {
		// maybe some folders are not added yet because not needed
		slog.Warn("Root mentioned not found", "root", root)
		return files, nil
	}

	projectRoot := utils.ResolveProjectRoot()

	applied := make(map[string]migrationHistory, len(histories))
	for _, migration := range histories {
		applied[migration.FilePath] = migration
	}

	err := filepath.WalkDir(path, func(currentPath string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			return nil
		}

		if !strings.EqualFold(filepath.Ext(entry.Name()), ".sql") {
			return nil
		}

		sqlBytes, err := os.ReadFile(currentPath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %q: %w", currentPath, err)
		}

		relativePath, err := filepath.Rel(projectRoot, currentPath)
		if err != nil {
			return fmt.Errorf("failed to resolve migration path %q: %w", currentPath, err)
		}

		relativePath = filepath.ToSlash(relativePath)

		filehash := generateHash(sqlBytes)

		if migration, exists := applied[relativePath]; exists {
			// The file hasn't changed since applied
			if migration.FileHash == filehash {
				slog.Warn("Migration Skipped", "file", entry.Name())
				return nil
			}

			// The file was modifed and we didn't intent to
			if failOnModified {
				return fmt.Errorf("Migration file was modified after getting applied: %v", entry.Name())
			}
		}

		files = append(files, migrationFile{
			FileName: entry.Name(),
			FilePath: relativePath,
			FileHash: filehash,
			SQL:      string(sqlBytes),
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].FilePath < files[j].FilePath
	})

	return files, nil
}

func generateHash(sql []byte) string {
	hash := sha256.Sum256(sql)

	return hex.EncodeToString(hash[:])
}
