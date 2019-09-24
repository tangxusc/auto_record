package notify

import (
	"github.com/tangxusc/auto_record/pkg/config"
	"testing"
)

func TestSendMail(t *testing.T) {
	config.Instance.Mail.Address = ``
	config.Instance.Mail.Password = ``
	SendMail(config.Instance.Mail.Address, `1111`)
}
