package main

import (
	"math/rand"
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

type Snowflake struct {
	x int
	y int
	speed int 
	char string 
}

func CreateSnow(flakes []Snowflake, xmax int) []Snowflake {
	if len(flakes) > 30 {
		return flakes
	}

	num := rand.Intn(3)

	for range num {
		if len(flakes) > 30 {
			break
		}

		newFlake := Snowflake{
			x: rand.Intn(xmax),
			y: 0,
			speed: 1,
			char: "❄",
		}

		flakes = append(flakes, newFlake)
	} 
	return flakes
}

func DrawSnow(flakes []Snowflake, s tcell.Screen) {
	snowStyle := tcell.StyleDefault.Background(color.Reset).Foreground(color.White)

	for flake := range flakes {
		if flakes[flake].y == 0 {	
			s.Put(flakes[flake].x, 0, flakes[flake].char, snowStyle)
		} else {
			s.Put(flakes[flake].x, flakes[flake].y, flakes[flake].char, snowStyle)
		}
	}
}

func UpdateSnow(flakes []Snowflake, ymax int, s tcell.Screen) {

	for flake := range flakes {
		if flakes[flake].y == ymax {
			flakes = append(flakes[:flake], flakes[flake+1:]...)
		} else {
			if flakes[flake].y + flakes[flake].speed <= ymax {	
				flakes[flake].y += flakes[flake].speed
			}
		}
	}
}
