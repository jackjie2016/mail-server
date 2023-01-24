#!/bin/bash

_os=`uname`
_path=`pwd`
_dir=`dirname $_path`

sed "s:{APP_PATH}:${_dir}:g" $_dir/scripts/init.d/mail-server.tpl > $_dir/scripts/init.d/mail-server
chmod +x $_dir/scripts/init.d/mail-server


if [ -d /etc/init.d ];then
	cp $_dir/scripts/init.d/mail-server /etc/init.d/mail-server
	chmod +x /etc/init.d/mail-server
fi

echo `dirname $_path`
