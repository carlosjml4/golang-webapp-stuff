package main

import (
	"fmt"
	"github.com/jimmykuu/webhelpers"
	"gopher"
)

func main() {
	c := gopher.DB.C("users")
	var users []gopher.User
	c.Find(nil).All(&users)

	smtpConfig := webhelpers.SmtpConfig{
		Username: gopher.Config.SmtpUsername,
		Password: gopher.Config.SmtpPassword,
		Host:     gopher.Config.SmtpHost,
		Addr:     gopher.Config.SmtpAddr,
	}

	for _, user := range users {
		subject := "Golang中国域名改为golangtc.com"
		message := user.Username + "，您好！\n\n由于golang.tc域名被Godaddy没收，现已不可继续使用，现在开始使用golangtc.com域名。希望继续参与到社区建设中来。\n\n Golang中国"
		err := webhelpers.SendMail(subject, message, gopher.Config.FromEmail, []string{user.Email}, smtpConfig, false)

		if err != nil {
			fmt.Println("error: ", err.Error())
		} else {
			fmt.Println("send to:", user.Username, user.Email)
		}
	}
}
