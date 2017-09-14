This is a [Go](https://golang.org) implementation of a simple PAM authentication
module. It returns `PAM_SUCCESS` for "test" and `PAM_USER_UNKNOWN` for everyone else.

Made this as a proof-of-concept for playing with the new `-buildmode=c-shared`
feature in Go 1.5.

This should not be used without review by someone with more experience with PAM / cgo, 
as I only did this as a minimal toy.

To build this project you need [wgo](https://github.com/skelterjohn/wgo).
