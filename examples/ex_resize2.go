package main

import (
	"github.com/tapir/allegro"
	"github.com/tapir/allegro/imageio"
	"log"
)

func main() {
	if !allegro.Init() {
		log.Fatalln("Could not init Allegro.")
	}

	if !allegro.InstallKeyboard() {
		log.Fatalln("Could not init Keyboard.")
	}
	if !imageio.Init() {
		log.Fatalln("Could not init ImageIO.")
	}

	allegro.SetNewDisplayFlags(allegro.Resizable | allegro.GenerateExposeEvents)
	display := allegro.CreateDisplay(640, 480)
	defer display.Destroy()
	if display == nil {
		log.Fatalln("Unable to set any graphic mode")
	}

	allegro.SetNewBitmapFlags(allegro.MemoryBitmap)
	bmp := allegro.LoadBitmap("data/mysha.pcx")
	defer bmp.Destroy()
	if bmp == nil {
		log.Fatalln("Unable to load image")
	}

	queue := allegro.CreateEventQueue()
	defer queue.Destroy()
	if queue == nil {
		log.Fatalln("Could not create Queue.")
	}

	queue.RegisterEventSource(allegro.GetKeyboardEventSource())
	queue.RegisterEventSource(display.GetEventSource())

	redraw := true
	for {
		if redraw && queue.IsEmpty() {
			allegro.ClearToColor(allegro.MapRgb(255, 0, 0))
			bmpW := float32(bmp.GetWidth())
			bmpH := float32(bmp.GetHeight())
			displayW := float32(display.GetWidth())
			displayH := float32(display.GetHeight())
			bmp.DrawScaled(0, 0, bmpW, bmpH, 0, 0, displayW, displayH, 0)
			allegro.FlipDisplay()
			redraw = false
		}

		event := queue.WaitForEvent()
		if event.Type == allegro.EventDisplayResize {
			display.AcknowledgeResize()
			redraw = true
		}
		if event.Type == allegro.EventDisplayExpose {
			redraw = true
		}
		if event.Type == allegro.EventKeyDown && event.KeyboardE.Keycode == allegro.KeyEscape {
			break
		}
		if event.Type == allegro.EventDisplayClose {
			break
		}
	}
}
