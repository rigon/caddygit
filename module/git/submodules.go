package git

import (
	// Submodules for the git app module registered here
	_ "github.com/rigon/caddygit/services/poll"
	_ "github.com/rigon/caddygit/services/webhook"
	_ "github.com/rigon/caddygit/services/webhook/generic"
	_ "github.com/rigon/caddygit/services/webhook/github"
)
