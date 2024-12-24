package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"server/global"
	"server/model/system"
	"time"
)

// InitGorm 初始化数据库并产生数据库全局变量
func InitGorm() {
	m := global.Config.Mysql
	if m.Dbname == "" {
		global.Log.Error("数据库名为空")
		panic("请检查是否存在数据库")
	}

	//
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:                  logger.Warn,            // 日志级别
			IgnoreRecordNotFoundError: true,                   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                  // 禁用彩色打印
		},
	)

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{Logger: newLogger})
	if err != nil {
		global.Log.Error("mysql连接失败", zap.Error(err))
		panic("数据连接出错了" + err.Error())
	}

	global.DB = db

	// 初始化数据库表
	registerTables()

	// 日志输出
	global.Log.Debug("数据库连接成功。开始运行", zap.Any("db", db))
}

// registerTables 初始化数据库表
func registerTables() {
	err := global.DB.AutoMigrate(
		// 系统
		system.UserModel{},
	)

	if err != nil {
		global.Log.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.Log.Info("register table success")
}
