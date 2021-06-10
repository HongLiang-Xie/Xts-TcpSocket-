package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	//格子ID
	GID int32
	//格子左边界X坐标
	MinX int32
	//格子右边界X坐标
	MaxX int32
	//格子上边界Y坐标
	MinY int32
	//格子下边界Y坐标
	MaxY int32
	//当前格子内物品和玩家的集合
	Sprites map[int]bool
	//保护当前集合的锁
	SpritesLock sync.RWMutex
}

//初始化一个格子信息
func NewGrid(GID,MinX,MaxX,MinY,MaxY int32)  *Grid {

	return &Grid{
		GID: GID,
		MinX: MinX,
		MaxX: MaxX,
		MinY: MinY,
		MaxY: MaxY,
		Sprites: make(map[int] bool),
	}

}

//添加玩家或物品
func (g *Grid) Add(Sprite int) {

	g.SpritesLock.Lock()
	defer g.SpritesLock.Unlock()

	g.Sprites[Sprite] = true

}

//移除玩家或物品
func (g *Grid) Remove(Sprite int) {

	g.SpritesLock.Lock()
	defer g.SpritesLock.Unlock()

	delete(g.Sprites, Sprite)

}

//获取格子内所有玩家和物品集合
func (g *Grid) GetSprite()(Sprites []int) {

	g.SpritesLock.Lock()
	defer g.SpritesLock.Unlock()
	for ID, _ := range g.Sprites {

		Sprites = append(Sprites, ID)

	}
	return Sprites

}
func (g*Grid)GridInfo()  {

	fmt.Println("gridinfo:",g)
}