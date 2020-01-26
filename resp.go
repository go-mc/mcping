package mcping

import (
	"encoding/base64"
	"errors"
	"strings"

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

var IconFormatErr = errors.New("data format error")
var IconAbsentErr = errors.New("icon not present")

// ToPNG decode base64-icon, return a PNG image
// Take care of there is not safety check, image may contain malicious code .
func (i Icon) ToPNG() ([]byte, error) {
	const prefix = "data:image/png;base64,"
	if i == "" {
		return nil, IconAbsentErr
	}
	if !strings.HasPrefix(string(i), prefix) {
		return nil, IconFormatErr
	}
	return base64.StdEncoding.DecodeString(strings.TrimPrefix(string(i), prefix))
}
