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
    long int fuel = 0;
    while (std::getline(file, line)) {
        int mass = std::stoi(line);
        fuel += (mass / 3) - 2;
    }
    std::cout << "Part 1: " << fuel << std::endl;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    long int fuel = 0;
    while (std::getline(file, line)) {
        int mass = std::stoi(line);
        int next = (mass / 3) - 2;
        while (next > 0) {
            fuel += next;
            next = (next / 3) - 2;
        }
    }
    std::cout << "Part 2: " << fuel << std::endl;
}
