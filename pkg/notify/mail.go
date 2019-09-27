package notify

import (
	"github.com/go-gomail/gomail"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/auto_record/pkg/config"
)

func SendMail(target string, timestring string, msg string) {
	m := gomail.NewMessage()
	//发件人
	m.SetAddressHeader("From", config.Instance.Mail.Address, "发件人")
	//收件人
	m.SetHeader("To", m.FormatAddress(target, "收件人"))
	//主题
	m.SetHeader("Subject", "auto_record")
	m.SetBody(`text/html`, `auto record at `+timestring+` error:`+msg)
	dialer := gomail.NewDialer(`smtp.qq.com`, 465, config.Instance.Mail.Address, config.Instance.Mail.Password)
	e := dialer.DialAndSend(m)
	if e != nil {
		logrus.Errorf("[mail]send mail at [%s] error:%v", timestring, e)
	} else {
		logrus.Infof("[mail]send mail at [%s] success", timestring)
	}

}
