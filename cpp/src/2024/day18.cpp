#include "grid.hpp"
#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <queue>
#include <string>
#include <vector>

constexpr int n_rows = 71;
constexpr int n_cols = n_rows;

using position_t = aoc::position_t;

using tile_t = struct {
    position_t position;
    char type;
    int state;
    int distance;
};

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::vector<position_t> positions;
    while (std::getline(file, line)) {
        int i = line.find(',');
        int col = std::stoi(line.substr(0, i));
        int row = std::stoi(line.substr(i + 1, line.size()));
        position_t position{row : row, col : col};
        positions.push_back(position);
    }

    std::vector<std::vector<tile_t>> grid;
    grid.reserve(n_rows);
    for (int i = 0; i < n_rows; i++) {
        std::vector<tile_t> row(n_cols);
        for (int j = 0; j < n_cols; j++) {
            row[j].type = '.';
            row[j].state = 0;
            row[j].distance = 0;
            row[j].position.row = grid.size();
            row[j].position.col = j;
        }
        grid.push_back(row);
    }

    constexpr int n_walls = 1024;
    for (int i = 0; i < n_walls; i++) {
        position_t position = positions[i];
        int row = position.row;
        int col = position.col;
        grid[row][col].type = '#';
    }

    // BFS
    position_t start{row : 0, col : 0};
    grid[0][0].state |= 0b1; // discovered
    std::queue<tile_t *> discovered;
    discovered.push(&(grid[0][0]));

    while (!discovered.empty()) {
        tile_t *current = discovered.front();
        discovered.pop();
        current->state |= 0b10; // mark visited

        if (current->position.row > 0) {
            position_t np(current->position);
            aoc::move_up(np);
            tile_t *nt = &(grid[np.row][np.col]);
            if (nt->type != '#' && nt->state == 0) {
                nt->state |= 0b1;
                nt->distance = current->distance + 1;
                discovered.push(nt);
            }
        }

        if (current->position.col > 0) {
            position_t np(current->position);
            aoc::move_left(np);
            tile_t *nt = &(grid[np.row][np.col]);
            if (nt->type != '#' && nt->state == 0) {
                nt->state |= 0b1;
                nt->distance = current->distance + 1;
                discovered.push(nt);
            }
        }

        if (current->position.row < n_rows - 1) {
            position_t np(current->position);
            aoc::move_down(np);
            tile_t *nt = &(grid[np.row][np.col]);
            if (nt->type != '#' && nt->state == 0) {
                nt->state |= 0b1;
                nt->distance = current->distance + 1;
                discovered.push(nt);
            }
        }

        if (current->position.col < n_cols - 1) {
            position_t np(current->position);
            aoc::move_right(np);
            tile_t *nt = &(grid[np.row][np.col]);
            if (nt->type != '#' && nt->state == 0) {
                nt->state |= 0b1;
                nt->distance = current->distance + 1;
                discovered.push(nt);
            }
        }

        if (current->position.row == n_rows - 1 &&
            current->position.col == n_cols - 1) {
            break;
        }
    }

    std::cout << "Part 1: " << grid[n_rows - 1][n_cols - 1].distance
              << std::endl;
}

bool exit_possible(std::vector<position_t> &positions, int n_walls) {

    std::vector<std::vector<tile_t>> grid;
    grid.reserve(n_rows);
    for (int i = 0; i < n_rows; i++) {
        std::vector<tile_t> row(n_cols);
        for (int j = 0; j < n_cols; j++) {
            row[j].type = '.';
            row[j].state = 0;
            row[j].distance = -1;
            row[j].position.row = grid.size();
            row[j].position.col = j;
        }
        grid.push_back(row);
    }

    for (int i = 0; i < n_walls; i++) {
        position_t position = positions[i];
        int row = position.row;
        int col = position.col;
        grid[row][col].type = '#';
    }

    // BFS
    position_t start{row : 0, col : 0};
    grid[0][0].state |= 0b1; // discovered
    grid[0][0].distance = 0;
    std::queue<tile_t *> discovered;
    discovered.push(&(grid[0][0]));

    while (!discovered.empty()) {
        tile_t *current = discovered.front();
        discovered.pop();
        current->state |= 0b10; // mark visited

        if (current->position.row > 0) {
            position_t np(current->position);
            aoc::move_up(np);
            tile_t *nt = &(grid[np.row][np.col]);
            if (nt->type != '#' && nt->state == 0) {
                nt->state |= 0b1;
                nt->distance = current->distance + 1;
                discovered.push(nt);
            }
        }

        if (current->position.col > 0) {
            position_t np(current->position);
            aoc::move_left(np);
            tile_t *nt = &(grid[np.row][np.col]);
            if (nt->type != '#' && nt->state == 0) {
                nt->state |= 0b1;
                nt->distance = current->distance + 1;
                discovered.push(nt);
            }
        }

        if (current->position.row < n_rows - 1) {
            position_t np(current->position);
            aoc::move_down(np);
            tile_t *nt = &(grid[np.row][np.col]);
            if (nt->type != '#' && nt->state == 0) {
                nt->state |= 0b1;
                nt->distance = current->distance + 1;
                discovered.push(nt);
            }
        }

        if (current->position.col < n_cols - 1) {
            position_t np(current->position);
            aoc::move_right(np);
            tile_t *nt = &(grid[np.row][np.col]);
            if (nt->type != '#' && nt->state == 0) {
                nt->state |= 0b1;
                nt->distance = current->distance + 1;
                discovered.push(nt);
            }
        }

        if (current->position.row == n_rows - 1 &&
            current->position.col == n_cols - 1) {
            return true;
        }
    }

    return false;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::vector<position_t> positions;
    while (std::getline(file, line)) {
        int i = line.find(',');
        int col = std::stoi(line.substr(0, i));
        int row = std::stoi(line.substr(i + 1, line.size()));
        position_t position{row : row, col : col};
        positions.push_back(position);
    }

    int n_walls_can = 1024;
    int n_walls_cant = 71 * 71;

    while (n_walls_can + 1 < n_walls_cant) {
        int i = (n_walls_can + n_walls_cant) / 2;
        bool can = exit_possible(positions, i);

        if (can) {
            n_walls_can = i;
        } else {
            n_walls_cant = i;
        }
    }

    position_t blocker = positions[n_walls_cant - 1];

    std::cout << "Part 2: " << blocker.col << ',' << blocker.row << std::endl;
}