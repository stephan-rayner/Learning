#include <iostream>
#include <vector>
#include "libs/app.h"

int main() {
    int len;
    std::cout << "Give me a number plz: ";
    std::cin >> len;
    std::vector<int> c;
    c.resize(len);
    run(c);
    return 0;
}