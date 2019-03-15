package clean

import (
	"github.com/ecletus/auth"
	"github.com/ecletus/auth/providers/password"
)

type Plugin struct {
}

func (Plugin) After() []interface{} {
	return []interface{}{&auth.Plugin{}, &password.Plugin{}}
}
