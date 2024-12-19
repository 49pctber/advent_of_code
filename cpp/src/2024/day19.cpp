#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <map>
#include <regex>
#include <set>
#include <string>
#include <vector>

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::getline(file, line);

    std::string prefix = "^((";
    std::string postfix = "))+$";
    std::string from = ", ";
    std::string to = ")|(";

    size_t pos = 0;
    while ((pos = line.find(from, pos)) != std::string::npos) {
        line.replace(pos, from.length(), to);
        pos += to.length();
    }

    std::regex pattern(prefix + line + postfix);

    int n_possible = 0;
    std::getline(file, line); // ignore blank line
    while (std::getline(file, line)) {
        if (std::regex_match(line, pattern)) {
            n_possible++;
        }
    }
    std::cout << "Part 1: " << n_possible << std::endl;
}

long int count(std::set<std::string> &tokens, std::string s,
               std::map<std::string, long int> &memos) {

    // check if end of string
    if (s.length() == 0) {
        return 1;
    }

    // check for memo
    if (memos.find(s) != memos.end()) {
        return memos[s];
    }

    // remove any prefixes and recursively count number of possibilities
    long int c = 0;
    for (int i = 1; i <= s.size(); i++) {
        std::string prefix = s.substr(0, i);
        if (tokens.find(prefix) != tokens.end()) {
            std::string remaining = s.substr(i, s.size() - i);
            c += count(tokens, remaining, memos);
        }
    }

    // add memo
    memos[s] = c;

    return c;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::map<std::string, long int> memos;
    std::set<std::string> tokens;
    std::string line;
    std::getline(file, line);

    std::stringstream ss(line);
    std::string token;
    while (std::getline(ss, token, ',')) {
        if (token[0] == ' ') {
            token = token.substr(1, token.size());
        }
        tokens.insert(token);
    }

    long int c = 0;
    std::getline(file, line); // ignore blank line
    while (std::getline(file, line)) {
        c += count(tokens, line, memos);
    }
    std::cout << "Part 2: " << c << std::endl;
}