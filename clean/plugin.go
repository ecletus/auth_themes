package clean

import (
	"github.com/aghape/auth"
	"github.com/aghape/auth/providers/password"
)

type Plugin struct {
}

func (Plugin) After() []interface{} {
	return []interface{}{&auth.Plugin{}, &password.Plugin{}}
}
