#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <stdbool.h>
#include <string>
#include <vector>

bool safe_levels(std::vector<int> levels) {

    bool ascending = levels[1] > levels[0];

    if (ascending) {
        for (size_t i = 0; i < levels.size() - 1; i++) {
            if (levels[i + 1] <= levels[i] || levels[i + 1] > levels[i] + 3) {
                return false;
            }
        }
    } else {
        for (size_t i = 0; i < levels.size() - 1; i++) {
            if (levels[i + 1] >= levels[i] || levels[i + 1] < levels[i] - 3) {
                return false;
            }
        }
    }

    return true;
}

std::vector<int> parse_line(std::string line) {
    std::vector<int> levels;
    std::stringstream ss(line);
    int level;
    while (ss >> level) {
        levels.push_back(level);
    }
    return levels;
}

void Solution::part1() {
    std::ifstream input(argv[1]);
    if (!input.is_open()) {
        std::cerr << "Failed to open input file" << std::endl;
    }
    std::string line;

    int safe_count = 0;

    while (std::getline(input, line)) {
        if (safe_levels(parse_line(line))) {
            safe_count++;
        }
    }

    std::cout << "Part 1: " << safe_count << std::endl;
}

void Solution::part2() {
    std::ifstream input(argv[1]);
    if (!input.is_open()) {
        std::cerr << "Failed to open input file" << std::endl;
    }
    std::string line;

    int safe_count = 0;

    while (std::getline(input, line)) {

        auto levels = parse_line(line);

        if (safe_levels(levels)) {
            safe_count++;
            continue;
        }

        for (size_t i = 0; i < levels.size(); i++) {
            std::vector<int> levels2(levels);
            levels2.erase(levels2.begin() + i);

            if (safe_levels(levels2)) {
                safe_count++;
                break;
            }
        }
    }

    std::cout << "Part 2: " << safe_count << std::endl;
}
