package db

import (
	"gitee.com/phper95/pkg/es"
	"github.com/phper95/mail-server/internal/conf"
	"github.com/phper95/mail-server/internal/log"
	"strings"
)

func InitES() error {
	es.InitClientWithOptions(es.DefaultClient, strings.Split(conf.ES.Host, ","),
		conf.ES.User,
		conf.ES.Password,
		es.WithScheme("https"))
	if err != nil {
		log.Errorf("InitClientWithOptions error", "client", es.DefaultClient)
	}
	return err
}
