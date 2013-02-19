package allegro

// #include <allegro5/allegro.h>
import "C"

type Path struct {
	allegroPath *C.ALLEGRO_PATH
}
