package allegro

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_shader.h>
// #include <allegro5/allegro_shader_hlsl.h>
import "C"

func GetDefaultHlslVertexShader() string {
	return C.GoString(C.al_get_default_hlsl_vertex_shader())
}

func GetDefaultHlslPixelShader() string {
	return C.GoString(C.al_get_default_hlsl_pixel_shader())
}
