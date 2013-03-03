package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

func Fopen(path, mode string) *File {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))
	m := C.CString(mode)
	defer C.free(unsafe.Pointer(m))
	return (*File)(unsafe.Pointer(C.al_fopen(p, m)))
}

func (f *File) FopenSlice(initialSize uint32, mode string) *File {
	m := C.CString(mode)
	defer C.free(unsafe.Pointer(m))
	return (*File)(unsafe.Pointer(C.al_fopen_slice((*C.ALLEGRO_FILE)(unsafe.Pointer(f)), C.size_t(initialSize), m)))
}

func (f *File) Fclose() {
	C.al_fclose((*C.ALLEGRO_FILE)(unsafe.Pointer(f)))
}

func (f *File) Fread(ptr unsafe.Pointer, size uint32) uint32 {
	return uint32(C.al_fread((*C.ALLEGRO_FILE)(unsafe.Pointer(f)), ptr, C.size_t(size)))
}

func (f *File) Fwrite(ptr unsafe.Pointer, size uint32) uint32 {
	return uint32(C.al_fwrite((*C.ALLEGRO_FILE)(unsafe.Pointer(f)), ptr, C.size_t(size)))
}

func (f *File) Fflush() bool {
	return bool(C.al_fflush((*C.ALLEGRO_FILE)(unsafe.Pointer(f))))
}

func (f *File) Ftell() int64 {
	return int64(C.al_ftell((*C.ALLEGRO_FILE)(unsafe.Pointer(f))))
}

func (f *File) Fseek(offset int64, whence int32) bool {
	return bool(C.al_fseek((*C.ALLEGRO_FILE)(unsafe.Pointer(f)), C.int64_t(offset), C.int(whence)))
}

func (f *File) Feof() bool {
	return bool(C.al_feof((*C.ALLEGRO_FILE)(unsafe.Pointer(f))))
}

func (f *File) Ferror() bool {
	return bool(C.al_ferror((*C.ALLEGRO_FILE)(unsafe.Pointer(f))))
}

func (f *File) Fclearerr() {
	C.al_fclearerr((*C.ALLEGRO_FILE)(unsafe.Pointer(f)))
}

func (f *File) Fungetc(c int32) int32 {
	return int32(C.al_fungetc((*C.ALLEGRO_FILE)(unsafe.Pointer(f)), C.int(c)))
}

func (f *File) Fsize() int64 {
	return int64(C.al_fsize((*C.ALLEGRO_FILE)(unsafe.Pointer(f))))
}

func (f *File) Fgetc() int32 {
	return int32(C.al_fgetc((*C.ALLEGRO_FILE)(unsafe.Pointer(f))))
}

func (f *File) Fputc(c int32) int32 {
	return int32(C.al_fputc((*C.ALLEGRO_FILE)(unsafe.Pointer(f)), C.int(c)))
}

func (f *File) Fread16le() int16 {
	return int16(C.al_fread16le((*C.ALLEGRO_FILE)(unsafe.Pointer(f))))
}

func (f *File) Fread16be() int16 {
	return int16(C.al_fread16be((*C.ALLEGRO_FILE)(unsafe.Pointer(f))))
}

func (f *File) Fwrite16le(w int16) int16 {
	return int16(C.al_fwrite16le((*C.ALLEGRO_FILE)(unsafe.Pointer(f)), C.int16_t(w)))
}

func (f *File) Fwrite16be(w int16) int16 {
	return int16(C.al_fwrite16be((*C.ALLEGRO_FILE)(unsafe.Pointer(f)), C.int16_t(w)))
}

func (f *File) Fread32le() int32 {
	return int32(C.al_fread32le((*C.ALLEGRO_FILE)(unsafe.Pointer(f))))
}

func (f *File) Fread32be() int32 {
	return int32(C.al_fread32be((*C.ALLEGRO_FILE)(unsafe.Pointer(f))))
}

func (f *File) Fwrite32le(w int32) int32 {
	return int32(C.al_fwrite32le((*C.ALLEGRO_FILE)(unsafe.Pointer(f)), C.int32_t(w)))
}

func (f *File) Fwrite32be(w int32) int32 {
	return int32(C.al_fwrite32be((*C.ALLEGRO_FILE)(unsafe.Pointer(f)), C.int32_t(w)))
}

func (f *File) Fgets(buf string, max uint32) string {
	b := C.CString(buf)
	defer C.free(unsafe.Pointer(b))
	return C.GoString(C.al_fgets((*C.ALLEGRO_FILE)(unsafe.Pointer(f)), b, C.size_t(max)))
}

func (f *File) Fputs(p string) int32 {
	pp := C.CString(p)
	defer C.free(unsafe.Pointer(pp))
	return int32(C.al_fputs((*C.ALLEGRO_FILE)(unsafe.Pointer(f)), pp))
}

func FopenFd(fd int32, mode string) *File {
	m := C.CString(mode)
	defer C.free(unsafe.Pointer(m))
	return (*File)(unsafe.Pointer(C.al_fopen_fd(C.int(fd), m)))
}

//func MakeTempFile(template string, retPath **Path) *File {
//	t := C.CString(template)
//	defer C.free(unsafe.Pointer(t))
//	return (*File)(unsafe.Pointer(C.al_make_temp_file(t, (*C.ALLEGRO_PATH)(unsafe.Pointer(retPath)))))
//}
