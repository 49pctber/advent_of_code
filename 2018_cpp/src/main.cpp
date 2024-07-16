#include "aoc.hpp"
#include <iostream>

int main(int argc, char **argv) {
    if (argc == 1) {
        std::cout << "Welcome to the Advent of Code!\nSpecify the day using "
                     "the first positional command line argument.\n";
    } else {
        int day = std::atoi(argv[1]);
        switch (day) {
        case 1:
            day1();
            break;
        case 2:
            day2();
            break;
        default:
            std::cout << "Invalid day specified.\n";
            break;
        }
    }

    return 0;
}