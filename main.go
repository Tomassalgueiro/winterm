package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)


func main() {
	defStyle := tcell.StyleDefault.Background(color.Reset).Foreground(color.Reset)
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.Clear()


	quit := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()

	running := true
	var flakes []Snowflake
	var lights []Light
	countUpdate := 0

	
	xmax, ymax := s.Size()
	lights = CreateLights(xmax, ymax)

	for running {
		select {

		case <- ticker.C:
			xmax, ymax := s.Size()
			flakes = CreateSnow(flakes, xmax)
			flakes = UpdateSnow(flakes, xmax, ymax)
			DrawSnow(flakes, s)

			DrawTree(xmax, ymax, s)

			switch countUpdate {
				case 0:
					DrawLightsUpdate(lights, s)
					countUpdate++
				case 3:
					DrawLightsNoUpdate(lights, s)
					countUpdate = 0
				default:
					DrawLightsNoUpdate(lights, s)
					countUpdate++
				}


			s.Show()

		case ev := <-s.EventQ():
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
				s.Clear()
				xmax, ymax = s.Size()
				lights = CreateLights(xmax, ymax)
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC || ev.Str() == "q" || ev.Str() == "q" {
					return
				}
			}
		}
	}
	s.Fini()
}
