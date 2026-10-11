package main

import (
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

// draws a tree (with trunk in the middle of the screen, trunk is ALWAYS three wide)
// max height is 18 (minus 3 (constant value, not changed with any size) from the trunk, so its 15 leaves)
func DrawTree(xsize int, ysize int, s tcell.Screen) {
	if xsize < 10 || ysize < 10 {
		return
	} 

	leafStyle := tcell.StyleDefault.Background(color.Reset).Foreground(color.DarkGreen)
	//trunkStyle := tcell.StyleDefault.Background(color.Reset).Foreground(color.BurlyWood)

	for i:=0;i<3;i++ {
		s.Put(i,i,"*",leafStyle)
	}


}
