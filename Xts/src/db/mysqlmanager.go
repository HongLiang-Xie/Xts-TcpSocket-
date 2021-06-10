package db

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
	"sync"
)




type MysqlManager struct {
	DB     *sqlx.DB     //数据库对象
	DBLock sync.RWMutex //数据库对象读写锁
}

func NewMysqlManager()*MysqlManager {
	init:=&MysqlManager{
		DB: nil,
		DBLock: sync.RWMutex{},
	}
	database,err:=sqlx.Open("mysql","root:88888888@tcp(localhost:3306)/login_db")
	if err!=nil{

		fmt.Println("open mysql failed",err)
		return nil
	}

	init.DB=database

	return init
}
func (mm* MysqlManager) Insert(query string, args ...interface{}) bool  { //增

	mm.DBLock.Lock()
	defer mm.DBLock.Unlock()

	_, err:=mm.DB.Exec(query,args...)
	if err!=nil{

		fmt.Println("mysql insert err",err)
		return false
	}
	return true
}
func (mm* MysqlManager) Select(dest interface{}, query string, args ...interface{})  { //查

	mm.DBLock.Lock()
	defer mm.DBLock.Unlock()

	err:=mm.DB.Select(dest,query, args ...)
	if err!=nil{

		fmt.Println("mysql Select err",err)
		return
	}
	return
}
func (mm* MysqlManager) Update(query string, args ...interface{}) bool  { //改

	mm.DBLock.Lock()
	defer mm.DBLock.Unlock()

	_,err:=mm.DB.Exec(query, args ...)
	if err!=nil{

		fmt.Println("mysql Update err",err)
		return false
	}
	return true
}
func (mm* MysqlManager) Delete(query string, args ...interface{}) bool  { //删除

	mm.DBLock.Lock()
	defer mm.DBLock.Unlock()

	_,err:=mm.DB.Exec(query, args ...)
	if err!=nil{

		fmt.Println("mysql Delete err",err)
		return false
	}
	return true
}