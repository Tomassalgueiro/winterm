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
}

var colors = []tcell.Color{color.Red, color.White, color.Purple, color.Orange, color.Yellow, color.Blue, color.Green, color.BlueViolet}

func CreateLights(xsize int, ysize int) []Light {

	var lights []Light

	xmid := xsize / 2
	ymid := ysize / 2
	height := min(ymid, 15)

	widthCount := 1

	for i := height - 1; i > 0; i-- {
		y := ysize - 3 - i
		left := xmid - widthCount
		width := 2*widthCount + 1

		row := height - 1 - i
		if row >= 2 {
			offset := (row - 2) * 3

			for pos := range width {
				if (pos+offset)%7 < 3 {
					lights = append(lights, Light{
						x:     left + pos,
						y:     y,
						color: color.Red,
					})
				}
			}
		}

		widthCount++
	}

	return lights
}

func DrawLightsUpdate(lights []Light, s tcell.Screen) {

	for light := range lights {
		num := rand.Intn(8)
		lightStyle := tcell.StyleDefault.Background(color.Reset).Foreground(colors[num])
		lights[light].color = colors[num]
		s.Put(lights[light].x, lights[light].y, "o", lightStyle)
	}
}

func DrawLightsNoUpdate(lights []Light, s tcell.Screen) {

	for light := range lights {
		lightStyle := tcell.StyleDefault.Background(color.Reset).Foreground(lights[light].color)
		s.Put(lights[light].x, lights[light].y, "o", lightStyle)
	}
}
