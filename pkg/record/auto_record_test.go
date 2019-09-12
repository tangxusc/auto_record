package record

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/tangxusc/auto_record/pkg/config"
	"github.com/tangxusc/auto_record/pkg/db"
	"log"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	time := getTime()
	fmt.Println(time)
}

func TestGetUuid(t *testing.T) {
	time := getUuid()
	fmt.Println(time)
}

func TestInsert(t *testing.T) {
	config.Instance.Record.EmployeeId = `000530`
	config.Instance.Db = &config.DbConfig{
		Address:  "10.130.0.210",
		Port:     "1433",
		Database: "jw",
		Username: "jw",
		Password: "jw@123",
		LifeTime: 10,
		MaxOpen:  5,
		MaxIdle:  5,
	}

	db.Conn(context.TODO())
	insert()
	db.Disconnection(context.TODO())
}

func TestStart(t *testing.T) {
	log.Println("Starting...")

	// 定义一个cron运行器
	c := cron.New(cron.WithSeconds())
	// 定时5秒，每5秒执行print5
	c.AddFunc("*/5 * * * * *", print5)
	// 定时15秒，每5秒执行print5
	c.AddFunc("*/15 * * * * *", print15)

	// 开始
	c.Start()
	c.Run()
	defer c.Stop()

	// 这是一个使用time包实现的定时器，与cron做对比
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
			print10()
		}
	}
}

func print5() {
	log.Println("Run 5s cron")
}

func print10() {
	log.Println("Run 10s cron")
}

func print15() {
	log.Println("Run 15s cron")
}
