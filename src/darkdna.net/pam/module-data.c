#include "_cgo_export.h"
#include "security/pam_appl.h"
#include "stdlib.h"
#include "stdio.h"
#include "stdint.h"

void _module_data_clear(pam_handle_t* hdl, void* ptr, int err_code) {
  uint64_t idx = (uint64_t)ptr;

  clearModuleData(hdl, idx);
}

uint64_t module_data_get(pam_handle_t* hdl, const char* str) {
  uint64_t ret;
  pam_get_data(hdl, str, (const void**)&ret);

  return ret;
}

void module_data_set(pam_handle_t* hdl, const char* str, uint64_t idx) {
  pam_set_data(hdl, str, (void*)idx, _module_data_clear);
}
