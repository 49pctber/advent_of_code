#include "solution.hpp"
#include <algorithm>
#include <fstream>
#include <iostream>
#include <set>
#include <vector>

const int target = 2020;
std::filesystem::path input_path = input_directory.append("1.txt");

void Solution::part1() {

    std::ifstream input(input_path);
    if (!input.is_open()) {
        std::cerr << "Failed to open input file" << std::endl;
    }

    std::string line;

    std::vector<int> entries;
    while (std::getline(input, line)) {
        int entry = std::stoi(line);
        entries.push_back(entry);
    }
    input.close();

    // sort to convert O(n^2) to O(n)
    std::sort(entries.begin(), entries.end());

    int low, high;
    low = 0;
    high = entries.size() - 1;

    while (low < high) {
        int sum = entries[low] + entries[high];
        if (sum > target) {
            high--;
        } else if (sum < target) {
            low++;
        } else {
            int product = entries[low] * entries[high];
            std::cout << "Part 1: " << product << std::endl;
            break;
        }
    }
}

void Solution::part2() {
    std::ifstream input(input_path);
    if (!input.is_open()) {
        std::cerr << "Failed to open input file" << std::endl;
    }

    std::string line;

    std::vector<int> entries;
    std::set<int> entry_set;
    while (std::getline(input, line)) {
        int entry = std::stoi(line);
        entries.push_back(entry);
        entry_set.insert(entry);
    }
    input.close();

    for (int low = 0; low < entries.size() - 2; low++) {
        for (int mid = low + 1; mid < entries.size() - 1; mid++) {
            int diff = target - entries[low] - entries[mid];
            if (entry_set.contains(diff)) {
                int product = entries[low] * entries[mid] * diff;
                std::cout << "Part 2: " << product << std::endl;
                return;
            }
        }
    }
}
