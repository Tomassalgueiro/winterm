package main 

import (
	"log"
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


	// Here's an example of how to inject a keystroke where it will
	// be picked up by a future read of the event queue.  Note that
	// care should be used to avoid blocking writes to the queue if
	// this is done from the same thread that is responsible for reading
	// the queue, or else a single-party deadlock might occur.
	// s.EventQ() <- tcell.NewEventKey(tcell.KeyRune, rune('a'), 0)

	for {
		xmax, ymax := s.Size()
		DrawTree(xmax, ymax, s)

		s.Show()

		ev := <-s.EventQ()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC || ev.Str() == "q" || ev.Str() == "Q" {
				return
			}
		}
	}
}
