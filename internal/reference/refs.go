package reference

import (
	"io"

	"github.com/whitequark/go-git-git/v6/storage"
	"github.com/whitequark/go-git/v6/plumbing"
)

// References returns all references from the storage.
func References(st storage.Storer) ([]*plumbing.Reference, error) {
	var localRefs []*plumbing.Reference

	iter, err := st.IterReferences()
	if err != nil {
		return nil, err
	}

	for {
		ref, err := iter.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		localRefs = append(localRefs, ref)
	}

	return localRefs, nil
}
