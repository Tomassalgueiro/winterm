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

var ColorThemes = map[string][]tcell.Color{
	"classic": {
		color.Red,
		color.Gold,
		color.White,
		color.LimeGreen,
		color.DarkOrange,
	},
	"neon": {
		color.DeepPink,
		color.Aqua,
		color.Lime,
		color.Chartreuse,
		color.HotPink,
	},
	"frozen": {
		color.DeepSkyBlue,
		color.DodgerBlue,
		color.LightCyan,
		color.White,
		color.GhostWhite,
		color.Turquoise,
	},
	"warm": {
		color.OrangeRed,
		color.Coral,
		color.DarkOrange,
		color.Orange,
		color.Gold,
		color.Yellow,
	},
	"berry": {
		color.Crimson,
		color.DeepPink,
		color.DarkViolet,
		color.Purple,
		color.BlueViolet,
		color.MediumVioletRed,
	},
	"all": {
		color.Red, color.Crimson, color.OrangeRed, color.Coral, color.DarkOrange,
		color.Orange, color.Gold, color.Yellow, color.DeepPink, color.HotPink,
		color.Fuchsia, color.Purple, color.DarkViolet, color.BlueViolet,
		color.DodgerBlue, color.DeepSkyBlue, color.RoyalBlue, color.Aqua,
		color.Lime, color.LimeGreen, color.White,
	},
}

var ActiveColors = ColorThemes["classic"]

func SetTheme(name string) {
	if selected, ok := ColorThemes[name]; ok {
		ActiveColors = selected
	} else {
		ActiveColors = ColorThemes["classic"]
	}
}

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
			c = ActiveColors[rand.Intn(len(ActiveColors))]
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
