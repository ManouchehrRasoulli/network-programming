#include <stdio.h>
#include <iostream>
#include <string.h>

#ifndef REQ
#define REQ

enum class OpCode {
    Read,
    Write,
    Ack,
};

class Req {
    public:
        void Encode(char * buff) {
            switch (this->op)
            {
            case OpCode::Read:
                {
                    int iter = 0;
                    const char* value1 = "myvalue";
                    memcpy(buff+iter, value1, strlen(value1));
                    iter+=strlen(value1);

                    double value2 = 0.1;
                    memcpy(buff+iter, &value2, sizeof(double));
                }
                break;
            default:
                {

                }
                break;
            }
        };
        void Decode(char * buff) {

        };
        OpCode op; // specify the request operation which we wish server perform for us
};

#endif