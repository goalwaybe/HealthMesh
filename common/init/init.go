package init

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"medi-biz/common/config"
	"medi-biz/common/models"
	pb "medi-biz/proto/medi"
	"time"
)

func init() {
	ViperInit()
	MysqlInit()
	RedisInit()
	MediGrpcInit()
}

func MediGrpcInit() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常： %s\n", err)
	}
	config.MediClient = pb.NewMediClient(conn)
}

func ViperInit() {
	viper.SetConfigFile("../../../common/config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config.ConfigData)
	if err != nil {
		panic(err)
	}
	fmt.Println("配置文件获取成功：", config.ConfigData)
}

func MysqlInit() {
	data := config.ConfigData.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		data.User,
		data.Password,
		data.Host,
		data.Port,
		data.Database)
	var err error
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")

	err = config.DB.AutoMigrate(
		&models.Patient{},
		&models.Schedule{},
		&models.Doctor{},
		&models.Department{},
		&models.Appointment{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库迁移成功")

	sqlDB, err := config.DB.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func RedisInit() {
	data := config.ConfigData.Redis
	Addr := fmt.Sprintf("%s:%d", data.Host, data.Port)
	config.Rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: data.Password, // no password set
		DB:       data.Database, // use default DB
	})
	err := config.Rdb.Ping(config.Ctx).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis连接成功")
}
