package filesystem

import (
	"github.com/whitequark/go-git-git/v6/storage"
	"github.com/whitequark/go-git-git/v6/storage/filesystem/dotgit"
	"github.com/whitequark/go-git/v6/plumbing/cache"
)

type ModuleStorage struct {
	dir *dotgit.DotGit
}

func (s *ModuleStorage) Module(name string) (storage.Storer, error) {
	fs, err := s.dir.Module(name)
	if err != nil {
		return nil, err
	}

	return NewStorage(fs, cache.NewObjectLRUDefault()), nil
}
