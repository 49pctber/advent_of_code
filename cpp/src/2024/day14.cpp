#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <list>
#include <regex>
#include <string>
#include <vector>

typedef struct {
    int x;
    int y;
    int vx;
    int vy;
} robot_t;

void print_robot(robot_t &robot) {
    std::cout << robot.x << ' ' << robot.y << ' ' << robot.vx << ' ' << robot.vy
              << '\n';
}
void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::list<robot_t> robots;
    std::regex pattern(R"(p=([0-9]+),([0-9]+) v=(\-?[0-9]+),(\-?[0-9]+))");
    std::smatch match;
    while (std::getline(file, line)) {
        if (std::regex_search(line, match, pattern)) {
            robot_t robot{
                x : std::stoi(match[1]),
                y : std::stoi(match[2]),
                vx : std::stoi(match[3]),
                vy : std::stoi(match[4])
            };
            robots.push_back(robot);
        }
    }

    const long int duration = 100;
    const int width = 101;
    const int height = 103;
    std::vector<int> counts(4, 0);

    for (auto robot : robots) {

        int quadrant = 0;

        int x = (robot.x + duration * robot.vx) % width;
        if (x < 0) {
            x = width + x;
        }

        if (x == width / 2) {
            continue;
        } else if (x > width / 2) {
            quadrant ^= 0b1;
        }

        int y = (robot.y + duration * robot.vy) % height;
        if (y < 0) {
            y = height + y;
        }

        if (y == height / 2) {
            continue;
        } else if (y > height / 2) {
            quadrant ^= 0b10;
        }

        counts[quadrant]++;
    }

    int safety_factor = 1;
    for (auto count : counts) {
        safety_factor *= count;
    }
    std::cout << "Part 1: " << safety_factor << std::endl;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::list<robot_t> robots;
    std::regex pattern(R"(p=([0-9]+),([0-9]+) v=(\-?[0-9]+),(\-?[0-9]+))");
    std::smatch match;
    while (std::getline(file, line)) {
        if (std::regex_search(line, match, pattern)) {
            robot_t robot{
                x : std::stoi(match[1]),
                y : std::stoi(match[2]),
                vx : std::stoi(match[3]),
                vy : std::stoi(match[4])
            };
            robots.push_back(robot);
        }
    }

    const int width = 101;
    const int height = 103;

    std::vector<std::vector<bool>> floor;
    floor.resize(height);
    for (int i = 0; i < floor.size(); i++) {
        floor[i].resize(width);
    }

    for (int duration = 0; duration < width * height; duration++) {

        std::vector<int> counts(4, 0);

        for (int row = 0; row < floor.size(); row++) {
            for (int col = 0; col < floor[0].size(); col++) {
                floor[row][col] = false;
            }
        }

        for (auto robot : robots) {

            int quadrant = 0;

            int x = (robot.x + duration * robot.vx) % width;
            if (x < 0) {
                x = width + x;
            }

            int y = (robot.y + duration * robot.vy) % height;
            if (y < 0) {
                y = height + y;
            }

            floor[x][y] = true;

            if (x == width / 2) {
                continue;
            } else if (x > width / 2) {
                quadrant ^= 0b1;
            }
            if (y == height / 2) {
                continue;
            } else if (y > height / 2) {
                quadrant ^= 0b10;
            }

            counts[quadrant]++;
        }

        // std::cout << duration << '\n';
        // for (auto row : floor) {
        //     for (auto col : row) {
        //         if (col) {
        //             std::cout << 'X';
        //         } else {
        //             std::cout << ' ';
        //         }
        //     }
        //     std::cout << '\n';
        // }
        // std::cout << '\n';

        for (auto row : floor) {
            int count = 0;
            for (int i = 0; i < row.size(); i++) {
                if (row[i]) {
                    count++;
                    if (count == 20) {
                        std::cout << "Part 2: possibly " << duration << '\n';
                    }
                } else {
                    count = 0;
                }
            }
        }
    }
}