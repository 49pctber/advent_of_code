#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

bool search(std::vector<long int> *terms, long int target, int depth,
            long int value) {

    depth++;

    if (depth == terms->size()) {
        return value == target;
    } else if (value > target) {
        return false;
    }

    return search(terms, target, depth, value + (*terms)[depth]) ||
           search(terms, target, depth, value * (*terms)[depth]);
}

bool search2(std::vector<long int> *terms, long int target, int depth,
             long int value) {

    depth++;

    if (depth == terms->size()) {
        return value == target;
    } else if (value > target) {
        return false;
    }
    long int concat =
        std::stol(std::to_string(value) + std::to_string((*terms)[depth]));
    return search2(terms, target, depth, value + (*terms)[depth]) ||
           search2(terms, target, depth, value * (*terms)[depth]) ||
           search2(terms, target, depth, concat);
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    long int sum = 0;
    std::string line;
    while (std::getline(file, line)) {
        int x = line.find(": ");
        long int target = std::stol(line.substr(0, x));
        std::string s = line.substr(x + 2, line.size() - 2 - x);
        std::vector<long int> terms;
        std::stringstream ss(s);
        while (ss >> x) {
            terms.push_back(x);
        }

        // check vector
        if (search(&terms, target, 0, terms[0])) {
            sum += target;
        }
    }
    std::cout << "Part 1: " << sum << std::endl;
}

void Solution::part2() {

    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    long int sum = 0;
    std::string line;
    while (std::getline(file, line)) {
        int x = line.find(": ");
        long int target = std::stol(line.substr(0, x));
        std::string s = line.substr(x + 2, line.size() - 2 - x);
        std::vector<long int> terms;
        std::stringstream ss(s);
        while (ss >> x) {
            terms.push_back(x);
        }

        // check vector
        if (search2(&terms, target, 0, terms[0])) {
            sum += target;
        }
    }
    std::cout << "Part 2: " << sum << std::endl;
}