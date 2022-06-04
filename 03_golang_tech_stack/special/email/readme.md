<!--
 * @Author: ChenLong longchen2008@126.com
 * @Date: 2022-06-04 23:58:28
 * @LastEditors: ChenLong longchen2008@126.com
 * @LastEditTime: 2022-06-05 00:01:43
 * @FilePath: \study\03_golang_tech_stack\special\email\readme.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
# Golang发送邮件库email

## 人性化的golang电子邮件库

email包装设计的易于使用，足够灵活，不会受到限制。是一个非常人性化的golang邮件库

该email库包目前支持以下内容：

- 发件人、收件人、密件抄送和抄送字段
- “ test@example.com ”和“< test@example.com >”格式的电子邮件地址
- 文本和 HTML 消息正文
- 附件
- 阅读回执
- 自定义标题

## 安装

```bash
go get github.com/jordan-wright/email
```

注意：此库的版本 > 1 需要 Go v1.5 或更高版本。

如果您需要与以前的 Go 版本兼容，可以使用 gopkg.in/jordan-wright/email.v1 中的以前的包

## 实例

### 使用 QQ 发送电子邮件

```bash
package main

import (
    "net/smtp"

    "github.com/jordan-wright/email"
)

func main() {
    e := email.NewEmail()
    e.From = "郭宏志 <524060020@qq.com>"
    e.To = []string{"524060020@qq.com"}
    e.Bcc = []string{"524060020@qq.com"}
    e.Cc = []string{"524060020@qq.com"}
    e.Subject = "测试golang email库"
    e.Text = []byte("本文邮件内容!")
    e.HTML = []byte("<h1>html 邮件内容!</h1>")
    e.Send("smtp.qq.com:587", smtp.PlainAuth("", "524060020@qq.com", "xxx", "smtp.qq.com"))
}
```

### 另一种创建电子邮件的方法

你还可以通过创建结构直接创建电子邮件，如下所示：

```bash
e := &email.Email {
    To: []string{"test@example.com"},
    From: "Jordan Wright <test@gmail.com>",
    Subject: "Awesome Subject",
    Text: []byte("Text Body is, of course, supported!"),
    HTML: []byte("<h1>Fancy HTML is supported, too!</h1>"),
    Headers: textproto.MIMEHeader{},
}
```

### 从 io.Reader 创建电子邮件

io.Reader你还可以使用实现该接口的任何类型创建电子邮件email.NewEmailFromReader。

### 附加文件

```bash
e := NewEmail()
e.AttachFile("test.txt")
```

### 可重用连接池

```bash
(var ch <-chan *email.Email)
p := email.NewPool(
    "smtp.gmail.com:587",
    4,
    smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"),
)
for i := 0; i < 4; i++ {
    go func() {
        for e := range ch {
            p.Send(e, 10 * time.Second)
        }
    }()
}
```

## 文档

```bash
http://godoc.org/github.com/jordan-wright/email
```
