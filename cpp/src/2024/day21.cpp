#include "grid.hpp"
#include "solution.hpp"
#include <algorithm>
#include <fstream>
#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace aoc;

using layout_t = std::vector<std::vector<char>>;
using lookup_t = std::map<char, position_t>;
using cost_t = long int;

layout_t numeric_layout{
    {'7', '8', '9'}, {'4', '5', '6'}, {'1', '2', '3'}, {'#', '0', 'A'}};

layout_t dpad_layout{{'#', '^', 'A'}, {'<', 'v', '>'}};

lookup_t numeric_lookup;
lookup_t dpad_lookup;
int stop_condition = 3;

std::map<std::string, cost_t> cache;

void initialize_lookups() {
    for (int row = 0; row < numeric_layout.size(); row++) {
        for (int col = 0; col < numeric_layout[0].size(); col++) {
            numeric_lookup[numeric_layout[row][col]] =
            position_t{row : row, col : col};
        }
    }

    for (int row = 0; row < dpad_layout.size(); row++) {
        for (int col = 0; col < dpad_layout[0].size(); col++) {
            dpad_lookup[dpad_layout[row][col]] =
            position_t{row : row, col : col};
        }
    }
}

cost_t compute_cost(std::string request, int depth) {

    if (depth == stop_condition) {
        return request.size();
    }

    std::string key = request + '/' + std::to_string(depth);
    if (cache.find(key) != cache.end()) {
        return cache[key];
    }

    layout_t *layout;
    lookup_t *lookup;
    position_t curr_position;

    if (depth == 0) {
        layout = &numeric_layout;
        lookup = &numeric_lookup;
    } else {
        layout = &dpad_layout;
        lookup = &dpad_lookup;
    }
    curr_position = (*lookup)['A'];
    position_t next_position;

    cost_t cost = 0;
    // for each step needed, find minimum cost route to get there
    for (char next : request) {
        // enumerate possible routes
        // ignore any that are oob
        next_position = (*lookup)[next];
        auto d = offset(curr_position, next_position);

        std::string needed;
        if (d.row > 0) {
            needed += std::string(d.row, 'v');
        } else if (d.row < 0) {
            needed += std::string(-d.row, '^');
        }

        if (d.col > 0) {
            needed += std::string(d.col, '>');
        } else if (d.col < 0) {
            needed += std::string(-d.col, '<');
        }

        std::sort(needed.begin(), needed.end());

        std::vector<std::string> candidates;
        do {
            position_t test_position(curr_position);
            bool valid = true;

            for (char instruction : needed) {
                switch (instruction) {
                case '^':
                    move_up(test_position);
                    break;

                case '>':
                    move_right(test_position);
                    break;
                case 'v':
                    move_down(test_position);
                    break;
                case '<':
                    move_left(test_position);
                    break;

                default:
                    break;
                }

                if ((*layout)[test_position.row][test_position.col] == '#') {
                    valid = false;
                    break;
                }
            }

            if (valid) {
                candidates.push_back(needed);
            }
        } while (std::next_permutation(needed.begin(), needed.end()));

        for (int i = 0; i < candidates.size(); i++) {
            candidates[i] += 'A';
        }

        // query how much they will cost
        // choose the lowest
        cost_t min_cost = INT64_MAX;
        for (auto candidate : candidates) {
            cost_t c = compute_cost(candidate, depth + 1);

            if (c < min_cost) {
                min_cost = c;
            }
        }

        // add to total cost
        cost += min_cost;

        // move position for next loop
        curr_position = next_position;
    }

    cache[key] = cost;

    return cost;
}

cost_t compute_complexity(std::string input) {
    cache.clear();
    cost_t shortest_sequence_length = compute_cost(input, 0);
    int numeric_part = std::stoi(input.substr(0, input.size() - 1));
    // std::cout << shortest_sequence_length << '*' << numeric_part <<
    // std::endl;
    return shortest_sequence_length * numeric_part;
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::vector<std::string> inputs;
    while (std::getline(file, line)) {
        inputs.push_back(line);
    }

    initialize_lookups();
    stop_condition = 3;

    cost_t sum = 0;
    for (auto input : inputs) {
        sum += compute_complexity(input);
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
    std::vector<std::string> inputs;
    while (std::getline(file, line)) {
        inputs.push_back(line);
    }

    // initialize_lookups();
    stop_condition = 26;

    cost_t sum = 0;
    for (auto input : inputs) {
        sum += compute_complexity(input);
    }
    std::cout << "Part 2: " << sum << std::endl;
}