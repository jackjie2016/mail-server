### 协议参考
- https://tools.ietf.org/html/rfc5321
- https://www.rfc-editor.org/rfc/rfc5321.txt

### SMTP设计图如下：
                  +----------+                +----------+
      +------+    |          |                |          |
      | User |<-->|          |      SMTP      |          |
      +------+    |  Client- |Commands/Replies| Server-  |
      +------+    |   SMTP   |<-------------->|    SMTP  |    +------+
      | File |<-->|          |    and Mail    |          |<-->| File |
      |System|    |          |                |          |    |System|
      +------+    +----------+                +----------+    +------+
                   SMTP client                SMTP server

### mail.qq.com | cmd | 例子
```
phper95@phper95 ~ % telnet smtp.qq.com 25
Trying 240e:ff:f100:8019::6a...
Connected to smtp.qq.com.
Escape character is '^]'.
220 newxmesmtplogicsvrsza9.qq.com XMail Esmtp QQ Mail Server.
helo m
250-newxmesmtplogicsvrsza9.qq.com-9.21.152.27-23666742
250-SIZE 73400320
250 OK
auth login
334 VXNlcm5hbWU6
用户    ---[base64]
334 UGFzc3dvcmQ6
授权码    ---[base64]
235 Authentication successful
mail from:<627293072@qq.com>
250 OK.
rcpt to: <phper95@163.com>
250 OK
data
354 End data with <CR><LF>.<CR><LF>.
from:627293072
to: phper95
subject: test

data
.
250 OK: queued as.
```

### localhost | cmd | debug
```
phper95@phper95 ~ % telnet 127.0.0.1 25
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
220 Anti-spam GT for Coremail System(mail-server)
helo m
250 ok
auth login 
334 dXNlcm5hbWU6
YWRtaW4=
334 UGFzc3dvcmQ6
YWRtaW4=
235 Authentication successful
mail from:<admin@xxx.com>
250 Mail OK
rcpt to:<phper95@163.com>
250 Mail OK
data
354 End data with <CR><LF>.<CR><LF>
```



