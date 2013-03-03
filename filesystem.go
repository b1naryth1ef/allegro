package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

const (
	FilemodeRead = C.ALLEGRO_FILEMODE_READ
	FilemodeWrite = C.ALLEGRO_FILEMODE_WRITE
	FilemodeExecute = C.ALLEGRO_FILEMODE_EXECUTE
	FilemodeHidden = C.ALLEGRO_FILEMODE_HIDDEN
	FilemodeIsfile = C.ALLEGRO_FILEMODE_ISFILE
	FilemodeIsdir = C.ALLEGRO_FILEMODE_ISDIR
)

type FsEntry C.ALLEGRO_FS_ENTRY

func CreateFsEntry(path string) *FsEntry {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))
	return (*FsEntry)(unsafe.Pointer(C.al_create_fs_entry(p)))
}

func (f *FsEntry) Destroy() {
	C.al_destroy_fs_entry((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f)))
}

func (f *FsEntry) GetName() string {
	return C.GoString(C.al_get_fs_entry_name((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f))))
}

func (f *FsEntry) Update() bool {
	return bool(C.al_update_fs_entry((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f))))
}

func (f *FsEntry) GetMode() uint32 {
	return uint32(C.al_get_fs_entry_mode((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f))))
}

func (f *FsEntry) GetAtime() int64 {
	return int64(C.al_get_fs_entry_atime((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f))))
}

func (f *FsEntry) GetCtime() int64 {
	return int64(C.al_get_fs_entry_ctime((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f))))
}

func (f *FsEntry) GetMtime() int64 {
	return int64(C.al_get_fs_entry_mtime((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f))))
}

func (f *FsEntry) GetSize() int64 {
	return int64(C.al_get_fs_entry_size((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f))))
}

func (f *FsEntry) Exists() bool {
	return bool(C.al_fs_entry_exists((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f))))
}

func (f *FsEntry) Remove() bool {
	return bool(C.al_remove_fs_entry((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f))))
}

func FilenameExists(path string) bool {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))
	return bool(C.al_filename_exists(p))
}

func RemoveFilename(path string) bool {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))
	return bool(C.al_filename_exists(p))
}

/***********************/
/* Directory Functions */
/***********************/

func (f *FsEntry) OpenDirectory() bool {
	return bool(C.al_open_directory((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f))))
}

func (f *FsEntry) ReadDirectory() *FsEntry {
	return (*FsEntry)(unsafe.Pointer(C.al_read_directory((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f)))))
}

func (f *FsEntry) CloseDirectory() bool {
	return bool(C.al_close_directory((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f))))
}

func GetCurrentDirectory() string {
	return C.GoString(C.al_get_current_directory())
}

func ChangeDirectory(path string) bool {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))
	return bool(C.al_change_directory(p))
}

func MakeDirectory(path string) bool {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))	
	return bool(C.al_make_directory(p))
}

func (f *FsEntry) Open(mode string) *File {
	m := C.CString(mode)
	defer C.free(unsafe.Pointer(m))
	return (*File)(unsafe.Pointer(C.al_open_fs_entry((*C.ALLEGRO_FS_ENTRY)(unsafe.Pointer(f)), m)))
}
