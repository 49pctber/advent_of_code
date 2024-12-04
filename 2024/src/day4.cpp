#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

bool check(std::vector<std::string> *grid, int row, int col, int dx, int dy) {
    if (row + 3 * dy < 0 || row + 3 * dy >= grid->size()) {
        return false;
    }

    if (col + 3 * dx < 0 || col + 3 * dx >= grid->size()) {
        return false;
    }

    return (*grid)[row][col] == 'X' && (*grid)[row + dy][col + dx] == 'M' &&
           (*grid)[row + 2 * dy][col + 2 * dx] == 'A' &&
           (*grid)[row + 3 * dy][col + 3 * dx] == 'S';
}

bool check2(std::vector<std::string> *grid, int row, int col) {
    if (row <= 0 || row >= grid->size() - 1) {
        return false;
    }

    if (col <= 0 || col >= grid->size() - 1) {
        return false;
    }

    if ((*grid)[row][col] != 'A') {
        return false;
    }

    if ((*grid)[row + 1][col + 1] == 'M' && (*grid)[row + 1][col - 1] == 'M' &&
        (*grid)[row - 1][col + 1] == 'S' && (*grid)[row - 1][col - 1] == 'S') {
        return true;
    }

    if ((*grid)[row + 1][col + 1] == 'M' && (*grid)[row + 1][col - 1] == 'S' &&
        (*grid)[row - 1][col + 1] == 'M' && (*grid)[row - 1][col - 1] == 'S') {
        return true;
    }
    if ((*grid)[row + 1][col + 1] == 'S' && (*grid)[row + 1][col - 1] == 'S' &&
        (*grid)[row - 1][col + 1] == 'M' && (*grid)[row - 1][col - 1] == 'M') {
        return true;
    }
    if ((*grid)[row + 1][col + 1] == 'S' && (*grid)[row + 1][col - 1] == 'M' &&
        (*grid)[row - 1][col + 1] == 'S' && (*grid)[row - 1][col - 1] == 'M') {
        return true;
    }
    return false;
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cout << "Error opening file" << std::endl;
        return;
    }

    std::vector<std::string> grid;
    std::string line;
    while (std::getline(file, line)) {
        grid.push_back(line);
    }

    int count = 0;
    for (int row = 0; row < grid.size(); row++) {
        for (int col = 0; col < grid[row].size(); col++) {
            if (grid[row][col] != 'X') {
                continue;
            }

            if (check(&grid, row, col, 1, 0)) {
                count++;
            }
            if (check(&grid, row, col, 1, 1)) {
                count++;
            }
            if (check(&grid, row, col, 0, 1)) {
                count++;
            }
            if (check(&grid, row, col, -1, 1)) {
                count++;
            }
            if (check(&grid, row, col, -1, 0)) {
                count++;
            }
            if (check(&grid, row, col, -1, -1)) {
                count++;
            }
            if (check(&grid, row, col, 0, -1)) {
                count++;
            }
            if (check(&grid, row, col, 1, -1)) {
                count++;
            }
        }
    }

    std::cout << "Part 1: " << count << std::endl;
    file.close();
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cout << "Error opening file" << std::endl;
        return;
    }

    std::vector<std::string> grid;
    std::string line;
    while (std::getline(file, line)) {
        grid.push_back(line);
    }

    int count = 0;
    for (int row = 0; row < grid.size(); row++) {
        for (int col = 0; col < grid[row].size(); col++) {
            if (check2(&grid, row, col)) {
                count++;
            }
        }
    }

    std::cout << "Part 2: " << count << std::endl;
    file.close();
}