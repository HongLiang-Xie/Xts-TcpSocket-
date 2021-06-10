package core

import "fmt"

//定义一些AOI的边界值 宏

const(
	AOI_MIN_X int32 =85
	AOI_MAX_X int32 =410
	AOI_CNTS_X int32 =10
	AOI_MIN_Y int32 =75
	AOI_MAX_Y int32 =400
	AOI_CNTS_Y int32 =20

)

type AOIManager struct {
	//区域的左边界坐标
	MinX int32
	//区域的右边界坐标
	MaxX int32
	// x方向格子的数量
	CntsX int32
	//区域的上边界坐标
	MinY int32
	//区域的下边界坐标
	MaxY int32
	// y方向格子的数量
	CntsY int32
	//当前区域内格子的集合
	Grids map[int]*Grid
}

//初始化

func NewAOIManager(MinX, MaxX, CntsX, MinY, MaxY, CntsY int32) *AOIManager {

	init := &AOIManager{
		MinX:  MinX,
		MaxX:  MaxX,
		CntsX: CntsX,
		MinY:  MinY,
		MaxY:  MaxY,
		CntsY: CntsY,
		Grids: make(map[int]*Grid),
	}
	for y := 0; y < int(CntsY); y++ {
		for x := 0; x < int(CntsX); x++ {
			gid := y*int(CntsX) + x
			init.Grids[gid] = NewGrid(int32(gid),
				init.MinX+int32(init.gridWidth())*int32(x),
				init.MinX+int32(init.gridWidth())*int32(x+1),
				init.MinY+int32(init.gridHeight())*int32(y),
				init.MinY+int32(init.gridHeight())*int32(y+1))
		}
	}

	return init
}

//得到每个格子在X轴方向的宽度
func (am *AOIManager) gridWidth() int {

	return int((am.MaxX - am.MinX) / am.CntsX)
}

//得到每个格子在Y轴方向的宽度
func (am *AOIManager) gridHeight() int {
	return int((am.MaxY - am.MinY) / am.CntsY)
}

//打印格子信息
func (am *AOIManager) AOIManagerInfo() {

	fmt.Println(am)
}

//获得周边格子
func (am *AOIManager) GetSurroundGrids(gID int) (grids []*Grid) {
	//判断ID是否在AOIManager中
	_, isTrue := am.Grids[gID]
	if !isTrue {

		return nil
	}
	//将当前gid本身加入九宫格切片中
	grids = append(grids, am.Grids[gID])
	//需要gid的左边是否有格子？右边是否有格子
	//需要通过gid的到当前格子的x轴编号 idx=id%am.CntsX
	idx := gID % int(am.CntsX)
	//判断idx左边是否有格子
	if idx > 0 {
		grids = append(grids, am.Grids[gID-1])
	}
	//判断idx右边是否有格子
	if idx < int(am.CntsX)-1 {

		grids = append(grids, am.Grids[gID+1])
	}
	//新建个TempGridsX用于存放GridsX的集合，此集合用于判断上下是否有格子idy:=v/int(am.CntsY)
	gridsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gridsX = append(gridsX, int(v.GID))
	}
	//遍历gridsx集合中每个格子的gid
	for _, v := range gridsX {

		idy := v / int(am.CntsY)
		//判断上边是否有格子
		if idy > 0 {
			grids = append(grids, am.Grids[v-int(am.CntsY)])
		}
		//判断下边是否有格子
		if idy < int(am.CntsY)-1 {
			grids = append(grids, am.Grids[v+int(am.CntsY)])
		}
	}
	return grids
}

//通过x,y坐标得到当前GID编号
func (am *AOIManager) GetGIDByPoint(x, y float32) int {
	idx := (int(x) - int(am.MinX)) / am.gridWidth()
	idy := (int(y) - int(am.MinY)) / am.gridHeight()

	return idy*int(am.CntsX) + idx

}

//通过坐标得到周边九宫格内全部的Sprite
func (am *AOIManager) GetSpriteByPoint(x, y float32) (playerIDs []int) {
	//得到当前玩家的Gid格子id
	gid := am.GetGIDByPoint(x, y)
	//通过GID得到周边九宫格信息
	grids := am.GetSurroundGrids(gid)
	//将九宫格信息里的全部sprite的id 全部累加到PlayerIDs中
	for _, v := range grids {

		playerIDs = append(playerIDs, v.GetSprite()...)
	//	fmt.Println("==> grid ID:",v.GID,"sprites:",v.GetSprite())
	}

	return
}
//添加一个sprite 到格子中
func(am*AOIManager) AddSpriteToGrid(sID,gID int){
	am.Grids[gID].Add(sID)
}
//移除格子中一个sprite
func(am*AOIManager) RemoveSpriteToGrid(sID,gID int){
	am.Grids[gID].Remove(sID)
}
//通过Gid获取全部的SpriteID
func (am*AOIManager)GetSpritesByGID(gID int)(Sprites[]int){

	return am.Grids[gID].GetSprite()
}
//通过坐标将Sprite 添加到格子中
func (am*AOIManager)AddSpritesByPoint(sID int,x,y float32){
	//得到当前玩家的Gid格子id
	gid := am.GetGIDByPoint(x, y)
	am.AddSpriteToGrid(sID,gid)

}
//通过坐标将Sprite 从格子中移除
func (am*AOIManager)RemoveSpritesByPoint(sID int,x,y float32){
	//得到当前玩家的Gid格子id
	gid := am.GetGIDByPoint(x, y)
	am.RemoveSpriteToGrid(sID,gid)

}