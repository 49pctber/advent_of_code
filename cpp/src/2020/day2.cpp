#include "solution.hpp"
#include <fstream>
#include <regex>
#include <string>

void Solution::part1() {
    std::ifstream input(argv[1]);
    if (!input.is_open()) {
        std::cerr << "Failed to open input file" << std::endl;
    }

    std::string line;
    int n_valid_passwords = 0;
    while (std::getline(input, line)) {
        std::regex pattern(R"((\d+)-(\d+) (\w): (\w+))");
        std::smatch matches;
        if (std::regex_search(line, matches, pattern)) {
            int min = std::stoi(matches[1]);
            int max = std::stoi(matches[2]);
            char c = std::string(matches[3])[0];
            std::string password = matches[4];
            int count = 0;
            for (char d : password) {
                if (d == c) {
                    count++;
                }
            }
            if (count <= max && count >= min) {
                n_valid_passwords++;
            }
        } else {
            std::cout << "No matches found for " << line << '\n';
        }
    }
    input.close();
    std::cout << "Part 1: " << n_valid_passwords << std::endl;
}

void Solution::part2() {
    std::ifstream input(argv[1]);
    if (!input.is_open()) {
        std::cerr << "Failed to open input file" << std::endl;
    }

    std::string line;
    int n_valid_passwords = 0;
    while (std::getline(input, line)) {
        std::regex pattern(R"((\d+)-(\d+) (\w): (\w+))");
        std::smatch matches;
        if (std::regex_search(line, matches, pattern)) {
            int pos1 = std::stoi(matches[1]) - 1;
            int pos2 = std::stoi(matches[2]) - 1;
            char c = std::string(matches[3])[0];
            std::string password = matches[4];
            if (password[pos1] == c ^ password[pos2] == c) {
                n_valid_passwords++;
            }
        } else {
            std::cout << "No matches found for " << line << '\n';
        }
    }
    input.close();
    std::cout << "Part 2: " << n_valid_passwords << std::endl;
}
