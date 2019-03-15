package clean

import (
	"errors"

	"github.com/moisespsena/template/html/template"
	"github.com/ecletus/auth"
	"github.com/ecletus/auth/claims"
	"github.com/ecletus/auth/providers/password"
	"github.com/ecletus/core"
	"github.com/ecletus/render"
)

// ErrPasswordConfirmationNotMatch password confirmation not match error
var ErrPasswordConfirmationNotMatch = errors.New("password confirmation doesn't match password")

// New initialize clean theme
func New(config *auth.Config) *auth.Auth {
	if config == nil {
		config = &auth.Config{}
	}

	if config.Render == nil {
		config.Render = render.New(&render.Config{
			FuncMapMaker: func(funcs *template.FuncValues, render *render.Render, context *core.Context) error {
				ctx := context.GetI18nContext()
				return funcs.Set("t", func(key string, args ...interface{}) template.HTML {
					return template.HTML(ctx.T(key).DefaultAndDataFromArgs(args...).Get())
				})
			},
		})
	}

	Auth := auth.New(config)

	Auth.RegisterProvider(password.New(&password.Config{
		Confirmable: true,
		RegisterHandler: func(context *auth.Context) (*claims.Claims, error) {
			context.Request.ParseForm()

			if context.Request.Form.Get("confirm_password") != context.Request.Form.Get("password") {
				return nil, ErrPasswordConfirmationNotMatch
			}

			return password.DefaultRegisterHandler(context)
		},
	}))

	return Auth
}