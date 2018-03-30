#include "app.h"
#include <iostream>
#include <vector>

void run(std::vector<int> c) {
    for (int i = 0; i < c.size(); i++)
    {
        c[i] = i * i;
        std::cout << c[i] << std::endl;
    }
}