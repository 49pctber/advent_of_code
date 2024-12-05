#include "solution.hpp"
#include <fstream>
#include <iostream>

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }
}

void Solution::part2() {}