#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <list>
#include <string>
#include <vector>

typedef std::vector<std::vector<char>> floor_t;
typedef char move_t;
// typedef std::list<move_t> moves_t;
typedef struct {
    int x;
    int y;
} position_t;

bool attempt_move(floor_t &floor, position_t pos, move_t move) {
    switch (move) {
    case '<':
        if (floor[pos.y][pos.x - 1] == '#') {
            return false;
        } else if (floor[pos.y][pos.x - 1] == '.') {
            floor[pos.y][pos.x - 1] = floor[pos.y][pos.x];
            floor[pos.y][pos.x] = '.';
            return true;
        } else if (attempt_move(floor, position_t{x : pos.x - 1, y : pos.y},
                                move)) {
            floor[pos.y][pos.x - 1] = floor[pos.y][pos.x];
            floor[pos.y][pos.x] = '.';
            return true;
        } else {
            return false;
        }
        break;

    case '>':
        if (floor[pos.y][pos.x + 1] == '#') {
            return false;
        } else if (floor[pos.y][pos.x + 1] == '.') {
            floor[pos.y][pos.x + 1] = floor[pos.y][pos.x];
            floor[pos.y][pos.x] = '.';
            return true;
        } else if (attempt_move(floor, position_t{x : pos.x + 1, y : pos.y},
                                move)) {
            floor[pos.y][pos.x + 1] = floor[pos.y][pos.x];
            floor[pos.y][pos.x] = '.';
            return true;
        } else {
            return false;
        }
        break;

    case '^':
        if (floor[pos.y - 1][pos.x] == '#') {
            return false;
        } else if (floor[pos.y - 1][pos.x] == '.') {
            floor[pos.y - 1][pos.x] = floor[pos.y][pos.x];
            floor[pos.y][pos.x] = '.';
            return true;
        } else if (attempt_move(floor, position_t{x : pos.x, y : pos.y - 1},
                                move)) {
            floor[pos.y - 1][pos.x] = floor[pos.y][pos.x];
            floor[pos.y][pos.x] = '.';
            return true;
        } else {
            return false;
        }
        break;

    case 'v':
        if (floor[pos.y + 1][pos.x] == '#') {
            return false;
        } else if (floor[pos.y + 1][pos.x] == '.') {
            floor[pos.y + 1][pos.x] = floor[pos.y][pos.x];
            floor[pos.y][pos.x] = '.';
            return true;
        } else if (attempt_move(floor, position_t{x : pos.x, y : pos.y + 1},
                                move)) {
            floor[pos.y + 1][pos.x] = floor[pos.y][pos.x];
            floor[pos.y][pos.x] = '.';
            return true;
        } else {
            return false;
        }
        break;

    default:
        std::cout << "This should never happen." << std::endl;
        return false;
        break;
    }
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    floor_t floor;
    std::list<move_t> moves;
    position_t robot_position;

    while (std::getline(file, line)) {
        if (line.size() > 1) {
            int x = line.find('@');
            if (x != std::string::npos) {
                robot_position.x = x;
                robot_position.y = floor.size();
            }

            std::vector<char> row(line.begin(), line.end());
            floor.push_back(row);

        } else {
            break;
        }
    }

    // std::cout << robot_position.x << ' ' << robot_position.y << '\n';

    while (std::getline(file, line)) {
        std::vector<move_t> row(line.begin(), line.end());
        for (move_t move : row) {
            moves.push_back(move);
        }
    }

    // for (move_t move : moves) {
    //     std::cout << move << ' ';
    // }
    // std::cout << std::endl;

    for (move_t move : moves) {
        if (attempt_move(floor, robot_position, move)) {
            switch (move) {
            case '<':
                robot_position.x--;
                break;

            case '^':
                robot_position.y--;
                break;

            case '>':
                robot_position.x++;
                break;
            case 'v':
                robot_position.y++;
                break;
            default:
                break;
            }
        }
    }

    long int sum = 0;
    for (int row = 0; row < floor.size(); row++) {
        for (int col = 0; col < floor[row].size(); col++) {
            std::cout << floor[row][col];
            if (floor[row][col] == 'O') {
                sum += row * 100 + col;
            }
        }
        std::cout << '\n';
    }
    std::cout << std::endl;

    std::cout << "Part 1: " << sum << std::endl;
}

void Solution::part2() {}