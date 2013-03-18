package color

// #cgo LDFLAGS: -lallegro_color
// #include <allegro5/allegro.h>
// #include <allegro5/allegro_color.h>
import "C"

import "github.com/tapir/allegro"


func ColorCmyk(c, m, y, k float32) allegro.Color {
	return (allegro.Color)(C.al_color_cmyk(C.float(c), C.float(m), C.float(y), C.float(k)))
}

func ColorCmykToRgb(cyan, magenta, yellow, key float32) (float32, float32, float32) {
	var r, g, b C.float
	C.al_color_cmyk_to_rgb(C.float(cyan), C.float(magenta), C.float(yellow), C.float(key), &r, &g, &b)
	return float32(r), float32(g), float32(b)
}
