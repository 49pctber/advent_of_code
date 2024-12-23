#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <map>
#include <set>
#include <string>
#include <vector>

const int mask = (0b1 << 24) - 1;

int next(int secret_number) {
    secret_number = ((secret_number << 6) ^ secret_number) & mask;
    secret_number = ((secret_number >> 5) ^ secret_number) & mask;
    secret_number = ((secret_number << 11) ^ secret_number) & mask;
    return secret_number;
}

int after(int secret_number, int n_iterations) {
    for (int i = 0; i < n_iterations; i++) {
        secret_number = next(secret_number);
    }
    return secret_number;
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    long int sum = 0;
    while (std::getline(file, line)) {
        int start = std::stoi(line);
        sum += after(start, 2000);
    }

    std::cout << "Part 1: " << sum << std::endl;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::vector<int> prices(2000, 0);
    std::vector<int> price_changes(2000, 0);
    std::vector<int> updated_keys(2000, 0); // for optimization

    std::vector<int> bananas(19 * 19 * 19 * 19, 0);
    std::vector<int> new_bananas(19 * 19 * 19 * 19);

    while (std::getline(file, line)) {
        int state = std::stoi(line);
        int prev_price = state % 10;

        for (int i = 0; i < 2000; i++) {
            state = next(state);
            int curr_price = state % 10;
            prices[i] = curr_price;
            price_changes[i] = curr_price - prev_price;
            prev_price = curr_price;
        }

        for (int i = 1999; i >= 3; i--) {
            int a = price_changes[i - 3];
            int b = price_changes[i - 2];
            int c = price_changes[i - 1];
            int d = price_changes[i];

            a = a >= 0 ? a : (9 - a);
            b = b >= 0 ? b : (9 - b);
            c = c >= 0 ? c : (9 - c);
            d = d >= 0 ? d : 9 - d;

            int key = ((a * 19 + b) * 19 + c) * 19 + d;

            updated_keys[i] = key;
            new_bananas[key] = prices[i];
        }

        for (size_t i = 3; i < updated_keys.size(); i++) {
            int key = updated_keys[i];
            bananas[key] += new_bananas[key];
            new_bananas[key] = 0;
        }
    }

    // find sequence that results in best price
    long int max_bananas = INT64_MIN;
    for (auto number : bananas) {
        if (number > max_bananas) {
            max_bananas = number;
        }
    }

    std::cout << "Part 2: " << max_bananas << std::endl;
}