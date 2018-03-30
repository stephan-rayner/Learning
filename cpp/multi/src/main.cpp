#include <iostream>
#include <vector>
#include <sstream> // for std::stringstream
#include "libs/app.h"

int main(int argc, char *argv[]) {
    int len;

    if (argc > 1) {
        std::stringstream geek(argv[1]);
        std::cout << "argc = " << argc << std::endl;
        std::cout << argv[1] << std::endl;
        geek >> len;

    } else {
        std::cout << "Give me a number plz: ";
        std::cin >> len;
    }

    std::vector<int> c;
    c.resize(len);
    run(c);
    exit(0);
}