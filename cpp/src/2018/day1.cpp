#include "solution.hpp"
#include <set>
#include <vector>

void Solution::part1() {
    std::ifstream f(argv[1]);

    std::string line;
    if (!f.is_open()) {
        std::cerr << "error opening file\n";
        exit(-1);
    }

    int current_frequency = 0;

    while (std::getline(f, line)) {
        int delta = std::stoi(line);
        current_frequency += delta;
    }

    std::cout << "Part 1: " << current_frequency << std::endl;
}

void Solution::part2() {
    std::ifstream f(argv[1]);

    std::string line;
    if (!f.is_open()) {
        std::cerr << "error opening file\n";
        exit(-1);
    }

    int current_frequency = 0;
    std::set<int> seen;
    seen.insert(current_frequency);
    std::vector<int> list;

    while (std::getline(f, line)) {
        int delta = std::stoi(line);
        list.push_back(delta);
    }

    int i = 0;
    while (true) {
        current_frequency += list[i];
        if (seen.find(current_frequency) != seen.end()) {
            break;
        } else {
            i += 1;
            i %= list.size();
            seen.insert(current_frequency);
        }
    }

    std::cout << "Part 2: " << current_frequency << std::endl;
}