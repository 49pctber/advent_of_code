/*
This code is pretty ugly.
And there's a lot of it.
I would be interested in techniques for handing this sort of movement problem
where both position and direction are needed.
*/

#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

typedef std::vector<std::string> grid_t;

typedef enum {
    DIR_UP = 0,
    DIR_RIGHT = 1,
    DIR_DOWN = 2,
    DIR_LEFT = 3
} direction_t;

typedef struct {
    int row;
    int col;
} position_t;

void step(position_t *p, direction_t d) {
    switch (d) {
    case DIR_UP:
        p->row--;
        break;

    case DIR_RIGHT:
        p->col++;
        break;

    case DIR_DOWN:
        p->row++;
        break;

    case DIR_LEFT:
        p->col--;
        break;

    default:
        break;
    }
}

void step_backwards(position_t *p, direction_t d) {
    switch (d) {
    case DIR_UP:
        p->row++;
        break;

    case DIR_RIGHT:
        p->col--;
        break;

    case DIR_DOWN:
        p->row--;
        break;

    case DIR_LEFT:
        p->col++;
        break;

    default:
        break;
    }
}

void turn_right(direction_t *d) {
    switch (*d) {
    case DIR_UP:
        *d = DIR_RIGHT;
        break;

    case DIR_RIGHT:
        *d = DIR_DOWN;
        break;

    case DIR_DOWN:
        *d = DIR_LEFT;
        break;

    case DIR_LEFT:
        *d = DIR_UP;
        break;

    default:
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
    grid_t grid;
    position_t pos;
    direction_t dir = DIR_UP;

    while (std::getline(file, line)) {
        grid.push_back(line);
        int col = line.find('^');
        if (col != std::string::npos) {
            pos.row = grid.size() - 1;
            pos.col = col;
        }
    }

    std::vector<std::vector<int>> counts;
    for (int row = 0; row < grid.size(); row++) {
        counts.push_back(std::vector<int>(grid[row].size(), 0));
    }

    while (true) {
        while (true) {

            // take a step
            step(&pos, dir);
            if (pos.row < 0 || pos.row >= grid.size() || pos.col < 0 ||
                pos.col >= grid[pos.row].size()) {
                goto out_of_bounds;
            }

            // backtrack if there is an obstacle, turn instead and step
            if (grid[pos.row][pos.col] == '#') {
                step_backwards(&pos, dir);
                turn_right(&dir);
                continue;
            }

            break;
        }

        // mark space sa having been visited
        counts[pos.row][pos.col]++;
    }

out_of_bounds:
    int n_pos = 0;
    for (auto row : counts) {
        for (auto col : row) {
            if (col > 0) {
                n_pos++;
            }
        }
    }

    std::cout << "Part 1: " << n_pos << std::endl;
}

bool has_cycle(grid_t *grid, position_t pos, position_t obstruction,
               direction_t dir) {
    std::vector<std::vector<int>> pos_state;
    for (int row = 0; row < grid->size(); row++) {
        pos_state.push_back(std::vector<int>((*grid)[row].size(), 0));
    }

    while (true) {
        while (true) {

            // take a step
            step(&pos, dir);

            // check if step would take us out of bounds
            if (pos.row < 0 || pos.row >= grid->size() || pos.col < 0 ||
                pos.col >= (*grid)[pos.row].size()) {
                return false;
            }

            // backtrack if there is an obstacle, turn instead and step again
            if ((*grid)[pos.row][pos.col] == '#') {
                step_backwards(&pos, dir);
                turn_right(&dir);
                continue;
            }

            if (pos.row == obstruction.row && pos.col == obstruction.col) {
                step_backwards(&pos, dir);
                turn_right(&dir);
                continue;
            }

            // we stepped forward successfully
            break;
        }

        // check if we have been in this position and orientation before
        if (pos_state[pos.row][pos.col] & 0b1 << int(dir)) {
            // we're in a cycle!
            return true;
        } else {
            // mark space as having been visited in a given orientation
            pos_state[pos.row][pos.col] |= 0b1 << int(dir);
        }
    }
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::vector<std::string> grid;
    position_t pos, start_pos;
    direction_t dir = DIR_UP;

    while (std::getline(file, line)) {
        grid.push_back(line);
        int col = line.find('^');
        if (col != std::string::npos) {
            start_pos.row = grid.size() - 1;
            start_pos.col = col;
        }
    }

    pos.row = start_pos.row;
    pos.col = start_pos.col;

    std::vector<std::vector<int>> pos_state;
    for (int row = 0; row < grid.size(); row++) {
        pos_state.push_back(std::vector<int>(grid[row].size(), 0));
    }

    while (true) {
        while (true) {

            // take a step
            step(&pos, dir);
            if (pos.row < 0 || pos.row >= grid.size() || pos.col < 0 ||
                pos.col >= grid[pos.row].size()) {
                goto out_of_bounds;
            }

            // backtrack if there is an obstacle, turn instead and step
            if (grid[pos.row][pos.col] == '#') {
                step_backwards(&pos, dir);
                turn_right(&dir);
                continue;
            }

            break;
        }

        // mark space sa having been visited
        pos_state[pos.row][pos.col] |= 0b1 < 4;
    }

out_of_bounds:
    int n_pos = 0;
    for (auto row : pos_state) {
        for (auto col : row) {
            if (col > 0) {
                n_pos++;
            }
        }
    }

    int count = 0;
    for (int row = 0; row < grid.size(); row++) {
        for (int col = 0; col < grid[row].size(); col++) {
            // only possible locations for obstructions are where guard would
            // visit during unobstructed tour
            if (pos_state[row][col] != 0) {
                bool cycle = has_cycle(
                    &grid, start_pos, position_t{row : row, col : col}, DIR_UP);
                if (cycle) {
                    count++;
                }
            }
        }
    }

    // cycle will repeat when guard is in same position and orientation as
    // any previous position/orientation
    std::cout << "Part 2: " << count << std::endl;
}