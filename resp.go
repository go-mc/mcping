package mcping

import (
	"github.com/Tnze/go-mc/chat"
	_ "github.com/Tnze/go-mc/data/lang/en-us"
	"github.com/google/uuid"
)

type Status struct {
	Description chat.Message
	Players     struct {
		Max    int
		Online int
		Sample []struct {
			ID   uuid.UUID
			Name string
		}
	}
	Version struct {
		Name     string
		Protocol int
	}
	Favicon Icon
}

// Icon should be a PNG image that is Base64 encoded
// (without newlines: \n, new lines no longer work since 1.13)
// and prepended with data:image/png;base64,.
type Icon string
