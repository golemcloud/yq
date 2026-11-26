//go:build wasip1

package yqlib

import (
	"errors"
	"io/fs"
	"os"
)

// isReadOnlyFile asks the host whether the file accepts writes. WASI has no
// permission bits, so Go reports every file as 0600; opening for writing
// without truncating or creating changes nothing.
func isReadOnlyFile(entryInfo os.FileInfo, _ os.FileInfo, name string) bool {
	if entryInfo.Mode()&os.ModeSymlink != 0 {
		return false
	}
	file, err := os.OpenFile(name, os.O_WRONLY, 0)
	if err != nil {
		return errors.Is(err, fs.ErrPermission)
	}
	_ = file.Close()
	return false
}
