package core

import (
	"fmt"
	"sync"
)

type WorldManager struct {
	//AoiManager 当前世界地图Aoi的管理模块
	AOIMgr*AOIManager
	//当前全部在线Player集合
	Players map[int32]  *Player
	//保护player集合的锁
	pLock sync.RWMutex

}
//初始化方法

 var WorldMgrObj*WorldManager
func init(){
	WorldMgrObj=&WorldManager{
		//初始化AOI管理
		AOIMgr:  NewAOIManager(AOI_MIN_X,AOI_MAX_X,AOI_CNTS_X,AOI_MIN_Y,AOI_MAX_Y,AOI_CNTS_Y),
		//初始化玩家集合
		Players: make(map[int32]  *Player),
		pLock:   sync.RWMutex{},
	}

}

//添加一个在线玩家

func (wm* WorldManager) AddPlayer(player *Player)  {
	//添加玩家在World中
	wm.pLock.Lock()
	defer wm.pLock.Unlock()
	wm.Players[player.Pid]=player
	//添加玩家在AOIManager中
	wm.AOIMgr.AddSpritesByPoint(int(player.Pid),player.X,player.Z)


}
//删除一个下线玩家
func (wm* WorldManager) RemovePlayer(player *Player)  {
	//删除玩家在World中
	wm.pLock.Lock()
	defer wm.pLock.Unlock()
	delete(wm.Players,player.Pid)
	//删除玩家在AOIManager中
	wm.AOIMgr.RemoveSpritesByPoint(int(player.Pid),player.X,player.Z)
}
//通过pid查询player对象
func (wm* WorldManager) GetPlayer(Pid int32)*Player {

	wm.pLock.Lock()
	defer wm.pLock.Unlock()
	player,isTrue:= wm.Players[Pid]

	if!isTrue{
		fmt.Println("player not FOUND!","Pid=",Pid)

	}
	return player
}

//获取全部在线玩家
func (wm* WorldManager) GetAllPlayer()[]*Player {

	var players[] *Player
	players=make([]*Player,0)
	wm.pLock.Lock()
	defer wm.pLock.Unlock()
	for _,value :=range  wm.Players{

		players=append(players, value)

	}

	return players
}