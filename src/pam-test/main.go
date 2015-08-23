package main

// #cgo LDFLAGS: -lpam
// #include <security/pam_modules.h>
// #include <security/pam_appl.h>
import "C"

import (
	"fmt"
	"time"
	"unsafe"

	"darkdna.net/pam"
)

var sessionBegin = C.CString("session-begin")

//export goAuthenticate
func goAuthenticate(handle *C.pam_handle_t, flags C.int, argv []string) C.int {
	hdl := pam.Handle{unsafe.Pointer(handle)}
	fmt.Println("authenticate:", argv)

	usr, err := hdl.GetUser()
	if err != nil {
		return C.PAM_AUTH_ERR
	}

	fmt.Println("User:", usr)
	if usr != "test" {
		return C.PAM_USER_UNKNOWN
	}

	resps, err := hdl.Conversation(
		pam.Message{
			Style: pam.MessageEchoOff,
			Msg:   "Password: ",
		},
	)

	if err != nil {
		fmt.Println("Error: ", err)
		return C.PAM_CONV_ERR
	}

	if resps[0] != "cake" {
		return C.PAM_AUTH_ERR
	}

	resps, err = hdl.Conversation(
		pam.Message{
			Style: pam.MessageEchoOn,
			Msg:   "Favourite colour: ",
		})

	if err != nil {
		fmt.Println("Error: ", err)
		return C.PAM_CONV_ERR
	}

	fmt.Println("I can't believe you like the colour", resps[0])
	hdl.SetModuleData("fav-colour", resps[0])

	return C.PAM_SUCCESS
}

//export setCred
func setCred(handle *C.pam_handle_t, flags C.int, argv []string) C.int {
	fmt.Println("setcred: ", argv)

	return C.PAM_SUCCESS
}

//export openSession
func openSession(handle *C.pam_handle_t, flags C.int, argv []string) C.int {
	hdl := pam.Handle{unsafe.Pointer(handle)}

	fmt.Println("open_session: ", argv)
	hdl.SetModuleData("session-begin", time.Now())

	return C.PAM_SUCCESS
}

//export closeSession
func closeSession(handle *C.pam_handle_t, flags C.int, argv []string) C.int {
	fmt.Println("close_session: ", argv)
	hdl := pam.Handle{unsafe.Pointer(handle)}

	tmp, err := hdl.GetModuleData("session-begin")
	if err == nil {
		signIn := tmp.(time.Time)

		fmt.Println("User was logged in for ", time.Now().Sub(signIn))
	} else {
		fmt.Println("User data error: ", err)
	}

	tmp, err = hdl.GetModuleData("fav-colour")
	if err == nil {
		favColour := tmp.(string)
		fmt.Println("Still can't believe their favourite colour is", favColour)
	} else {
		fmt.Println("Test failed, no colour data.")
	}

	return C.PAM_SUCCESS
}

// Does Nothing?
func main() {}
