TODO List
=========
* Implement modules: direct3d, opengl
* Implement addons: image, audio, audio_codec, font, color, physfs, primitives, video, memfile
* Use GO's test facilities.
* Port some allegro examples.
* Documentation? (Is it worth copying over the Allegro docs? How will it be maintained?).
* Idiomatic error handling integration?
* void* []byte intptr_t -> unsafe.Pointer

Callbacks not implemented
=========================
* system.go -> al_install_system
* display.go -> al_acknowledge_drawing_resume
* events.go -> al_emit_user_event
* files.go -> al_fopen_interface | al_set_new_file_interface | al_set_standard_file_interface | al_get_new_file_interface | al_create_file_handle | al_get_file_userdata
* filesystem.go -> al_set_fs_interface | al_set_standard_fs_interface | al_get_fs_interface
* graphics.go -> al_create_custom_bitmap | al_register_bitmap_loader | al_register_bitmap_saver | al_register_bitmap_loader_f | al_register_bitmap_saver_f

Quirks?
=======
* display.go -> SetIcons()
* file.go -> MakeTempFile()

Macros not implemented
======================
* events.h -> ALLEGRO_GET_EVENT_TYPE
* events.h -> ALLEGRO_EVENT_TYPE_IS_USER
* timer.h -> ALLEGRO_USECS_TO_SECS
* timer.h -> ALLEGRO_MSECS_TO_SECS
* timer.h -> ALLEGRO_BPS_TO_SECS
* timer.h -> ALLEGRO_BPM_TO_SECS

Not going to be implemented
===========================
* file.go -> al_fget_ustr -> See utf8.
* memory -> No need. Go is garbage collected.
* path -> Ne need. Go has "path" and "filepath" packages.
* utf8 -> Go supports UTF8. Do we need this?
* threads -> Ne need. Go has goroutines.
* time -> Ne need. Go has "time" package.
* misc -> Not applicable.
* platform.ios platform.android -> Not fully supported by go.
* main addon -> Not applicable.
* native dialogs addon -> ?!?
