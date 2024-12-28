#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    int prev = INT32_MAX;
    int n_inc = 0;
    while (std::getline(file, line)) {
        int curr = std::stoi(line);
        if (curr > prev) {
            n_inc++;
        }
        prev = curr;
    }

    std::cout << "Part 1: " << n_inc << std::endl;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::vector<int> measurements;
    while (std::getline(file, line)) {
        measurements.push_back(std::stoi(line));
    }

    int prev = measurements[0] + measurements[1] + measurements[2];
    int n_inc = 0;
    for (size_t i = 3; i < measurements.size(); i++) {
        int curr = prev + measurements[i] - measurements[i - 3];
        if (curr > prev) {
            n_inc++;
        }
        prev = curr;
    }

    std::cout << "Part 2: " << n_inc << std::endl;
}