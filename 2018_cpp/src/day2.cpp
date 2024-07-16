#include "aoc.hpp"

void day2() {
    std::cout << "Day 2\n";
    std::filesystem::path input_dir(INPUT_DIR);
    std::filesystem::path path = input_dir / "2.txt";
    std::cout << "Part 1: " << day2part1(path) << "\n";
    std::cout << "Part 2: " << day2part2(path) << "\n";
};

int day2part1(std::filesystem::path path) {
    int double_count = 0;
    int triple_count = 0;

    std::ifstream f(path);
    std::string line;

    if (!f.is_open()) {
        std::cerr << "error opening file" << path << "\n";
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

    return double_count * triple_count;
};

std::string day2part2(std::filesystem::path path) {
    std::ifstream f(path);
    std::string line;

    if (!f.is_open()) {
        std::cerr << "error opening file" << path << "\n";
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
                return line.erase(idx, 1);
            }
        }
        prev.push_back(line);
    }

    std::cout << "not found\n";
    exit(-1);
};
