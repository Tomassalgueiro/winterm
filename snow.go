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
	lastMoveRight bool 
}

func CreateSnow(flakes []Snowflake, xmax int) []Snowflake {
	if len(flakes) >= 50 {
		return flakes
	}

	num := rand.Intn(3)

	for range num {
		if len(flakes) >= 50 {
			break
		}

		newFlake := Snowflake{
			x: rand.Intn(xmax),
			y: 0,
			speed: 1,
			char: "❄",
			lastMoveRight: false,
		}
		flakes = append(flakes, newFlake)
	} 

	return flakes
}

func DrawSnow(flakes []Snowflake, s tcell.Screen) {
	defStyle := tcell.StyleDefault.Background(color.Reset).Foreground(color.Reset)
	snowStyle := tcell.StyleDefault.Background(color.Reset).Foreground(color.White)

	for flake := range flakes {
		if flakes[flake].y == 0 {
			s.Put(flakes[flake].x, flakes[flake].y, flakes[flake].char, snowStyle)
		} else {
			if  flakes[flake].lastMoveRight == true {
				s.Put(flakes[flake].x-1, flakes[flake].y-1, "", defStyle)
			} else {
				s.Put(flakes[flake].x+1, flakes[flake].y-1, "", defStyle)
			}
			s.Put(flakes[flake].x, flakes[flake].y, flakes[flake].char, snowStyle)
		}
	}
}

func UpdateSnow(flakes []Snowflake, xmax int, ymax int) []Snowflake {

	flakesSize := len(flakes)
	for flake := 0 ; flake < flakesSize; flake++ {
		// 0 to go left and 1 to go right
		movement := rand.Intn(2)
		if flakes[flake].y == ymax {
			flakesSize--
			flakes = append(flakes[:flake], flakes[flake+1:]...)
		} else {
			if flakes[flake].y + flakes[flake].speed <= ymax {	
				flakes[flake].y += flakes[flake].speed
			}
			if movement == 1 {
				flakes[flake].x += 1
				flakes[flake].lastMoveRight = true
			} else {
				flakes[flake].x -= 1 
				flakes[flake].lastMoveRight = false
			}
		}
	}

	return flakes
}
