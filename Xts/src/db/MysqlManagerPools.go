package db

import (
	"fmt"
	"math/rand"
	"sync"
)

type MysqlManagerPools struct {

	MMPool map[int] *MysqlManager
	 MMPLock  sync.RWMutex //MMPool读写锁
}
var MySqlPoos *MysqlManagerPools
func init(){
	init:=&MysqlManagerPools{MMPool: make(map[int]*MysqlManager,10)}

	for i:=0;i<10;i++{
		init.MMPool[i]=NewMysqlManager()
	}
	fmt.Println("open mysql")
	MySqlPoos= init
}
func (mmp* MysqlManagerPools) GetMySqlManger()*MysqlManager  {
	mmp.MMPLock.Lock()
	defer mmp.MMPLock.Unlock()
	return mmp.MMPool[rand.Intn(9)]
}