This is a (Go)[https://golang.org] implementation of a simple PAM authentication
module. It returns `PAM_SUCCESS` for "test" and `PAM_USER_UNKNOWN` for everyone else.

Made this as a proof-of-concept for playing with the new `-buildmode=c-shared`
feature in Go 1.5.

To build this project you need (wgo)[https://github.com/skelterjohn/wgo].
