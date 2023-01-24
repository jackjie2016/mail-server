# IMAP CMD
```
xx@xx ~ % telnet 127.0.0.1 143
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
* OK [CAPABILITY IMAP4 IMAP4rev1 ID AUTH=PLAIN AUTH=LOGIN AUTH=XOAUTH2 NAMESPACE] mail-server ready
B CAPABILITY
* OK Coremail System IMap Server Ready(mail-server)
* CAPABILITY IMAP4rev1 XLIST SPECIAL-USE ID LITERAL+ STARTTLS XAPPLEPUSHSERVICE UIDPLUS X-CM-EXT-1
B OK CAPABILITY completed
C ID ("name" "com.tencent.foxmail" "version" "1.5.2.94526" "vendor" "Tencent Limited" "contact" "foxmail@foxmail.com")
* ID NIL
C OK ID completed
D LOGIN 627293072 "xxxxx" 
```

### SIMPLE CMD
```
C CLOSE
C OK CLOSE Done
------------

G LOGOUT
* BYE LOGOUT received
G OK LOGOUT Completed
------------
```





### LIST
```
a list "" *
* LIST (\NoSelect \HasChildren) "/" "&UXZO1mWHTvZZOQ-"
* LIST (\HasNoChildren) "/" "INBOX"
* LIST (\HasNoChildren) "/" "Sent Messages"
* LIST (\HasNoChildren) "/" "Drafts"
* LIST (\HasNoChildren) "/" "Deleted Messages"
* LIST (\HasNoChildren) "/" "Junk"
a OK LIST completed
```

### SELECT
```
cc select "Sent Messages"
* 4438 EXISTS
* 0 RECENT
* OK [UIDVALIDITY 1393914013] UID validity status
* OK [UIDNEXT 4502] Predicted next UID
* FLAGS (\Answered \Flagged \Deleted \Draft \Seen)
* OK [PERMANENTFLAGS (\* \Answered \Flagged \Deleted \Draft \Seen)] Permanent flags
cc OK [READ-WRITE] SELECT complete
```
### FETCH
```
G FETCH 2 BODY[TEXT]
* 2 FETCH (BODY[TEXT] {853}
This is a multi-part message in MIME format.

------=_NextPart_5D0CE15E_8847F2B0_78A0120B
Content-Type: text/plain;
	charset="utf-8"
Content-Transfer-Encoding: base64

c2QKc2QKcwpk

------=_NextPart_5D0CE15E_8847F2B0_78A0120B
Content-Type: text/html;
	charset="utf-8"
Content-Transfer-Encoding: base64

PGRpdiBzdHlsZT0iZm9udDoxNHB4LzEuNSAnTHVjaWRhIEdyYW5kZScsICflvq7ova/pm4Xp
u5EnO2NvbG9yOiMzMzM7Ij48cCBzdHlsZT0iZm9udDoxMnB4LzEuNSAnLlNGIE5TIFRleHQn
O21hcmdpbjowOyI+PGJyPjwvcD48cCBzdHlsZT0iZm9udDoxMnB4LzEuNSAnLlNGIE5TIFRl
eHQnO21hcmdpbjowOyI+c2Q8L3A+PHAgc3R5bGU9ImZvbnQ6MTJweC8xLjUgJy5TRiBOUyBU
ZXh0JzttYXJnaW46MDsiPnNkPC9wPjxwIHN0eWxlPSJmb250OjEycHgvMS41ICcuU0YgTlMg
VGV4dCc7bWFyZ2luOjA7Ij5zPC9wPjxwIHN0eWxlPSJmb250OjEycHgvMS41ICcuU0YgTlMg
VGV4dCc7bWFyZ2luOjA7Ij5kPC9wPjwvZGl2Pg==

------=_NextPart_5D0CE15E_8847F2B0_78A0120B--
)
G OK FETCH Completed
```

### UID FETCH
```
B UID FETCH 3440:3441 (uid flags rfc822.size bodystructure body.peek[header])
* 153 FETCH (UID 3440 FLAGS (\Seen) RFC822.SIZE 21210 BODYSTRUCTURE (("TEXT" "PLAIN" ("charset" "UTF-8") NIL NIL "QUOTED-PRINTABLE" 1485 34 NIL NIL NIL)("TEXT" "HTML" ("charset" "UTF-8") NIL NIL "QUOTED-PRINTABLE" 16011 220 NIL NIL NIL) "ALTERNATIVE" ("BOUNDARY" "----=_Part_288952104_622524367.1622600794892") NIL NIL) BODY[HEADER] {3335}
Received: from 17.179.250.241 (unknown [11.176.227.221])
	by newmx31.qq.com (NewMx) with SMTP id 
	for <627293072@qq.com>; Wed, 02 Jun 2021 11:36:59 +0800
X-QQ-FEAT: nHaaMjwLeTwma0n/3c+mkKIi0t1ttWAsrpGEOKF5aYU=
X-QQ-MAILINFO: M9mpTqh4QKvqS3wheh6UiYeV12VwznCqWASphqCrPs6AwYjLvDel8FdDW
	/K2XRAbScy4a/O/8i002j9cpRDPoRjnf0Fx5Dy3w3iKSw7QXiqT9nUgG6z2PK/eCiA5aBY0
	MaA01gxn2RXKfMGzWzjejRQSRZtpPj7dYQ==
X-QQ-mid: mxszb88t1622605017txojeeuuj
X-QQ-CSender: developer_bounces@insideapple.apple.com
X-QQ-ORGSender: developer_bounces@insideapple.apple.com
X-QQ-XMAILINFO: Ml8T/6JCI2CtvmjVm5VetX6NANP4GHkd7ApZdXT0MM9X88OalUqj+gxj6Wo+6Q
	 bUUNyzLn86J3YYZy6btRDF9Cd/4oGQpAV7bUQlPfMfi5nTlK8RxqHyZZCOCe0UyylcX4HMz1ShkQ
	 LifkqhVnnwyqndeD+zqKJCmhs3r/mge0L2fhhXvXb9b2/UAhILwjnjWCpoo+R5WXFLibe2mew3RN
	 t7sl9iRNRQxxk0Hlu1TLyxcUF1mizj+JT7ReTjMIGshRuPvnbSdgQxxuPQEcAIIsfs+oLfcAjfrd
	 WByElXYjm6lBuMo6mrXPF0eMfAWZ8947mAko65dx7oHTdSrjPHWKuIkOI3kJ9THTWLqP1ndLbZm/
	 QLBATnXyZWt9uSJF5a0LvwEWApae7W4Ww/cHLfvWvA1u77qCDRoEHzLhOOlPR5xSKQRIqaB21yHu
	 YwR4s348Vd54nOlvMTmfuqZRBcmDELbbKPrVm6egYzOBu9ZK3dYTbRf47vm9u3cp8Yt28w7FnWLp
	 ATelHel68PlF3SzjwgnrFrmeBYfTpXAJS+k+t+qEMWWk8I/avONJqmlRm380yOx2nDxpvkP4jAIx
	 ssKUfInmr/D1xh+e4evW1D9F6caWJRxU4Wt6JakVHGU+OYScioZ2dNUp5g7ApZ7bBLojvoyM0A10
	 NCy/Z9pjdSunNpdTNKI/9KpQ7TyoriotC6JbdrAJVutebdI1Up+LXBvocANWxxj/GFOJtMUzEoug
	 O3u4dm3SzKLhzwZXgVuNddjh1pv25x/rQNcSTS2yAAZKFTu3RbdpRZnoj2P0ptJO8e9rynSKU/qb
	 evb7I//n8DuKtzrb8/NbBSZ85dkLuK8SPHyCYeDZt2b0ByJV/OqOVasVNUImTKRZumgomkM3PNYj
	 Dl9laywRi0jQG+Vr90KoVeRPBAvlq0iYDMieXQe9CA7+7TLG3NMLtls7l2PcWyvWfb/q2Vc9iK3L
	 4pqV1WDYpaiFfgBwVM61HJKYtppUm1
DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
	d=insideapple.apple.com; s=insideapple0517; t=1622600794;
	bh=h+NQMcEzI3C9xQHHRNmsLrw2bwEoMdcWC04BYbfx1ME=;
	h=Date:From:To:Message-ID:Subject:Content-Type;
	b=xThMNnVZDsGc6cucFvCW/YwsvKR+E2OisDivhR7Ba3DMa/Jr75lwnmwS18jBI6/hL
	 rX9b3MsM68Z2hQTpAWimNT0qRpMY3mV4UsonLgLG6tKnxCA3AhC1nOY4a0UFUSimUT
	 uAulmf7WMyOaeWqCjadYmciKsI4qduaYu8SgL4ECi73sSLCUvTcBHRyI8fGimKDzGX
	 Q0oLo1K02ng1Gmv6r6naRDi1+PtdjyPjfiK9RzjuBBnsixxUNPeMCT1Rur6Mt/HaPH
	 BpaoafjxaKWOziPIE7/wEfg1/B3VjFazIKqI94mdD4tT90Rq1cE/oQ80DGs5ta/UZC
	 32r0kYhtviUsA==
Date: Wed, 2 Jun 2021 02:26:34 +0000 (GMT)
From: Apple Developer <developer@insideapple.apple.com>
To: 627293072@qq.com
Message-ID: <1563273383.288952106.1622600794892@Insideapple.apple.com>
Subject: =?UTF-8?Q?<=E5=B9=BF=E5=91=8A>_One_week_until_WWDC21?=
MIME-Version: 1.0
Content-Type: multipart/alternative; 
	boundary="----=_Part_288952104_622524367.1622600794892"
X-Attach-Flag: N
X-Sent-To: 627293072@qq.com,2,dVSSzOMDZAAu6UHtqk4P%2F4PLeVwmALo5LV6PQfb6rqXSMJQK5zEwUdx501t%2FSC1qLwrmOU9FnljRtc0S94Q5qB%2FCx0yUt4oZFf87548AAQLK26Sdl5CBoZBjOxjGe9gV11OovB4crPz8epGd8GQR4r4872t8Ag5HFdbDKsnCkre3fV4xy%2FsPCEmm2a9czUh6
X-TXN_ID: e603a91c-a15c-46a5-a87e-5e1c3adf14d7
X-DKIM_SIGN_REQUIRED: YES
x-evs: BYPASS
X-EmailType-Id: 2210431
X-Business-Group: cbx_wlm
X-Broadcast-Id: 2210431
List-Unsubscribe: <https://developer.apple.com/unsubscribe/?v=2&la=en_us&a=0j9Q%2BCQP8SbooeZ9pEWjmNxZ8iIJXK4K7WEZr0g3VxzvejIezNyBa%2FHQHycnSJmhRUKImFygEmVJ46lLIf7esyVzIBI6dny5D6c9vpWE8qxg6c8pGaJdLFeNztwA2o2UrzYrM8%2FTygZrUbb8FE9FskhcJFgBMmGsevmnTamp7dfa3KY4oWHT71%2Fw7ZpNEm43M44mp2qCAThUV0lk76H8iQ%3D%3D>
X-BTPH: reg
X-mp: d

)
B OK UID FETCH Completed
```
