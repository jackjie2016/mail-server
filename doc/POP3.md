
### POP | DEBUG
```
phper95@phper95 ~ % telnet pop.qq.com 110 
Trying 203.205.232.7...
Connected to pop.qq.com.
Escape character is '^]'.
+OK XMail POP3 Server v1.0 Service Ready(XMail v1.0)
user 627293072
+OK
pass xxxxxx
+OK
```

## 相关命令
```
USER username 认证用户名
PASS password 认证密码认证，认证通过则状态转换 
APOP name,digest 认可一种安全传输口令的办法，执行成功导致状态转换，请参见 RFC 1321 。
STAT 处理请求 server 回送邮箱统计资料，如邮件数、 邮件总字节数
UIDL n 处理 server 返回用于该指定邮件的唯一标识， 如果没有指定，返回所有的。
LIST n 处理 server 返回指定邮件的大小等 
RETR n 处理 server 返回邮件的全部文本 
DELE n 处理 server 标记删除，QUIT 命令执行时才真正删除
RSET 处理撤消所有的 DELE 命令 
TOP n,m 处理 返回 n 号邮件的前 m 行内容，m 必须是自然数 
NOOP 处理 server 返回一个肯定的响应 
QUIT 希望结束会话。如果 server 处于"处理" 状态，则现在进入"更新"状态，删除那些标记成删除的邮件。如果 server 处于"认可"状态，则结束会话时 server 不进入"更新"状态 。
```
