package router

import (
	"github.com/phper95/mail-server/internal/app/context"
	"github.com/phper95/mail-server/internal/app/router/mail"
	"github.com/phper95/mail-server/internal/db"
)

const (
	HOME = "mail/list"
)

func Home(c *context.Context) {
	c.Data["PageIsHome"] = true
	c.Data["PageIsMail"] = true

	mail.RenderMailSearch(c, &mail.MailSearchOptions{
		PageSize: 10,
		OrderBy:  "id ASC",
		TplName:  HOME,
		Type:     db.MailSearchOptionsTypeInbox,
	})
}
