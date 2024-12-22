#include "grid.hpp"
#include "solution.hpp"
#include <algorithm>
#include <fstream>
#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace aoc;

const std::vector<std::vector<char>> numeric_layout{
    {'7', '8', '9'}, {'4', '5', '6'}, {'1', '2', '3'}, {'#', '0', 'A'}};

const std::vector<std::vector<char>> dpad_layout{{'#', '^', 'A'},
                                                 {'<', 'v', '>'}};

class Keypad {
  public:
    Keypad(const std::vector<std::vector<char>> &layout,
           position_t start_position)
        : keys(layout), position(start_position) {
        for (int row = 0; row < keys.size(); row++) {
            for (int col = 0; col < keys[row].size(); col++) {
                key_lookup[keys[row][col]] = position_t{row : row, col : col};
            }
        }
    }

    bool follow_instruction(char instruction) {

        switch (instruction) {
        case '^':
            return move_north();

        case '>':
            return move_east();

        case 'v':
            return move_south();

        case '<':
            return move_west();

        case 'A':
            // return press();
            return true;

        default:
            std::cout << "SOMETHING WENT WRONG" << std::endl;
            return false;
        }
    }

    bool follow_instructions(std::string instructions) {
        for (char instruction : instructions) {
            if (!follow_instruction(instruction)) {
                return false;
            }
        }
        return true;
    }

    long unsigned int find_shortest(std::vector<std::string> desired_outputs) {

        // base case
        // the number of keypresses is equal to the desired instructions for the
        // human operator
        if (parent == nullptr) {
            return desired_outputs[0].size();
        }

        // recursion
        long unsigned int min_cost = UINT64_MAX;
        std::string best_instructions;

        const position_t initial_position(position);

        for (std::string desired_output : desired_outputs) {
            long unsigned int cost = 0;
            position = initial_position;

            for (char next_out : desired_output) {
                // determine candidate sequences of movements needed to get
                // there
                // TODO remove any invalid sequences

                auto d = offset(position, key_lookup[next_out]);

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
                    candidates.push_back(needed);
                } while (std::next_permutation(needed.begin(), needed.end()));

                for (int i = 0; i < candidates.size(); i++) {
                    candidates[i] += 'A';
                }

                // give candidates to parent to figure out most efficient route
                if (parent != nullptr) {
                    // pass up the chain
                    cost += parent->find_shortest(candidates);
                } else {
                    cost = candidates[0].size();
                }
            }

            if (cost < min_cost) {
                min_cost = cost;
                best_instructions += desired_output;
            }
        }

        follow_instructions(best_instructions);

        return min_cost;
    }

    Keypad *child;  // keypad that this keypad controls
    Keypad *parent; // keypad that controls this keypad

  private:
    std::vector<std::vector<char>> keys;
    std::map<char, position_t> key_lookup;
    position_t position;

    bool move_north() {
        if (position.row <= 0) {
            return false;
        }

        move_up(position);
        return true;
    }

    bool move_south() {
        if (position.row >= keys.size() - 1) {
            return false;
        }

        move_down(position);
        return true;
    }

    bool move_east() {
        if (position.col >= keys[0].size() - 1) {
            return false;
        }

        move_right(position);
        return true;
    }

    bool move_west() {
        if (position.col <= 0) {
            return false;
        }

        move_left(position);
        return true;
    }

    bool press() {
        char instruction = keys[position.row][position.col];

        if (child != nullptr) {
            return child->follow_instruction(instruction);
        }

        return true;
    }
};

class State {
  public:
    State()
        : numeric(numeric_layout, position_t{row : 3, col : 2}),
          dpad1(dpad_layout, position_t{row : 0, col : 2}),
          dpad2(dpad_layout, position_t{row : 0, col : 2}),
          dpad3(dpad_layout, position_t{row : 0, col : 2}) {
        numeric.child = nullptr;
        numeric.parent = &dpad1;
        dpad1.child = &numeric;
        dpad1.parent = &dpad2;
        dpad2.child = &dpad1;
        dpad2.parent = &dpad3;
        dpad3.child = &dpad2;
        dpad3.parent = nullptr;
    }

    void instruct(char instruction) { dpad2.follow_instruction(instruction); }

    long int complexity(std::string desired_output) {
        long int shortest_sequence_length =
            numeric.find_shortest(std::vector<std::string>{desired_output});

        int numeric_part = std::stoi(desired_output.substr(0, 3));

        std::cout << shortest_sequence_length << '*' << numeric_part << '\n';

        return shortest_sequence_length * numeric_part;
    }

  private:
    Keypad numeric;
    Keypad dpad1;
    Keypad dpad2;
    Keypad dpad3;
};

// void test() {
//     State s;
//     // std::string input =
//     //
//     "<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A";
//     std::string input =
//         "<v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A";
//     for (char instruction : input) {
//         s.instruct(instruction);
//     }
//     std::cout << std::endl;
// }

void test() {
    long unsigned int sum = 0;
    for (std::string input : {"029A", "980A", "179A", "456A", "379A"}) {
        State s;
        long unsigned int complexity = s.complexity(input);
        sum += complexity;
        std::cout << input << ": " << complexity << '\n';
    }

    std::cout << "Test (sum): " << sum << std::endl;
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
    test();

    long unsigned int part1 = 0;
    std::cout << "Part 1: " << part1 << std::endl;
}

void Solution::part2() {
    long int part2 = 0;
    std::cout << "Part 2: " << part2 << std::endl;
}