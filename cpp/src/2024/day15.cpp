#include "solution.hpp"
#include <algorithm>
#include <fstream>
#include <iostream>
#include <list>
#include <set>
#include <string>
#include <vector>

typedef std::vector<std::vector<char>> floor_t;
typedef char move_t;
// typedef std::list<move_t> moves_t;
typedef struct {
    int x;
    int y;
} position_t;

bool operator<(const position_t &a, const position_t &other) {
    if (a.y == other.y) {
        return a.x < other.x;
    }
    return a.y < other.y;
}

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
            if (floor[row][col] == 'O') {
                sum += row * 100 + col;
            }
        }
    }

    std::cout << "Part 1: " << sum << std::endl;
}

void actually_move(floor_t &floor, move_t move,
                   std::set<position_t> &affected) {
    std::vector<position_t> sorted(affected.begin(), affected.end());
    std::sort(sorted.begin(), sorted.end());

    // for (position_t pos : sorted) {
    //     std::cout << pos.x << ',' << pos.y << "; ";
    // }
    // std::cout << '\n';

    switch (move) {
    case '<':
        for (auto p : sorted) {
            char tmp = floor[p.y][p.x - 1];
            floor[p.y][p.x - 1] = floor[p.y][p.x];
            floor[p.y][p.x] = tmp;
        }
        break;

    case '>':
        std::reverse(sorted.begin(), sorted.end());
        for (auto p : sorted) {
            char tmp = floor[p.y][p.x + 1];
            floor[p.y][p.x + 1] = floor[p.y][p.x];
            floor[p.y][p.x] = tmp;
        }
        break;

    case '^':
        for (auto p : sorted) {
            char tmp = floor[p.y - 1][p.x];
            floor[p.y - 1][p.x] = floor[p.y][p.x];
            floor[p.y][p.x] = tmp;
        }
        break;

    case 'v':
        std::reverse(sorted.begin(), sorted.end());
        for (auto p : sorted) {
            char tmp = floor[p.y + 1][p.x];
            floor[p.y + 1][p.x] = floor[p.y][p.x];
            floor[p.y][p.x] = tmp;
        }
        break;

    default:
        break;
    }
}

bool attempt_move_2(floor_t &floor, position_t pos, move_t move,
                    std::set<position_t> &affected) {

    if (floor[pos.y][pos.x] == '#') {
        return false;
    }
    if (floor[pos.y][pos.x] == '.') {
        return true;
    }

    if (affected.find(pos) != affected.end()) {
        return true;
    } else {
        affected.insert(pos);
    }

    // bool able_to_move;

    switch (move) {
    case '<':
        if (attempt_move_2(floor, position_t{x : pos.x - 1, y : pos.y}, move,
                           affected)) {
            // able_to_move = true;
        } else {
            return false;
        }
        break;

    case '>':
        if (attempt_move_2(floor, position_t{x : pos.x + 1, y : pos.y}, move,
                           affected)) {
            // able_to_move = true;
        } else {
            return false;
        }
        break;

    case '^':
        if (attempt_move_2(floor, position_t{x : pos.x, y : pos.y - 1}, move,
                           affected)) {
            // able_to_move = true;
        } else {
            return false;
        }

        if (floor[pos.y][pos.x] == '[') {
            if (attempt_move_2(floor, position_t{x : pos.x + 1, y : pos.y},
                               move, affected)) {
                // able_to_move = true;
            } else {
                return false;
            }
        }

        if (floor[pos.y][pos.x] == ']') {
            if (attempt_move_2(floor, position_t{x : pos.x - 1, y : pos.y},
                               move, affected)) {
                // able_to_move = true;
            } else {
                return false;
            }
        }

        break;

    case 'v':
        if (attempt_move_2(floor, position_t{x : pos.x, y : pos.y + 1}, move,
                           affected)) {
            // able_to_move = true;
        } else {
            return false;
        }

        if (floor[pos.y][pos.x] == '[') {
            if (attempt_move_2(floor, position_t{x : pos.x + 1, y : pos.y},
                               move, affected)) {
                // able_to_move = true;
            } else {
                return false;
            }
        }

        if (floor[pos.y][pos.x] == ']') {
            if (attempt_move_2(floor, position_t{x : pos.x - 1, y : pos.y},
                               move, affected)) {
                // able_to_move = true;
            } else {
                return false;
            }
        }
        break;

    default:
        std::cout << "This should never happen." << std::endl;
        return false;
        break;
    }

    if (floor[pos.y][pos.x] == '@') {
        actually_move(floor, move, affected);
    }
    return true;
}

void Solution::part2() {
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
                robot_position.x = 2 * x;
                robot_position.y = floor.size();
            }

            std::vector<char> row;
            row.reserve(line.size() * 2);
            for (char c : line) {
                switch (c) {
                case '#':
                    row.push_back('#');
                    row.push_back('#');
                    break;

                case '.':
                    row.push_back('.');
                    row.push_back('.');
                    break;

                case 'O':
                    row.push_back('[');
                    row.push_back(']');
                    break;

                case '@':
                    row.push_back('@');
                    row.push_back('.');
                    break;

                default:
                    break;
                }
            }
            floor.push_back(row);

        } else {
            break;
        }
    }

    while (std::getline(file, line)) {
        std::vector<move_t> row(line.begin(), line.end());
        for (move_t move : row) {
            moves.push_back(move);
        }
    }

    for (move_t move : moves) {

        // for (auto row : floor) {
        //     for (auto col : row) {
        //         std::cout << col;
        //     }
        //     std::cout << '\n';
        // }
        // std::cout << move << std::endl;

        std::set<position_t> affected;
        if (attempt_move_2(floor, robot_position, move, affected)) {
            switch (move) {
            case '<':
                robot_position.x--;
                break;
            case '>':
                robot_position.x++;
                break;
            case '^':
                robot_position.y--;
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
            if (floor[row][col] == '[') {
                sum += row * 100 + col;
            }
        }
    }

    std::cout << "Part 2: " << sum << std::endl; // 1480222 < ans
}