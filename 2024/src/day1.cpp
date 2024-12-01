#include "solution.hpp"
#include <algorithm>
#include <fstream>
#include <iostream>
#include <map>
#include <set>
#include <stdio.h>
#include <string>
#include <vector>

std::filesystem::path input_path = input_directory.append("1.txt");

void Solution::part1() {
    std::ifstream input(input_path);
    if (!input.is_open()) {
        std::cerr << "Failed to open input file" << std::endl;
    }

    std::string line;

    std::vector<int> list1;
    std::vector<int> list2;

    std::vector<int> entries;
    while (std::getline(input, line)) {
        int a, b;
        sscanf(line.c_str(), "%d %d", &a, &b);
        list1.push_back(a);
        list2.push_back(b);
    }
    input.close();

    std::sort(list1.begin(), list1.end());
    std::sort(list2.begin(), list2.end());

    int sum = 0;
    for (int i = 0; i < list1.size(); i++) {
        sum += std::abs(list1[i] - list2[i]);
    }
    std::cout << "Part 1: " << sum << std::endl;
}

void Solution::part2() {
    std::ifstream input(input_path);
    if (!input.is_open()) {
        std::cerr << "Failed to open input file" << std::endl;
    }

    std::string line;

    std::vector<int> list_left;
    std::vector<int> list_right;

    std::vector<int> entries;
    while (std::getline(input, line)) {
        int a, b;
        sscanf(line.c_str(), "%d %d", &a, &b);
        list_left.push_back(a);
        list_right.push_back(b);
    }
    input.close();

    std::sort(list_left.begin(), list_left.end());
    std::sort(list_right.begin(), list_right.end());

    std::map<int, int> counts;
    for (int x : list_right) {
        counts[x] += 1;
    }

    int sum = 0;
    for (int x : list_left) {
        if (counts.find(x) != counts.end()) {
            sum += x * counts[x];
        }
    }

    std::cout << "Part 2: " << sum << std::endl;
}
