package app

import (
	"io/fs"
	"os"
)

type fsContainer struct {
	fsys fs.FS
}

func newFsContainer(fsys fs.FS) *fsContainer {
	if fsys == nil {
		fsys = os.DirFS("/")
	}
	return &fsContainer{
		fsys: fsys,
	}
}

func (s *fsContainer) FS() fs.FS {
	return s.fsys
}
