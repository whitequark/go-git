package main

import (
	"fmt"
	"os"

	. "github.com/whitequark/go-git-git/v6/_examples"
	"github.com/whitequark/go-git-git/v6/plumbing/transport/ssh"
	git "github.com/whitequark/go-git/v6"
)

func main() {
	CheckArgs("<url>", "<directory>")
	url, directory := os.Args[1], os.Args[2]

	authMethod, err := ssh.NewSSHAgentAuth("git")
	CheckIfError(err)

	// Clone the given repository to the given directory
	Info("git clone %s ", url)

	r, err := git.PlainClone(directory, &git.CloneOptions{
		Auth:     authMethod,
		URL:      url,
		Progress: os.Stdout,
	})
	CheckIfError(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
}
