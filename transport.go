package git

// Default supported transports.
import (
	_ "github.com/whitequark/go-git-git/v6/plumbing/transport/git"  // git transport
	_ "github.com/whitequark/go-git-git/v6/plumbing/transport/http" // http transport
	_ "github.com/whitequark/go-git-git/v6/plumbing/transport/ssh"  // ssh transport
	_ "github.com/whitequark/go-git/v6/plumbing/transport/file"     // file transport
)
