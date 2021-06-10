package core

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	am:=NewAOIManager(100,300,4,200,450,5)
	am.AOIManagerInfo()
	for _,grid:=range am.Grids {
		grid.GridInfo()
	}
	fmt.Println(am)

}
func TestGetSurroundGrids(t *testing.T) {

	am:=NewAOIManager(0,250,5,0,250,5)

	for gid,_:=range am.Grids{

		grids:= am.GetSurroundGrids(gid)
		fmt.Println("gid",gid,"grids len=",len(grids))
		gIDs:=make([]int,0,len(grids))
		for _,grid:=range grids{
			gIDs=append(gIDs,int(grid.GID))
		}
		fmt.Println("SurroundGrids:",gIDs)
	}


}
