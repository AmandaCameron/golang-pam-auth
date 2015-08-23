package pam

/*
#cgo LDFLAGS: -lpam
#include <security/pam_appl.h>
#include <security/pam_modules.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Handle is a handle type to hang the PAM methods off of.
type Handle struct {
	Ptr unsafe.Pointer
}

func (hdl Handle) ptr() *C.pam_handle_t {
	return (*C.pam_handle_t)(hdl.Ptr)
}

// GetUser maps to the pam_get_user call, and returns the user that we're trying to auth as.
func (hdl Handle) GetUser() (string, error) {
	var usr *C.char
	if err := C.pam_get_user(hdl.ptr(), &usr, nil); err != C.PAM_SUCCESS {
		return "", pamError(err)
	}

	return C.GoString(usr), nil
}

type pamError C.int

func (pe pamError) Error() string {
	return fmt.Sprintf("PAM error code %d", pe)
}
