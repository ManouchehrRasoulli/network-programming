#include <stdlib.h>
#include "req.h"

int main(int argc, char ** args) {
    char * buff = new char[512];

    Req * r = new Req();
    r->op = OpCode::Read;
    r->Encode(buff);

    std::cout << buff << std::endl;

    return EXIT_SUCCESS;
}