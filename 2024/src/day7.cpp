/*
I updated this file with a more optimized approach.

search() and search2() were my original solutions. Both parts take about 0.16
seconds on my computer. They are standard recursive backtracking algorithms.

new_search() and new_search2() take a different approach. Instead of starting at
zero and finding the target, they instead work backwards. They start at the
target and see if multiplication and concatenation are even possible operations.
This leads to far more efficient pruning.
*/
#include "solution.hpp"
#include <cmath>
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

bool new_search(std::vector<long int> *terms, int depth, long int value) {
    if (depth == 0) {
        return value == 0;
    }

    depth--;
    if (value % (*terms)[depth] == 0) {
        if (new_search(terms, depth, value / (*terms)[depth])) {
            return true;
        }
    }
    return new_search(terms, depth, value - (*terms)[depth]);
}

bool new_search2(std::vector<long int> *terms, int depth, long int value) {
    if (depth == 0) {
        return value == 0;
    }
    if (value < 0) {
        return false;
    }

    depth--;

    int b = std::pow(10, std::ceil(std::log10((*terms)[depth] + 1)));
    if (value % b == (*terms)[depth]) {
        if (new_search2(terms, depth, value / b)) {
            return true;
        }
    }

    if (value % (*terms)[depth] == 0) {
        if (new_search2(terms, depth, value / (*terms)[depth])) {
            return true;
        }
    }

    return new_search2(terms, depth, value - (*terms)[depth]);
}

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

    // commented code takes about 10x longer than uncommented code
    // long int concat =
    //     std::stol(std::to_string(value) + std::to_string((*terms)[depth]));
    int b = (*terms)[depth];
    long int concat = value * std::pow(10, std::ceil(std::log10(b + 1))) + b;

    return search2(terms, target, depth, value + b) ||
           search2(terms, target, depth, value * b) ||
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

        if (new_search(&terms, terms.size(), target)) {
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

        if (new_search2(&terms, terms.size(), target)) {
            sum += target;
        }
    }
    std::cout << "Part 2: " << sum << std::endl;
}