package mysqlCli

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {

	err := GetIntence()
	if err != nil {
		fmt.Printf("init db failed,err %v\n", err)
	}
	id := DB.Insert()
	DB.Query(id)
	DB.DeleteRow(id)
	DB.QueryMore(0)

}
func TestGorm(t *testing.T) {
	gormInsert()
	gormQuery()
}
