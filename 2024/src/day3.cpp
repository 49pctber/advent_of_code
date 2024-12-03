#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <regex>
#include <string>

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cout << "Error opening file" << std::endl;
        return;
    }

    std::string line;
    int sum = 0;
    std::regex pattern(R"(mul\(([0-9]+),([0-9]+)\))");
    while (std::getline(file, line)) {
        auto begin = std::sregex_iterator(line.begin(), line.end(), pattern);
        auto end = std::sregex_iterator();
        for (auto it = begin; it != end; ++it) {
            std::smatch match = *it;
            sum += std::stoi(match[1], NULL) * std::stoi(match[2], NULL);
        }
    }
    file.close();

    std::cout << "Part 1: " << sum << std::endl;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cout << "Error opening file" << std::endl;
        return;
    }

    std::string line;
    int sum = 0;
    std::regex pattern(R"(mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\))");
    bool ignore = false;

    while (std::getline(file, line)) {

        auto begin = std::sregex_iterator(line.begin(), line.end(), pattern);
        auto end = std::sregex_iterator();
        for (auto it = begin; it != end; ++it) {
            std::smatch match = *it;

            if (match.str() == "do()") {
                ignore = false;
            } else if (match.str() == "don't()") {
                ignore = true;
            } else {
                if (!ignore) {
                    sum +=
                        std::stoi(match[1], NULL) * std::stoi(match[2], NULL);
                }
            }
        }
    }

    std::cout << "Part 2: " << sum << std::endl;
}