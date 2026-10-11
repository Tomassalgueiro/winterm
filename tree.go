package main

import (

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

// draws a tree (with trunk in the middle of the screen, trunk is ALWAYS three wide)
// max height is 18 (minus 3 (constant value, not changed with any size) from the trunk, so its 15 leaves)
func DrawTree(xsize int, ysize int, s tcell.Screen) {
	if xsize < 10 || ysize < 10 {
		s.PutStr(0,0,"Screen size too small")
		return
	} 

	leafStyle := tcell.StyleDefault.Background(color.Reset).Foreground(color.DarkGreen)
	trunkStyle := tcell.StyleDefault.Background(color.Reset).Foreground(color.BurlyWood)

	xmid := xsize/2
	ymid := ysize/2
	widthCount := 0
	height := min(ymid, 15)

	// we want width per height = past width + 3
	// first loop for height, second for width
	for i:= range 3 {
		for j := xmid - 1; j <= xmid + 1; j++ {
			s.Put(j,ysize-i-1,"|",trunkStyle)
		}
	}

	for i := height; i > 0; i-- {
		for j := xmid - widthCount; j <= xmid + widthCount; j++ {
			s.Put(j,ysize-3-i,"o",leafStyle)
		}
		widthCount += 1
	}

}
