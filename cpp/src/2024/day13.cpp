#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <list>
#include <regex>
#include <string>

typedef struct {
    int XA, YA, XB, YB;
    long int targetX, targetY;
} game_t;

long int get_n_tokens(game_t &game) {

    int det = game.XA * game.YB - game.XB * game.YA;

    long int A = game.YB * game.targetX - game.XB * game.targetY;
    long int B = game.XA * game.targetY - game.YA * game.targetX;

    if (A % det != 0 || B % det != 0) {
        return 0;
    }

    long int PA = A / det;
    long int PB = B / det;

    return PA * 3 + PB;
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::list<game_t> games;
    game_t game;
    std::regex pattern(
        R"(Button (A): X\+([0-9]+), Y\+([0-9]+)|Button (B): X\+([0-9]+), Y\+([0-9]+)|(Prize): X=([0-9]+), Y=([0-9]+))");
    std::smatch match;
    while (std::getline(file, line)) {
        if (std::regex_search(line, match, pattern)) {
            if (match[1] == "A") {
                game.XA = std::stoi(match[2]);
                game.YA = std::stoi(match[3]);
            } else if (match[4] == "B") {
                game.XB = std::stoi(match[5]);
                game.YB = std::stoi(match[6]);
            } else {
                game.targetX = std::stoi(match[8]);
                game.targetY = std::stoi(match[9]);
            }
        } else {
            games.push_back(game);
        }
    }
    games.push_back(game);

    int count = 0;
    for (auto game : games) {
        int n_tokens = get_n_tokens(game);
        count += n_tokens;
    }

    std::cout << "Part 1: " << count << std::endl;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::list<game_t> games;
    game_t game;
    std::regex pattern(
        R"(Button (A): X\+([0-9]+), Y\+([0-9]+)|Button (B): X\+([0-9]+), Y\+([0-9]+)|(Prize): X=([0-9]+), Y=([0-9]+))");
    std::smatch match;
    while (std::getline(file, line)) {
        if (std::regex_search(line, match, pattern)) {
            if (match[1] == "A") {
                game.XA = std::stoi(match[2]);
                game.YA = std::stoi(match[3]);
            } else if (match[4] == "B") {
                game.XB = std::stoi(match[5]);
                game.YB = std::stoi(match[6]);
            } else {
                game.targetX = std::stoi(match[8]) + 10000000000000;
                game.targetY = std::stoi(match[9]) + 10000000000000;
            }
        } else {
            games.push_back(game);
        }
    }
    games.push_back(game);

    long int count = 0;
    for (auto game : games) {
        long int n_tokens = get_n_tokens(game);
        count += n_tokens;
    }

    std::cout << "Part 2: " << count << std::endl;
}