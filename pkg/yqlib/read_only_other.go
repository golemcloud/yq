//go:build !wasip1

package yqlib

import "os"

// isReadOnlyFile reports a regular file without any write permission bit.
func isReadOnlyFile(entryInfo os.FileInfo, info os.FileInfo, _ string) bool {
	return entryInfo.Mode()&os.ModeSymlink == 0 && info.Mode().Perm()&0o222 == 0
}
