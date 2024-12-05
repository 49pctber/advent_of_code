#include "solution.hpp"
#include <array>
#include <fstream>
#include <iostream>
#include <regex>
#include <set>
#include <sstream>
#include <string>
#include <vector>

void print_vector(std::vector<int> v) {
    for (auto n : v) {
        std::cout << n << ' ';
    }
    std::cout << '\n';
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cout << "Error opening file" << std::endl;
        exit(EXIT_FAILURE);
    }

    int sum = 0;

    std::array<std::vector<int>, 100> rules;
    std::string line;

    while (std::getline(file, line)) {
        if (line.size() == 5) {
            int a = std::stoi(line.substr(0, 2));
            int b = std::stoi(line.substr(3, 2));
            rules[a].push_back(b);
            continue;
        }

        if (line.size() > 1) {
            std::stringstream ss(line);
            std::vector<int> v;
            std::string token;
            bool valid = true;

            while (std::getline(ss, token, ',')) {
                v.push_back(std::stoi(token));
            }

            // There is a weird assumption in here. Beware.
            // This assumes that all the numbers are include in some rule.
            for (std::size_t i = 0; i < v.size(); i++) {
                for (std::size_t j = i + 1; j < v.size(); j++) {
                    if (std::find(rules[v[i]].begin(), rules[v[i]].end(),
                                  v[j]) == rules[v[i]].end()) {
                        valid = false;
                        goto exit_checking;
                    }
                }
            }

        exit_checking:
            if (valid) {
                // find middle element and add to sum
                sum += v[v.size() / 2];
            }
        }
    }

    // for (int i = 0; i < 100; i++) {
    //     std::cout << i << ": ";
    //     for (int n : rules[i]) {
    //         std::cout << n << ' ';
    //     }
    //     std::cout << '\n';
    // }

    std::cout << "Part 1: " << sum << std::endl;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cout << "Error opening file" << std::endl;
        exit(EXIT_FAILURE);
    }

    int sum = 0;

    std::array<std::vector<int>, 100> rules;
    std::string line;

    while (std::getline(file, line)) {
        if (line.size() == 5) {
            int a = std::stoi(line.substr(0, 2));
            int b = std::stoi(line.substr(3, 2));
            rules[a].push_back(b);
            continue;
        }

        if (line.size() > 1) {
            std::stringstream ss(line);
            std::vector<int> v;
            std::string token;
            bool valid = true;

            while (std::getline(ss, token, ',')) {
                v.push_back(std::stoi(token));
            }

            // There is a weird assumption in here. Beware.
            // This assumes that all the numbers are include in some rule.
            for (std::size_t i = 0; i < v.size(); i++) {
                for (std::size_t j = i + 1; j < v.size(); j++) {
                    if (std::find(rules[v[i]].begin(), rules[v[i]].end(),
                                  v[j]) == rules[v[i]].end()) {
                        valid = false;
                        goto exit_checking;
                    }
                }
            }

        exit_checking:
            if (!valid) {
                // sort vector
                std::vector<int> w;

                while (v.size() > 0) {
                    // find element in v that doesn't need to be preceeded
                    std::set<int> e;
                    for (int n : v) {
                        for (int i : rules[n]) {
                            e.insert(i);
                        }
                    }
                    for (int n : v) {
                        if (!e.contains(n)) {
                            // add that element to w
                            w.push_back(n);

                            // remove that element from v
                            auto it = std::find(v.begin(), v.end(), n);
                            if (it != v.end()) {
                                v.erase(it);
                            }
                        }
                    }
                }

                // find middle element and add to sum
                sum += w[w.size() / 2];
            }
        }
    }

    std::cout << "Part 1: " << sum << std::endl;
}