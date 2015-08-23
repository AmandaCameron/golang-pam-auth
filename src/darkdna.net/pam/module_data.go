package pam

/*
#cgo LDFLAGS: -lpam
#include "stdlib.h"
#include "stdint.h"
#include "security/pam_appl.h"

extern void module_data_set(pam_handle_t*, const char*, uint64_t);
extern uint64_t module_data_get(pam_handle_t*, const char*);
*/
import "C"
import "errors"

import "unsafe"

var dataMap map[C.uint64_t]interface{}
var dataIdx C.uint64_t

func init() {
	dataMap = make(map[C.uint64_t]interface{})
	dataIdx = C.uint64_t(0)
}

// SetModuleData sets the speciified module data.
func (hdl Handle) SetModuleData(name string, data interface{}) error {
	nameBuf := C.CString(name)
	defer C.free(unsafe.Pointer(nameBuf))

	dataIdx++
	idx := dataIdx

	C.module_data_set((*C.pam_handle_t)(hdl.Ptr), nameBuf, idx)
	dataMap[idx] = data

	return nil
}

// GetModuleData gets the specified module data.
func (hdl Handle) GetModuleData(name string) (interface{}, error) {
	nameBuf := C.CString(name)
	defer C.free(unsafe.Pointer(nameBuf))

	idx := C.module_data_get((*C.pam_handle_t)(hdl.Ptr), nameBuf)
	if data, ok := dataMap[idx]; ok {
		return data, nil
	}

	return nil, errors.New("No such data.")
}

//export clearModuleData
func clearModuleData(hdl *C.pam_handle_t, idx C.uint64_t) {
	dataMap[idx] = nil
}
