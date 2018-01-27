#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int x, y;
} inputs;


inputs make_point(int x, int y) {
    inputs data;
    data.x = x;
    data.y = y;
    return data;
}


int dbl_int(int x){
    return x * 2;
}

char* hello(void) {
    char* txt = "Hello";
    return txt;
}

