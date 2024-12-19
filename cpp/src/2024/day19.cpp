#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <regex>
#include <string>
#include <vector>

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::getline(file, line);

    std::string prefix = "^((";
    std::string postfix = "))+$";
    std::string from = ", ";
    std::string to = ")|(";

    size_t pos = 0;
    while ((pos = line.find(from, pos)) != std::string::npos) {
        line.replace(pos, from.length(), to);
        pos += to.length();
    }

    std::regex pattern(prefix + line + postfix);

    int n_possible = 0;
    std::getline(file, line); // ignore blank line
    while (std::getline(file, line)) {
        // std::cout << line << ": ";
        if (std::regex_match(line, pattern)) {
            // std::cout << "match!\n";
            n_possible++;
        } else {
            // std::cout << "no match.\n";
        }
    }
    std::cout << "Part 1: " << n_possible << std::endl;
}

void Solution::part2() {}