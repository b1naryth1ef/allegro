package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

type MonitorInfo struct {
	X1 int32
	Y1 int32
	X2 int32
	Y2 int32
}

func GetNewDisplayAdapter() int32 {
	return int32(C.al_get_new_display_adapter())
}

func SetNewDisplayAdapter(adapter int32) {
	C.al_set_new_display_adapter(C.int(adapter))
}

func GetMonitorInfo(adapter int32) *MonitorInfo {
	var m C.ALLEGRO_MONITOR_INFO
	b := bool(C.al_get_monitor_info(C.int(adapter), &m))
	if b == false {
		return nil
	}
	return (*MonitorInfo)(unsafe.Pointer(&m))
}

func GetNumVideoAdapters() int32 {
	return int32(C.al_get_num_video_adapters())
}
