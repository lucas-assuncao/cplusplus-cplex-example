#include "lib_interface.h"
#include "model.h"

const char * callLib(const char *input)
{
  return buildAndRunModel(input);
}