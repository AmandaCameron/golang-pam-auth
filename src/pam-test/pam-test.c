#include "_cgo_export.h"
#include "string.h"

#define PAM_SM_AUTH
#include <security/pam_modules.h>

GoSlice argcvToSlice(int, const char**);

PAM_EXTERN int pam_sm_authenticate(pam_handle_t* pamh, int flags, int argc, const char** argv) {
  GoSlice args = argcvToSlice(argc, argv);

	return Authenticate(pamh, flags, args);
}

GoSlice argcvToSlice(int argc, const char** argv) {
  GoString* strs = malloc(sizeof(GoString) * argc);

  GoSlice ret;
  ret.cap = argc;
  ret.len = argc;
  ret.data = (void*)strs;

  int i;
  for(i = 0; i < argc; i++) {
    strs[i] = *((GoString*)malloc(sizeof(GoString)));

    strs[i].p = (char*)argv[i];
    strs[i].n = strlen(argv[i]);
  }

  return ret;
}
