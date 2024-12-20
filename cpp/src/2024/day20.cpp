#include "grid.hpp"
#include "solution.hpp"
#include <algorithm>
#include <fstream>
#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace aoc;

struct cell_t {
    position_t position;
    int distance;
    std::vector<int> neighbor_distances;
    char type;
};

using grid_t = std::vector<std::vector<cell_t>>;

void fill_direction(grid_t &grid, orientation_t orientation, int distance) {
    int row = orientation.position.row;
    int col = orientation.position.col;
    direction_t dir = orientation.direction;

    if (grid[row][col].type == '#') {
        grid[row][col].neighbor_distances.push_back(distance - 1);
        return;
    }

    grid[row][col].distance = distance;

    orientation_t o1(orientation);
    move_forward(o1);
    fill_direction(grid, o1, distance + 1);

    orientation_t o2(orientation);
    turn_right(o2);
    move_forward(o2);
    fill_direction(grid, o2, distance + 1);

    orientation_t o3(orientation);
    turn_left(o3);
    move_forward(o3);
    fill_direction(grid, o3, distance + 1);
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    grid_t grid;
    orientation_t start;
    while (std::getline(file, line)) {
        std::vector<cell_t> row(line.size());
        for (int i = 0; i < line.size(); i++) {
            row[i].position.row = grid.size();
            row[i].position.col = i;
            row[i].type = line[i];
            if (line[i] == 'S') {
                start.position.row = grid.size();
                start.position.col = i;
                start.direction = dir_right;
                if (line[i - 1] == '.') {
                    start.direction = dir_left;
                }
            }
        }
        grid.push_back(row);
    }

    fill_direction(grid, start, 0);

    int count = 0;
    std::map<int, int> cheat_counts;

    for (auto row : grid) {
        for (auto cell : row) {
            if (cell.type == '#') {
                if (cell.neighbor_distances.size() >= 2) {
                    int max = *std::max_element(cell.neighbor_distances.begin(),
                                                cell.neighbor_distances.end());
                    int min = *std::min_element(cell.neighbor_distances.begin(),
                                                cell.neighbor_distances.end());
                    int distance = max - min - 2;
                    cheat_counts[distance]++;
                    if (distance >= 100) {
                        count++;
                    }
                }
            }
        }
    }

    std::cout << "Part 1: " << count << std::endl;
}

void Solution::part2() {}