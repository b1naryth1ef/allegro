package allegro

// #include <allegro5/allegro.h>
import "C"

type Config struct {
	allegroConfig *C.ALLEGRO_CONFIG
}
