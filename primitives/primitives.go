package primitives

// #cgo LDFLAGS: -lallegro_primitives
// #include <allegro5/allegro.h>
// #include <allegro5/allegro_primitives.h>
import "C"

import (
	"github.com/tapir/allegro"
	"unsafe"
)

func Init() bool {
	return bool(C.al_init_primitives_addon())
}

func Shutdown() {
	C.al_shutdown_primitives_addon()
}

func DrawLine(x1, y1, x2, y2 float32, c *allegro.Color, thickness float32) {
	C.al_draw_line(C.float(x1), C.float(y1), C.float(x2), C.float(y2), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))), C.float(thickness))
}

func DrawTriangle(x1, y1, x2, y2, x3, y3 float32, c *allegro.Color, thickness float32) {
	C.al_draw_triangle(C.float(x1), C.float(y1), C.float(x2), C.float(y2), C.float(x3), C.float(y3), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))), C.float(thickness))
}

func DrawFilledTriangle(x1, y1, x2, y2, x3, y3 float32, c *allegro.Color) {
	C.al_draw_filled_triangle(C.float(x1), C.float(y1), C.float(x2), C.float(y2), C.float(x3), C.float(y3), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))))
}

func DrawRectangle(x1, y1, x2, y2 float32, c *allegro.Color, thickness float32) {
	C.al_draw_rectangle(C.float(x1), C.float(y1), C.float(x2), C.float(y2), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))), C.float(thickness))
}

func DrawFilledRectangle(x1, y1, x2, y2 float32, c *allegro.Color) {
	C.al_draw_filled_rectangle(C.float(x1), C.float(y1), C.float(x2), C.float(y2), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))))
}

func DrawRoundedRectangle(x1, y1, x2, y2, rx, ry float32, c *allegro.Color, thickness float32) {
	C.al_draw_rounded_rectangle(C.float(x1), C.float(y1), C.float(x2), C.float(y2), C.float(rx), C.float(ry), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))), C.float(thickness))
}

func DrawRoundedFilledRectangle(x1, y1, x2, y2, rx, ry float32, c *allegro.Color) {
	C.al_draw_filled_rounded_rectangle(C.float(x1), C.float(y1), C.float(x2), C.float(y2), C.float(rx), C.float(ry), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))))
}

func DrawEllipse(cx, cy, rx, ry float32, c *allegro.Color, thickness float32) {
	C.al_draw_ellipse(C.float(cx), C.float(cy), C.float(rx), C.float(ry), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))), C.float(thickness))
}

func DrawFilledEllipse(cx, cy, rx, ry float32, c *allegro.Color) {
	C.al_draw_filled_ellipse(C.float(cx), C.float(cy), C.float(rx), C.float(ry), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))))
}

func DrawCircle(cx, cy, r float32, c *allegro.Color, thickness float32) {
	C.al_draw_circle(C.float(cx), C.float(cy), C.float(r), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))), C.float(thickness))
}

func DrawFilledCircle(cx, cy, r float32, c *allegro.Color) {
	C.al_draw_filled_circle(C.float(cx), C.float(cy), C.float(r), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))))
}

func DrawArc(cx, cy, r, start_theta, delta_theta float32, c *allegro.Color, thickness float32) {
	C.al_draw_arc(C.float(cx), C.float(cy), C.float(r), C.float(start_theta), C.float(delta_theta), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))), C.float(thickness))
}

func DrawSpline(points [8]float32, c *allegro.Color, thickness float32) {
	C.al_draw_spline((*C.float)(unsafe.Pointer(&points)), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))), C.float(thickness))
}

func DrawRibbon(points []float32, points_stride int32, c *allegro.Color, thickness float32, num_segments int32) {
	C.al_draw_ribbon((*C.float)(unsafe.Pointer(&points)), C.int(points_stride), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))), C.float(thickness), C.int(num_segments))
}

