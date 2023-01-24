package mail

import (
	"bytes"
	"context"
	"errors"
	// "fmt"
	"github.com/phper95/mail-server/internal/conf"
	"github.com/phper95/mail-server/internal/rspamd"
	"strings"
)

func RspamdCheck(content string) (bool, error, float64) {
	rspamdUrl := conf.Rspamd.Domain
	rspamdPassword := conf.Rspamd.Password
	rspamdJCS := conf.Rspamd.RecjectConditionScore

	if conf.Rspamd.Enable {

		client := rspamd.New(rspamdUrl)
		if !strings.EqualFold(rspamdPassword, "") {
			client.SetAuth(rspamdPassword)
		}

		_, err := client.Ping(context.Background())
		if err == nil {

			f := bytes.NewBuffer([]byte(content))
			email := rspamd.NewEmailFromReader(f)
			checkRes, err := client.Check(context.Background(), email)
			if err == nil {
				// for _, symVal := range checkRes.Symbols {
				// 	if symVal.Score > 0 {
				// 		fmt.Println(symVal.Name, symVal.Score, symVal.Description)
				// 	}
				// }
				if checkRes.Score > rspamdJCS {
					return true, errors.New("Judged as spam"), checkRes.Score
				}
			}

			return true, nil, checkRes.Score
		} else {
			return true, err, 0
		}
	}
	return true, nil, 0
}
