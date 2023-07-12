package _mongo

import (
	"fmt"
	"testing"
)

func TestMongoClient_Insert(t *testing.T) {
	err := New()
	if err != nil {
		fmt.Println(err)
	}
	UserDB.Insert()
}
func TestMongoCollection_Query(t *testing.T) {
	err := New()
	if err != nil {
		fmt.Println(err)
	}
	UserDB.Query()
}
func TestMongoCollection_Update(t *testing.T) {
	err := New()
	if err != nil {
		fmt.Println(err)
	}
	UserDB.UpdateById()
}
func TestMongoCollection_UpdateByOne(t *testing.T) {
	err := New()
	if err != nil {
		fmt.Println(err)
	}
	UserDB.UpdateByOne()
}
func TestMongoCollection_UpdateByMany(t *testing.T) {
	err := New()
	if err != nil {
		fmt.Println(err)
	}
	UserDB.UpdateByMany()
}
func TestMongoCollection_ReplaceOne(t *testing.T) {
	err := New()
	if err != nil {
		fmt.Println(err)
	}
	UserDB.ReplaceOne()
}
func TestMongoCollection_DeleteOne(t *testing.T) {
	err := New()
	if err != nil {
		fmt.Println(err)
	}
	UserDB.DeleteOne()
}
