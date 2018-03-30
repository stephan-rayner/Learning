#include "app.h"
#include <iostream>
#include <vector>

void run(std::vector<int> c) {
    for (int i = 1; i <= c.size(); i++) {
        c[i] = i * i;
        std::cout << i << " * " << i << " = " << c[i] << " | ";
        if ( i % 3 == 0) {
            std::cout << std::endl;
        }
    }
    std::cout << std::endl;
}