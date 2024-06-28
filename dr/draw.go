package dr

/*
#include draw.c
*/

import "C"

func dr() {
	C.hello()
}

// func MyDrawText(text string, posX float32, posY float32, fontSize int32, col color.RGBA) {
// 	ctext := C.CString(text)
// 	defer C.free(unsafe.Pointer(ctext))
// 	cposX := (C.float)(posX)
// 	cposY := (C.float)(posY)
// 	cfontSize := (C.int)(fontSize)
// 	ccolor := colorCptr(col)
// 	C.DrawText(ctext, cposX, cposY, cfontSize, *ccolor)
// }
