package allegro

// #include <allegro5/allegro.h>
import "C"

type File struct {
	allegroFile *C.ALLEGRO_FILE
}
