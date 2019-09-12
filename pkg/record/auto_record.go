package record

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/auto_record/pkg/config"
	"github.com/tangxusc/auto_record/pkg/db"
	"math/rand"
	"strings"
	"time"
)

var AttendMachineNo = 100
var ci *cron.Cron

func Start(ctx context.Context) {
	ci = cron.New(cron.WithSeconds())
	morning(ci)
	night(ci)
	interval(ci)
	ci.Start()
}

func interval(c *cron.Cron) {
	_, _ = c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("now:", time.Now().Format(`2006-01-02 15:04:05`))
	})
}

func Stop(ctx context.Context) {
	ci.Stop()
}

func morning(c *cron.Cron) {
	id, e := c.AddFunc("0 51 8 * * *", func() {
		logrus.Debugf(`[record]trigger morning ...`)
		insert()
	})
	logrus.Debugf("[record]morning register:%v,error:%v", id, e)
}

func insert() {
	e := db.Exec(`INSERT INTO HR_AttendMachineData_Middle(Id,AttendMachineNo,EmployeeId,AttendTime,Status,ErrorMessage)
VALUES(?,?,?,convert(datetime,?, 20),?,?)`, getUuid(), AttendMachineNo, config.Instance.Record.EmployeeId, getTime(), `0`, ``)
	if e != nil {
		logrus.Warningf(`[record]morning insert error:%v`, e)
	}
}

//当前时间+随机值
func getTime() string {
	intn := rand.Intn(120)
	add := time.Now().Add(time.Duration(intn*-1) * time.Second)
	return add.Format(`2006-01-02 15:04:05`)
}

func getUuid() string {
	u2, err := uuid.NewV4()
	if err != nil {
		return ""
	}

	return strings.ToUpper(u2.String())
}

func night(c *cron.Cron) {
	id, e := c.AddFunc("0 7 18 * * *", func() {
		logrus.Debugf(`[record]trigger night ...`)
		insert()
	})
	logrus.Debugf("[record]night register:%v,error:%v", id, e)
}
