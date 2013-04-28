/* An example showing bitmap flipping flags, original C version by Steven Wallace. */

package main

import (
	"github.com/tapir/allegro"
	"github.com/tapir/allegro/font"
	"github.com/tapir/allegro/imageio"
	"log"
)

const (
	INTERVAL float64 = 0.01
)

var (
	bmpX    float32 = 200
	bmpY    float32 = 200
	bmpDx   float32 = 96
	bmpDy   float32 = 96
	bmpFlag int32   = 0
)

func update(bmp *allegro.Bitmap) {
	target := allegro.GetTargetBitmap()
	displayW := target.GetWidth()
	displayH := target.GetHeight()
	bitmapW := bmp.GetWidth()
	bitmapH := bmp.GetHeight()

	bmpX += bmpDx * float32(INTERVAL)
	bmpY += bmpDy * float32(INTERVAL)

	/* Make sure bitmap is still on the screen. */
	if bmpY < 0 {
		bmpY = 0
		bmpDy *= -1
		bmpFlag &= allegro.FlipVertical
	}

	if bmpX < 0 {
		bmpX = 0
		bmpDx *= -1
		bmpFlag &= allegro.FlipHorizontal
	}

	if bmpY > float32(displayH)-float32(bitmapH) {
		bmpY = float32(displayH) - float32(bitmapH)
		bmpDy *= -1
		bmpFlag |= allegro.FlipVertical
	}

	if bmpX > float32(displayW)-float32(bitmapW) {
		bmpX = float32(displayW) - float32(bitmapW)
		bmpDx *= -1
		bmpFlag |= allegro.FlipHorizontal
	}
}

func main() {
	done := false
	redraw := true

	if !allegro.Init() {
		log.Fatalln("Failed to init Allegro.")
	}

	if !imageio.Init() {
		log.Fatalln("Failed to init IIO addon.")
	}

	font.InitFont()

	display := allegro.CreateDisplay(640, 480)
	defer display.Destroy()
	if display == nil {
		log.Fatalln("Error creating display.")
	}

	if !allegro.InstallKeyboard() {
		log.Fatalln("Error installing keyboard.")
	}

	font := font.LoadFont("data/fixed_font.tga", 0, 0)
	if font == nil {
		log.Fatalln("Error loading data/fixed_font.tga")
	}

	bmp := allegro.LoadBitmap("data/mysha.pcx")
	dispBmp := bmp
	defer bmp.Destroy()
	if bmp == nil {
		log.Fatalln("Error loading data/mysha.pcx")
	}
	text := "Display bitmap (space to change)"

	allegro.SetNewBitmapFlags(allegro.MemoryBitmap)
	memBmp := allegro.LoadBitmap("data/mysha.pcx")
	defer memBmp.Destroy()
	if memBmp == nil {
		log.Fatalln("Error loading data/mysha.pcx")
	}

	timer := allegro.CreateTimer(INTERVAL)
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
	queue.RegisterEventSource(display.GetEventSource())
	queue.RegisterEventSource(timer.GetEventSource())

	timer.Start()

	allegro.SetBlender(allegro.Add, allegro.One, allegro.InverseAlpha)

	for !done {
		if redraw && queue.IsEmpty() {
			update(bmp)
			allegro.ClearToColor(allegro.MapRgbF(0, 0, 0))
			bmp.DrawTinted(allegro.MapRgbaF(1, 1, 1, 0.5), bmpX, bmpY, bmpFlag)
			font.DrawText(allegro.MapRgbaF(1, 1, 1, 0.5), 0, 0, 0, text)
			allegro.FlipDisplay()
			redraw = false
		}

		event := queue.WaitForEvent()
		switch event.Type() {
		case allegro.EventKeyDown:
			if event.KeyboardKeycode() == allegro.KeyEscape {
				done = true
			} else if event.KeyboardKeycode() == allegro.KeySpace {
				if bmp == memBmp {
					bmp = dispBmp
					text = "Display bitmap (space to change)"
				} else {
					bmp = memBmp
					text = "Memory bitmap (space to change)"
				}
			}
		case allegro.EventDisplayClose:
			done = true
		case allegro.EventTimer:
			redraw = true
		}
	}
}
