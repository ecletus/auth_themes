package clean

import (
	"errors"

	"github.com/qor/auth"
	"github.com/qor/auth/claims"
	"github.com/qor/auth/providers/password"
	"github.com/qor/render"
	"github.com/qor/qor"
	"github.com/moisespsena/template/html/template"
)

// ErrPasswordConfirmationNotMatch password confirmation not match error
var ErrPasswordConfirmationNotMatch = errors.New("password confirmation doesn't match password")


// New initialize clean theme
func New(config *auth.Config) *auth.Auth {
	if config == nil {
		config = &auth.Config{}
	}
	config.ViewPaths = append(config.ViewPaths, "github.com/qor/auth_themes/clean/views")

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

	Auth.Config.SetupDB(func(db *qor.DB) error {
		// Migrate Auth Identity model
		db.DB.AutoMigrate(Auth.Config.AuthIdentityModel)
		return nil
	})

	return Auth
}
