package main

// #cgo LDFLAGS: -lpam
// #define PAM_SM_AUTH
// #include <security/pam_modules.h>
import "C"
import (
	"fmt"
	"strings"
)

//export Authenticate
func Authenticate(handle *C.pam_handle_t, flags C.int, argv []string) C.int {
	fmt.Printf("Argv: %+v\n", argv)

	var user *C.char
	user = C.CString(strings.Repeat(" ", 54))

	if C.pam_get_user(handle, &user, nil) != C.PAM_SUCCESS {
		return C.PAM_USER_UNKNOWN
	}

	fmt.Printf("User: %+v\n", C.GoString(user))
	if C.GoString(user) != "test" {
		return C.PAM_USER_UNKNOWN
	}

	return C.PAM_SUCCESS
}

// Does Nothing?
func main() {

}
