package clean

import (
	"errors"

	"github.com/moisespsena/template/html/template"
	"github.com/aghape/auth"
	"github.com/aghape/auth/claims"
	"github.com/aghape/auth/providers/password"
	"github.com/aghape/aghape"
	"github.com/aghape/render"
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
			FuncMapMaker: func(funcs *template.FuncValues, render *render.Render, context *qor.Context) error {
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