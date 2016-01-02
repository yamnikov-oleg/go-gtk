// +build !cgocheck

package gtkgl

// #include "gtkgl.go.h"
// #cgo CPPFLAGS: -w -static
// #cgo pkg-config: gtkgl-3.0
import "C"

import "github.com/yamnikov-oleg/go-gtk/gtk"
import "unsafe"

type GLArea struct {
	gtk.DrawingArea
}

func NewGLArea(attrList []int) *GLArea {
	ptr := unsafe.Pointer(nil)
	if len(attrList) > 0 {
		ptr = unsafe.Pointer(&attrList[0])
	}
	return &GLArea{gtk.DrawingArea{
		*gtk.WidgetFromNative(C.make_area(C.int(len(attrList)), (*C.int)(ptr)))},
	}
}

func (v *GLArea) MakeCurrent() {
	C.ggla_area_make_current((*C.GglaArea)(unsafe.Pointer(v.GWidget)))
}

func (v *GLArea) SwapBuffers() {
	C.ggla_area_swap_buffers((*C.GglaArea)(unsafe.Pointer(v.GWidget)))
}
