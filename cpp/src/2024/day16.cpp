#include "grid.hpp"
#include "solution.hpp"
#include <climits>
#include <fstream>
#include <iostream>
#include <set>
#include <stack>
#include <string>
#include <vector>

using namespace aoc;

typedef long int cost_t;

typedef struct {
    cost_t min_cost;
    char type;
    int visited;
} tile_t;

typedef std::vector<std::vector<tile_t>> maze_t;

cost_t search(maze_t &maze, orientation_t orientation, cost_t cost) {
    int row = orientation.position.row;
    int col = orientation.position.col;
    int dir = orientation.direction;
    tile_t *here = &(maze[row][col]);

    if (here->type == '#') {
        return INT_MAX;
    }

    // check if we have visited this tile from this direction
    if (here->visited >> ((dir ^ 0b10) & 0b1)) {
        if (cost > here->min_cost) {
            return here->min_cost;
        }
    }
    // mark visited from this direction
    here->visited |= 1 << (dir ^ 0b10);
    here->min_cost = std::min(here->min_cost, cost);

    // go to neighbors
    orientation_t forward(orientation);
    move_forward(forward);
    search(maze, forward, cost + 1);

    orientation_t right(orientation);
    turn_right(right);
    move_forward(right);
    search(maze, right, cost + 1001);

    orientation_t left(orientation);
    turn_left(left);
    move_forward(left);
    search(maze, left, cost + 1001);

    return here->min_cost;
}

bool search_2(maze_t &maze, orientation_t orientation, cost_t cost,
              cost_t target, std::set<position_t> &best_seats) {
    int row = orientation.position.row;
    int col = orientation.position.col;
    int dir = orientation.direction;
    tile_t *here = &(maze[row][col]);

    if (here->type == '#') {
        return false;
    }

    if (cost > target) {
        return false;
    }

    if (cost > here->min_cost + 1000) {
        // no possible way for this to happen on puzzle input
        // two extra turns will certainly result in a worse path
        return false;
    }

    if (here->type == 'E' && cost == target) {
        best_seats.insert(orientation.position);
        return true;
    }

    bool on_best_path = false;

    // go to neighbors
    orientation_t forward(orientation);
    move_forward(forward);
    on_best_path |= search_2(maze, forward, cost + 1, target, best_seats);

    orientation_t right(orientation);
    turn_right(right);
    move_forward(right);
    on_best_path |= search_2(maze, right, cost + 1001, target, best_seats);

    orientation_t left(orientation);
    turn_left(left);
    move_forward(left);
    on_best_path |= search_2(maze, left, cost + 1001, target, best_seats);

    if (on_best_path) {
        best_seats.insert(orientation.position);
    }

    return on_best_path;
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    maze_t maze;
    orientation_t start, end;
    while (std::getline(file, line)) {
        std::vector<tile_t> row;
        row.reserve(line.size());
        for (int i = 0; i < line.size(); i++) {
            row.push_back(tile_t{INT_MAX, line[i], 0});
            if (line[i] == 'S') {
                start.position.row = maze.size();
                start.position.col = i;
                start.direction = dir_right;
            } else if (line[i] == 'E') {
                end.position.row = maze.size();
                end.position.col = i;
                end.direction = dir_right;
            }
        }
        maze.push_back(row);
    }

    search(maze, start, 0);
    cost_t min_cost = maze[end.position.row][end.position.col].min_cost;
    std::cout << "Part 1: " << min_cost << std::endl;

    std::set<position_t> best_seats;
    for (int row = 0; row < maze.size(); row++) {
        for (int col = 0; col < maze[row].size(); col++) {
            maze[row][col].visited = 0;
        }
    }
    search_2(maze, start, 0, min_cost, best_seats);
    std::cout << "Part 2: " << best_seats.size() << std::endl;
}

void Solution::part2() {}