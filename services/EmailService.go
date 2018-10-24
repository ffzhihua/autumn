package services

import (
	"autumn/models"
	"autumn/tools/cfg"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"
	"time"
	"strconv"
)

type EmailService struct {
}

//发送邮件验证码
func (this *EmailService) SendCode(email string, code string, language string) bool {

	subject := cfg.Lang("email", language, "title").String()
	content := fmt.Sprintf(cfg.Lang("email", language, "content").String(), code,
		strconv.Itoa(int(cfg.Get("sms", "expire").Int())/60))

	done := make(chan error)

	go this.SendMailUsingTLS(done, email, subject, content)

	if err := <-done; err != nil {
		return false
	}

	return true
}

//验证接口
func (this *EmailService) Verify(email string, code string, usage bool) int {

	//验证码输入错误
	if len(code) != 6 {
		return 11000
	}

	//短信测试开启
	if code == cfg.Get("sms", "test").String() {
		return 0
	}

	eml := new(models.Email).Info(email, code)

	//无效
	if eml.Id == 0 {
		return 11005
	}

	//过期
	if eml.Expire < time.Now().Unix() {
		return 11006
	}

	if eml.Status == 1 {
		return 11007
	}

	if usage {
		eml.Status = 1
		eml.Update()
	}

	return 0
}

//return a smtp client
func (this *EmailService) dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func (this *EmailService) SendMailUsingTLS(done chan error, email string, subject string, content string) {
	auth := smtp.PlainAuth("",
		cfg.Get("smtp", "username").String(),
		cfg.Get("smtp", "password").String(),
		cfg.Get("smtp", "host").String(),
	)

	//create smtp client
	c, err := this.dial(cfg.Get("smtp", "server").String())
	if err != nil {
		log.Println("Service.EmailService.smtp error:", err)
		done <- err
		return
	}

	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Service.EmailService.smtp error:", err)
				done <- err
				return
			}
		}
	}

	if err = c.Mail(cfg.Get("smtp", "from").String()); err != nil {
		log.Println("Service.EmailService.smtpe error:", err)
		done <- err
		return
	}

	if err = c.Rcpt(email); err != nil {
		log.Println("Service.EmailService.smtpe error:", err)
		done <- err
		return
	}

	w, err := c.Data()
	if err != nil {
		log.Println("Service.EmailService.smtpe error:", err)
		done <- err
		return
	}

	mailContent := []byte(strings.Replace("From:RRNC<"+(cfg.Get("smtp", "username").String())+">~To :"+
		email+"~Subject:"+subject+"~Content-Type: text/html;charset=UTF-8~",
		"~", "\r\n", -1) + "\r\n" + content)

	_, err = w.Write(mailContent)
	if err != nil {
		log.Println("Service.EmailService.smtpe error:", err)
		done <- err
		return
	}

	err = w.Close()
	if err != nil {
		log.Println("Service.EmailService.smtpe error:", err)
		done <- err
	}

	c.Quit()

	done <- nil
	return
}

func (this *EmailService) sendMail(done chan error, email string, subject string, content string) {
	auth := smtp.PlainAuth("",
		cfg.Get("smtp", "username").String(),
		cfg.Get("smtp", "password").String(),
		cfg.Get("smtp", "host").String(),
	)

	header := []byte(strings.Replace("From:RRNC<"+(cfg.Get("smtp", "username").String())+">~To :"+
		email+"~Subject:"+subject+"~Content-Type: text/html;charset=UTF-8~",
		"~", "\r\n", -1) + "\r\n" + content)

	sendTo := strings.Split(email, ":")

	err := smtp.SendMail(cfg.Get("smtp", "server").String(),
		auth, cfg.Get("smtp", "username").String(),
		sendTo,
		header)

	if err != nil {
		log.Println("Send Email Error:", err)
	}

	done <- err

}