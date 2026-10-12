package main

import (
	"math/rand"
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

type Light struct {
	x int
	y int 
	color tcell.Color
	groupID int 
}

var colors = []tcell.Color{color.Red, color.White, color.Purple, color.Orange, color.Yellow, color.Blue, color.Green, color.BlueViolet}

func CreateLights(xsize int, ysize int) []Light {

	var lights []Light

	xmid := xsize / 2
	ymid := ysize / 2
	height := min(ymid, 15)
	widthCount := 1
	currentGroupID := 0

	for i := height - 1; i > 0; i-- {
		y := ysize - 3 - i
		left := xmid - widthCount
		width := 2*widthCount + 1
		row := height - 1 - i

		if row >= 2 {
			startPos := (row * 2) % 4 

			for pos := startPos; pos+2 < width; pos += 6  {
				for offset := range 3 {
					lights = append(lights, Light{
						x:     left + pos + offset,
						y:     y,
						color: color.Red,
						groupID: currentGroupID,
					})
				}
				currentGroupID++
				
			}
		}

		widthCount++
	}

	return lights
}

func DrawLightsUpdate(lights []Light, s tcell.Screen) {

	groupColors := make(map[int]tcell.Color)

	for i := range lights {
		gid := lights[i].groupID
		c, exists := groupColors[gid]
		if !exists {
			c = colors[rand.Intn(len(colors))]
			groupColors[gid] = c
		}

		lights[i].color = c
		lightStyle := tcell.StyleDefault.Background(color.Reset).Foreground(c)
		s.Put(lights[i].x, lights[i].y, "o", lightStyle)
	}
}

func DrawLightsNoUpdate(lights []Light, s tcell.Screen) {

	for light := range lights {
		lightStyle := tcell.StyleDefault.Background(color.Reset).Foreground(lights[light].color)
		s.Put(lights[light].x, lights[light].y, "o", lightStyle)
	}
}
