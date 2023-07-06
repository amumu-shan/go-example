package mysqlCli

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type guser struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

func (usr guser) TableName() string {
	return "user"
}

func (u *guser) BeforeCreate(tx *gorm.DB) (err error) {

	if u.Age > 10 {
		return errors.New("insert username21")
	}
	fmt.Println("BeforeCreate", u.Age, u.Name)
	return
}

var db *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	//gorm.Open(mysql.Open(dsn),&gorm.Config{})
	mysqlDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   //string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	sqlDB, err := mysqlDB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	db = mysqlDB
	err = sqlDB.Ping()
	if err != nil {
		fmt.Printf("gorm mysql failed,err:%#v", err)
	}
	fmt.Println("mysql connect success")
	return
}

func gormInsert() {
	gUser := []*guser{
		{Name: "ps", Age: 18},
		{Name: "lisi", Age: 18},
	}
	result := db.Create(gUser)
	if err := result.Error; err != nil {
		fmt.Printf("insert failed.,err:%#v", err)
	}

	fmt.Print("add id:")
	for _, user := range gUser {
		fmt.Printf("%d,", user.Id)
	}
	fmt.Printf("insert success.rowsAffected=%d", result.RowsAffected)

	users := [5000]guser{}
	db.Session(&gorm.Session{CreateBatchSize: 1000})
	for i := 0; i < 5000; i++ {
		users[i] = guser{Name: fmt.Sprintf("username%d", i), Age: i + 1}
	}
	db.Create(&users)
	fmt.Println("\ncreateBatchSize insert success")
}
func gormQuery() {
	user := guser{}
	// 获取第一条记录（主键升序）
	db.First(&user)
	fmt.Println(user.Id, user.Name, user.Age)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	user2 := guser{}
	db.Where(guser{Name: "lisi"}).First(&user2)
	fmt.Println("user2:", user2.Id, user2.Name, user2.Age)
	// 获取一条记录，没有指定排序字段
	user3 := guser{}
	db.Where("name like ?", "username%").Take(&user3)
	fmt.Println(user3.Id, user3.Name, user3.Age)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	db.Last(&user)
	fmt.Println(user.Id, user.Name, user.Age)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	result := db.First(&user)
	if result.Error != nil {
		fmt.Println("result.error:", result.Error)
	} else {
		fmt.Println("result.RowsAffected:", result.RowsAffected)
	}

	var user4 []guser
	db.Select("id", "name").Limit(10).Offset(5).Find(&user4)
	for _, v := range user4 {
		fmt.Println("limit:", v.Id, v.Name, v.Age)
	}

	var user5 []guser
	db.Table("user").Select("name", "max(age) as age").Group("name").Find(&user5)
	for _, v := range user5 {
		fmt.Println("Group:", v.Id, v.Name, v.Age)
	}

	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)

}
