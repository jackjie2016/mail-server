#!/bin/bash
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin


TAGRT_DIR=/usr/local/mail-server_dev
mkdir -p $TAGRT_DIR
cd $TAGRT_DIR

export GIT_COMMIT=$(git rev-parse HEAD)
export BUILD_TIME=$(date -u '+%Y-%m-%d %I:%M:%S %Z')


if [ ! -d $TAGRT_DIR/mail-server ]; then
	git clone https://github.com/phper95/mail-server
	cd $TAGRT_DIR/mail-server
else
	cd $TAGRT_DIR/mail-server
	git pull https://github.com/phper95/mail-server
fi

go mod tidy
go mod vendor


rm -rf mail-server
go build  -ldflags "-X \"github.com/phper95/mail-server/internal/conf.BuildTime=${BUILD_TIME}\" -X \"github.com/phper95/mail-server/internal/conf.BuildCommit=${GIT_COMMIT}\"" ./


cd $TAGRT_DIR/mail-server/scripts

sh make.sh

systemctl daemon-reload

service mail-server restart

cd $TAGRT_DIR/mail-server && ./mail-server -v


