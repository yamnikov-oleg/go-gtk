// +build !cgocheck

package gio

// #include "gio.go.h"
// #cgo pkg-config: gio-2.0
import "C"
import (
	"unsafe"

	"github.com/yamnikov-oleg/go-gtk/glib"
)

type Icon struct {
	GIcon *C.GIcon
}

func NewIconForString(str string) (*Icon, *glib.Error) {
	var err *C.GError
	ptr := C.CString(str)
	defer C.free(unsafe.Pointer(ptr))
	gicon := C.g_icon_new_for_string((*C.gchar)(ptr), &err)
	if err != nil {
		return nil, glib.ErrorFromNative(unsafe.Pointer(err))
	}
	return &Icon{gicon}, nil
}

func GetIconType() int {
	return int(C.g_icon_get_type())
}
