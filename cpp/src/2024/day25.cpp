#include "solution.hpp"
#include <array>
#include <fstream>
#include <iostream>
#include <set>
#include <string>

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::set<std::array<int, 5>> locks;
    std::set<std::array<int, 5>> keys;

    std::string line;
    while (std::getline(file, line)) {
        bool lock = line == "#####";

        std::array<int, 5> pins{0, 0, 0, 0, 0};

        for (int i = 0; i < 5; i++) {
            std::getline(file, line);
            for (int c = 0; c < 5; c++) {
                if (line[c] == '#') {
                    pins[c]++;
                }
            }
        }
        std::getline(file, line);
        std::getline(file, line);

        if (lock) {
            locks.insert(pins);
        } else {
            keys.insert(pins);
        }
    }

    // std::cout << "Locks: " << locks.size() << "\n";
    // std::cout << "Keys:  " << keys.size() << "\n";

    int count = 0;
    for (auto lock : locks) {
        for (auto key : keys) {
            bool fits = true;
            for (int i = 0; i < 5; i++) {
                if (lock[i] + key[i] > 5) {
                    fits = false;
                }
            }
            if (fits) {
                count++;
            }
        }
    }
    std::cout << "Part 1: " << count << std::endl;
}

void Solution::part2() {}