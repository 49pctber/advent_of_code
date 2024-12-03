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
    std::regex pattern(R"(mul\(([0-9]{1,3}),([0-9]{1,3})\))");
    std::regex pattern_do(R"(do\(\))");
    std::regex pattern_dont(R"(don't\(\))");

    std::vector<int> donts;
    std::vector<int> dos;
    int offset = 0;
    while (std::getline(file, line)) {
        // std::cout << line << std::endl;
        auto begin_dont =
            std::sregex_iterator(line.begin(), line.end(), pattern_dont);
        auto end_dont = std::sregex_iterator();

        for (auto it = begin_dont; it != end_dont; ++it) {
            std::smatch match = *it;
            std::cout << match.position() << " ";
            donts.push_back(match.position() + offset);
        }

        auto begin_do =
            std::sregex_iterator(line.begin(), line.end(), pattern_do);
        auto end_do = std::sregex_iterator();

        for (auto it = begin_do; it != end_do; ++it) {
            std::smatch match = *it;
            // std::cout << match.position() << " ";
            dos.push_back(match.position() + offset);
        }

        auto begin = std::sregex_iterator(line.begin(), line.end(), pattern);
        auto end = std::sregex_iterator();
        for (auto it = begin; it != end; ++it) {
            std::smatch match = *it;
            int pos = match.position() + offset;
            int dontpos = -2;
            int dopos = -1;

            for (int n : donts) {
                if (n < pos) {
                    dontpos = n;
                }
            }
            for (int n : dos) {
                if (n < pos) {
                    dopos = n;
                }
            }
            if (dopos > dontpos) {
                sum += std::stoi(match[1], NULL) * std::stoi(match[2], NULL);
            }
        }
        offset += line.size();
    }

    std::cout << "Part 2: " << sum << std::endl; // not 66677244
}