package tools

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eryajf/go-ldap-admin/config"
	"github.com/patrickmn/go-cache"

	"strconv"

	"gopkg.in/gomail.v2"
)

// 验证码放到缓存当中
var VerificationCodeCache = cache.New(24*time.Hour, 48*time.Hour)

func email(mailTo []string, subject string, body string) error {
	mailConn := map[string]string{
		"user": config.Conf.Email.User,
		"pass": config.Conf.Email.Pass,
		"host": config.Conf.Email.Host,
		"port": config.Conf.Email.Port,
	}
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	newmail := gomail.NewMessage()

	newmail.SetHeader("From", newmail.FormatAddress(mailConn["user"], config.Conf.Email.From))
	newmail.SetHeader("To", mailTo...)    //发送给多个用户
	newmail.SetHeader("Subject", subject) //设置邮件主题
	newmail.SetBody("text/html", body)    //设置邮件正文

	do := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	return do.DialAndSend(newmail)
}

func SendMail(sendto []string, pass string) error {
	subject := "重置LDAP密码成功"
	// 邮件正文
	body := fmt.Sprintf("<li><a>更改之后的密码为: %s </a></li>", pass)
	return email(sendto, subject, body)
}

// SendCode 发送验证码
func SendCode(sendto []string) error {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	// 把验证码信息放到cache，以便于验证时拿到
	VerificationCodeCache.Set(sendto[0], vcode, time.Minute*5)
	subject := "验证码-重置密码"
	//发送的内容
	body := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为 %s ,为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, vcode)
	return email(sendto, subject, body)
}
