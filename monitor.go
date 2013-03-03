package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

func GetNewDisplayAdapter() int32 {
	return int32(C.al_get_new_display_adapter())
}

func SetNewDisplayAdapter(adapter int32) {
	C.al_set_new_display_adapter(C.int(adapter))
}

func GetMonitorInfo(adapter int32) (bool, *MonitorInfo) {
	var m C.ALLEGRO_MONITOR_INFO
	b := bool(C.al_get_monitor_info(C.int(adapter), &m))
	return b, (*MonitorInfo)(unsafe.Pointer(&m))
}

func GetNumVideoAdapters() int32 {
	return int32(C.al_get_num_video_adapters())
}
