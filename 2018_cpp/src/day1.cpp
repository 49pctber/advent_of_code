#include "aoc.hpp"

int day1part1(std::filesystem::path path) {
    std::ifstream f(path);

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

    return current_frequency;
}

int day1part2(std::filesystem::path path) {
    std::ifstream f(path);

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
            return current_frequency;
        } else {
            i += 1;
            i %= list.size();
            seen.insert(current_frequency);
        }
    }
}

void day1() {
    std::cout << "Day 1\n";
    std::filesystem::path input_dir(INPUT_DIR);
    std::filesystem::path path = input_dir / "1.txt";
    std::cout << "Part 1: " << day1part1(path) << "\n";
    std::cout << "Part 2: " << day1part2(path) << "\n";
}