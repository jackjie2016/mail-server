package context

import (
	"fmt"
	"github.com/phper95/mail-server/internal/conf"
	"github.com/phper95/mail-server/internal/tools"
	"gopkg.in/macaron.v1"
	"net/url"
)

type ToggleOptions struct {
	SignInRequired  bool
	SignOutRequired bool
	AdminRequired   bool
	DisableCSRF     bool
}

func Toggle(options *ToggleOptions) macaron.Handler {

	return func(c *Context) {
		if !tools.IsFile("./install.lock") {
			fmt.Println("Toggle:not install")
			c.RedirectSubpath("/install")
			return
		}

		if options.SignInRequired {
			if !c.IsLogged {
				c.SetCookie("redirect_to", url.QueryEscape(conf.Web.Subpath+c.Req.RequestURI), 0, conf.Web.Subpath)
				c.RedirectSubpath("/user/login")
				return
			}
		}
	}
}
