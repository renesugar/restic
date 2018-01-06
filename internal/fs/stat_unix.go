// +build !windows

package fs

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

// extendedStat extracts info into an ExtendedFileInfo for unix based operating systems.
func extendedStat(fi os.FileInfo) ExtendedFileInfo {
	s, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		panic(fmt.Sprintf("conversion to syscall.Stat_t failed, type is %T", fi.Sys()))
	}

	extFI := ExtendedFileInfo{
		FileInfo:  fi,
		DeviceID:  s.Dev,
		Inode:     s.Ino,
		Links:     s.Nlink,
		UID:       s.Uid,
		GID:       s.Gid,
		Device:    s.Rdev,
		BlockSize: s.Blksize,
		Blocks:    s.Blocks,
		Size:      s.Size,

		AccessTime: time.Unix(s.Atim.Unix()),
		ModTime:    time.Unix(s.Mtim.Unix()),
	}

	return extFI
}
