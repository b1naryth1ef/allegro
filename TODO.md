TODO List
=========
* Implement addons: video
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
* direct3d.go -> al_d3d_set_release_callback | al_d3d_set_restore_callback
* audio.go -> al_set_mixer_postprocess_callback | al_register_sample_loader | al_register_sample_loader_f | al_register_sample_saver | al_register_sample_saver_f | al_register_audio_stream_loader | al_register_audio_stream_loader_f
* font.go -> al_register_font_loader
* primitives -> al_calculate_spline | al_calculate_ribbon | al_calculate_arc | all low-level drawing routines

Quirks?
=======
* display.go -> SetIcons()
* file.go -> MakeTempFile()

Macros not implemented
======================
* events.h -> ALLEGRO_GET_EVENT_TYPE | ALLEGRO_EVENT_TYPE_IS_USER
* timer.h -> ALLEGRO_USECS_TO_SECS | ALLEGRO_MSECS_TO_SECS | ALLEGRO_BPS_TO_SECS | ALLEGRO_BPM_TO_SECS

Not going to be implemented
===========================
* file.go -> al_fget_ustr -> See utf8.
* memory -> No need. Go is garbage collected.
* path -> No need. Go has "path" and "filepath" packages.
* utf8 -> Go supports UTF8. Do we need this?
* threads -> No need. Go has goroutines.
* misc -> Not applicable.
* platform.ios platform.android -> Not fully supported by Go.
* main addon -> Not applicable.
* native dialogs addon -> ?!?
* al_get_opengl_extension_list | al_get_opengl_proc_address | al_get_d3d_device | al_get_d3d_system_texture | al_get_d3d_video_texture
* font.go -> al_get_ustr_width | al_draw_ustr | al_draw_justified_ustr | al_get_ustr_dimensions -> See utf8.
