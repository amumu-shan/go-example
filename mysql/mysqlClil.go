package mysqlCli

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
	age  int
}

type dbStruct struct {
	db *sql.DB
}

var DB *dbStruct

func GetIntence() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s faild,err:%v\n", dsn, err)
	}
	db.SetMaxIdleConns(10)
	fmt.Println("mysql连接成功~")
	DB = &dbStruct{db: db}
	return
}

func (db *dbStruct) Insert() (id int64) {
	sqlStr := `insert into user(name,age) values('zhangsan',18)`
	ret, err := db.db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	lastId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Println("id=", lastId)
	return lastId
}
func (db *dbStruct) DeleteRow(id int64) {
	sqlStr := `delete from user where id=?`
	ret, err := db.db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete faild,err:%v\n", err)
		return
	}
	n, _ := ret.RowsAffected()
	fmt.Printf("删除了%d行数据\n", n)
}

func (db *dbStruct) Query(id int64) {
	sqlStr := "select id,name,age from user where id =?;"
	row := db.db.QueryRow(sqlStr, id)
	var u1 user
	row.Scan(&u1.id, &u1.name, &u1.age)
	fmt.Printf("u1:%+v\n", u1)

}
func (db *dbStruct) QueryMore(id int64) {
	sqlStr := "select id,name,age from user where id >?;"
	rows, err := db.db.Query(sqlStr, id)
	defer rows.Close()
	if err != nil {
		fmt.Printf("%s query failed,err:%v \n", sqlStr, err)
		return
	}
	for rows.Next() {
		var u1 user
		rows.Scan(&u1.id, &u1.name, &u1.age)
		fmt.Printf("u1:%#v\n", u1)
	}
}
