package common

import (
	"fmt"

	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局mysql数据库变量
var DB *gorm.DB

// 初始化mysql数据库
func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.Conf.Mysql.Username,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
		config.Conf.Mysql.Charset,
		config.Conf.Mysql.Collation,
		config.Conf.Mysql.Query,
	)
	// 隐藏密码
	showDsn := fmt.Sprintf(
		"%s:******@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.Conf.Mysql.Username,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
		config.Conf.Mysql.Charset,
		config.Conf.Mysql.Collation,
		config.Conf.Mysql.Query,
	)
	// Log.Info("数据库连接DSN: ", showDsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
		//// 指定表前缀
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix: config.Conf.Mysql.TablePrefix + "_",
		//},
	})
	if err != nil {
		Log.Panicf("初始化mysql数据库异常: %v", err)
		panic(fmt.Errorf("初始化mysql数据库异常: %v", err))
	}

	// 开启mysql日志
	if config.Conf.Mysql.LogMode {
		db.Debug()
	}
	// 全局DB赋值
	DB = db
	// 自动迁移表结构
	dbAutoMigrate()
	Log.Infof("初始化mysql数据库完成! dsn: %s", showDsn)
}

// 自动迁移表结构
func dbAutoMigrate() {
	DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Group{},
		&model.Menu{},
		&model.Api{},
		&model.OperationLog{},
	)
}
