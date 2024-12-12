#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <string>

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    while (std::getline(file, line)) {
        std::cout << line << std::endl;
    }
}

void Solution::part2() {}