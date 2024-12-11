#include "solution.hpp"
#include <cmath>
#include <fstream>
#include <iostream>
#include <list>
#include <map>
#include <sstream>
#include <string>

typedef long int stone_t;
typedef int blink_t;
typedef long int count_t;
typedef std::map<stone_t, std::map<blink_t, count_t>> memo_t;

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::getline(file, line);
    std::stringstream ss(line);

    std::list<stone_t> stones;
    stone_t new_stone;
    while (ss >> new_stone) {
        stones.push_back(new_stone);
    }

    for (int i = 0; i < 25; i++) {
        for (auto it = stones.begin(); it != stones.end(); ++it) {
            if (*it == 0) {
                *it = 1;
                continue;
            }

            int n_digits = std::ceil(std::log10(*it + 1));
            if (n_digits % 2 == 0) {
                int divisor = std::pow(10, n_digits / 2);
                stone_t new_value = *it % divisor;
                *it /= divisor;
                ++it;
                stones.insert(it, new_value);
                --it;
                continue;
            }

            *it *= 2024;
        }
    }

    std::cout << "Part 1: " << stones.size() << std::endl;
}

long int partial_search(memo_t &memo, stone_t stone, blink_t n_blinks) {

    if (n_blinks == 0) {
        return 1;
    }

    // check if memo'd
    if (memo.find(stone) != memo.end() &&
        memo[stone].find(n_blinks) != memo[stone].end()) {
        return memo[stone][n_blinks];
    }

    // compute
    count_t n_stones = 0;
    int n_digits = std::ceil(std::log10(stone + 1));

    if (stone == 0) {
        n_stones += partial_search(memo, 1, n_blinks - 1);
    } else if (n_digits % 2 == 0) {
        int divisor = std::pow(10, n_digits / 2);
        n_stones += partial_search(memo, stone % divisor, n_blinks - 1);
        n_stones += partial_search(memo, stone / divisor, n_blinks - 1);
    } else {
        n_stones += partial_search(memo, stone * 2024, n_blinks - 1);
    }

    // add memo and return
    memo[stone][n_blinks] = n_stones;
    return n_stones;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::getline(file, line);
    std::stringstream ss(line);

    std::list<stone_t> stones;
    stone_t new_stone;
    while (ss >> new_stone) {
        stones.push_back(new_stone);
    }

    memo_t memo;

    count_t count = 0;
    const blink_t n_blinks = 75;
    for (auto it = stones.begin(); it != stones.end(); ++it) {
        count += partial_search(memo, *it, n_blinks);
    }

    std::cout << "Part 2: " << count << std::endl;
}