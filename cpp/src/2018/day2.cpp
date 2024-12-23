#include "solution.hpp"
#include <map>
#include <vector>

void Solution::part1() {
    int double_count = 0;
    int triple_count = 0;

    std::ifstream f(argv[1]);
    std::string line;

    if (!f.is_open()) {
        std::cerr << "error opening file\n";
        exit(-1);
    }

    while (std::getline(f, line)) {

        std::map<char, int> counts;
        for (char c : line) {
            counts[c]++;
        }

        for (std::pair<char, int> p : counts) {
            if (p.second == 2) {
                double_count++;
                break;
            }
        }

        for (std::pair<char, int> p : counts) {
            if (p.second == 3) {
                triple_count++;
                break;
            }
        }
    }

    std::cout << "Part 1: " << double_count * triple_count << std::endl;
};

void Solution::part2() {
    std::ifstream f(argv[1]);
    std::string line;

    if (!f.is_open()) {
        std::cerr << "error opening file\n";
        exit(-1);
    }
    std::vector<std::string> prev;

    while (std::getline(f, line)) {

        // compare against previous strings
        for (std::string p : prev) {

            // check if line only differs in one spot
            int idx = -1;
            bool found = false;
            for (int i = 0; i < line.length(); i++) {
                if (p[i] == line[i]) {
                    continue;
                } else if (idx == -1) {
                    idx = i;
                    found = true;
                } else {
                    found = false;
                    break;
                }
            }
            if (found) {
                std::cout << "Part 2: " << line.erase(idx, 1) << std::endl;
                return;
            }
        }
        prev.push_back(line);
    }

    std::cout << "not found\n";
    exit(-1);
};
