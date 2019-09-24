package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const debugArgName = "debug"

func InitLog() {
	if Instance.Debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetReportCaller(true)
		logrus.Debug("[config]已开启debug模式...")
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}
}

func BindParameter(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&Instance.Debug, debugArgName, "v", false, "debug mod")

	cmd.PersistentFlags().StringVarP(&Instance.Db.Address, "db-address", "", "10.130.0.210", "db数据库连接地址")
	cmd.PersistentFlags().StringVarP(&Instance.Db.Port, "db-port", "", "1433", "db数据库端口")
	cmd.PersistentFlags().StringVarP(&Instance.Db.Database, "db-Database", "", "jw", "db数据库实例")
	cmd.PersistentFlags().StringVarP(&Instance.Db.Username, "db-Username", "", "jw", "db数据库用户名")
	cmd.PersistentFlags().StringVarP(&Instance.Db.Password, "db-Password", "", "jw@123", "db数据库密码")
	cmd.PersistentFlags().IntVarP(&Instance.Db.LifeTime, "db-LifeTime", "", 10, "db数据库连接最大连接周期(秒)")
	cmd.PersistentFlags().IntVarP(&Instance.Db.MaxOpen, "db-MaxOpen", "", 5, "db数据库最大连接数")
	cmd.PersistentFlags().IntVarP(&Instance.Db.MaxIdle, "db-MaxIdle", "", 5, "db数据库最大等待数量")

	cmd.PersistentFlags().StringVarP(&Instance.Record.EmployeeId, "r-employeeId", "", "000530", "工号")

	cmd.PersistentFlags().StringVarP(&Instance.Mail.Address, "m-addr", "", "562050688@qq.com", "邮箱地址")
	cmd.PersistentFlags().StringVarP(&Instance.Record.EmployeeId, "m-pwd", "", "", "邮箱密码")

}

type RecordConfig struct {
	EmployeeId string
}

type MailConfig struct {
	//发送人账号密码
	Address string
	//接收人
	Password string
}

type Config struct {
	Debug  bool
	Db     *DbConfig
	Record *RecordConfig
	Mail   *MailConfig
}

var Instance = &Config{
	Debug:  true,
	Db:     &DbConfig{},
	Record: &RecordConfig{},
	Mail:   &MailConfig{},
}

type DbConfig struct {
	Address  string
	Port     string
	Database string
	Username string
	Password string

	LifeTime int
	MaxOpen  int
	MaxIdle  int
}
