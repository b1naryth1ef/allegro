package main

import (
	"fmt"
	"github.com/tapir/allegro"
	"github.com/tapir/allegro/imageio"
	"log"
	"time"
)

func main() {
	var zoom float32 = 1.0
	redraw := true
	filename := "data/mysha.pcx"

	if !allegro.Init() {
		log.Fatalln("Could not init Allegro.")
	}

	if !allegro.InstallMouse() {
		log.Fatalln("Could not init Mouse.")
	}
	if !allegro.InstallKeyboard() {
		log.Fatalln("Could not init Keyboard.")
	}

	if !imageio.Init() {
		log.Fatalln("Could not init ImageIO.")
	}

	display := allegro.CreateDisplay(640, 480)
	defer display.Destroy()
	if display == nil {
		log.Fatalln("Could not create Display.")
	}
	display.SetWindowTitle(filename)

	t := time.Now()
	bitmap := allegro.LoadBitmap(filename)
	defer bitmap.Destroy()
	duration := time.Since(t)
	if bitmap == nil {
		log.Fatalf("%s not found or failed to load\n", filename)
	}
	fmt.Printf("Loading took %s\n", duration.String())

	timer := allegro.CreateTimer(1.0 / 30)
	defer timer.Destroy()
	if timer == nil {
		log.Fatalln("Could not create Timer.")
	}

	queue := allegro.CreateEventQueue()
	defer queue.Destroy()
	if queue == nil {
		log.Fatalln("Could not create Queue.")
	}
	queue.RegisterEventSource(allegro.GetKeyboardEventSource())
	queue.RegisterEventSource(allegro.GetMouseEventSource())
	queue.RegisterEventSource(display.GetEventSource())
	queue.RegisterEventSource(timer.GetEventSource())

	timer.Start()

	for {
		event := queue.WaitForEvent()

		if event.Type() == allegro.EventDisplayOrientation {
			o := event.DisplayOrientation()
			if o == allegro.DisplayOrientation0Degrees {
				fmt.Println("0 degrees")
			} else if o == allegro.DisplayOrientation90Degrees {
				fmt.Println("90 degrees")
			} else if o == allegro.DisplayOrientation180Degrees {
				fmt.Println("180 degrees")
			} else if o == allegro.DisplayOrientation270Degrees {
				fmt.Println("270 degrees")
			} else if o == allegro.DisplayOrientationFaceUp {
				fmt.Println("Face up")
			} else if o == allegro.DisplayOrientationFaceDown {
				fmt.Println("Face down")
			}
		}
		if event.Type() == allegro.EventDisplayClose {
			fmt.Println("Closing.")
			break
		}
		if event.Type() == allegro.EventMouseButtonDown {
			fmt.Println("Mouse down + ", event.MouseButton())
		}
		if event.Type() == allegro.EventMouseButtonUp {
			fmt.Println("Mouse up + ", event.MouseButton())
		}
		if event.Type() == allegro.EventKeyChar {
			k := event.KeyboardKeycode()
			u := event.KeyboardUnichar()
			fmt.Println("Keycode + ", k)
			fmt.Println("Unichar + ", u)

			if k == allegro.KeyEscape {
				fmt.Println("Closing (Escape key).")
				break
			}
			if u == '1' {
				zoom = 1
			}
			if u == '+' {
				zoom *= 1.1
			}
			if u == '-' {
				zoom /= 1.1
			}
			if u == 'f' {
				zoom = float32(display.GetWidth()) / float32(bitmap.GetWidth())
			}
		}
		if event.Type() == allegro.EventTimer {
			redraw = true
		}

		if redraw && queue.IsEmpty() {
			redraw = false
			allegro.ClearToColor(allegro.MapRgb(255, 255, 255))
			if zoom == 1 {
				bitmap.Draw(0, 0, 0)
			} else {
				bitmap.DrawScaledRotated(0, 0, 0, 0, zoom, zoom, 0, 0)
			}
			allegro.FlipDisplay()
		}
	}
}
