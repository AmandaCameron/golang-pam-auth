package pam

/*
#cgo LDFLAGS: -lpam
#include <stdlib.h>
#include <security/pam_appl.h>

int do_conv(pam_handle_t* hdlr, int count, const struct pam_message** msgs, struct pam_response** responses) {
	int err;
	struct pam_conv* conv;

	err = pam_get_item(hdlr, PAM_CONV, (const void**)&conv);
	if(err != PAM_SUCCESS) {
		return err;
	}

	return conv->conv(count, msgs, responses, conv->appdata_ptr);
}
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

// MessageStyle is a style of Message
type MessageStyle int

const (
	// MessageEchoOff is for messages that shouldn't gave an echo.
	MessageEchoOff = MessageStyle(C.PAM_PROMPT_ECHO_OFF)

	// MessageEchoOn is for messages that should have an echo.
	MessageEchoOn = MessageStyle(C.PAM_PROMPT_ECHO_ON)

	// MessageErrorMsg is for messages that should be displayed as an error.
	MessageErrorMsg = MessageStyle(C.PAM_ERROR_MSG)

	// MessageTextInfo is for textual blurbs to be spat out.
	MessageTextInfo = MessageStyle(C.PAM_TEXT_INFO)
)

// Message represents something to ask / show in a Conv.Conversation call.
type Message struct {
	Style MessageStyle
	Msg   string
}

// Conversation passes on the specified messages.
func (hdl Handle) Conversation(_msgs ...Message) ([]string, error) {
	if len(_msgs) == 0 {
		return nil, errors.New("Must pass at least one Message.")
	}

	msg := []*C.struct_pam_message{}
	resp := []*C.struct_pam_response{}

	for _, _msg := range _msgs {
		msgStruct := &C.struct_pam_message{msg_style: C.int(_msg.Style), msg: C.CString(_msg.Msg)}
		defer C.free(unsafe.Pointer(msgStruct.msg))

		msg = append(msg, msgStruct)
		resp = append(resp, &C.struct_pam_response{})
	}

	code := C.do_conv(hdl.ptr(), C.int(len(_msgs)), &msg[0], &resp[0])
	if code != C.PAM_SUCCESS {
		return nil, fmt.Errorf("Got non-success from the function: %d", code)
	}

	var ret []string
	for _, r := range resp {
		ret = append(ret, C.GoString(r.resp))
		C.free(unsafe.Pointer(r.resp))
	}

	return ret, nil
}
