#include "solution.hpp"
#include <climits>
#include <fstream>
#include <iostream>
#include <set>
#include <stack>
#include <string>
#include <vector>

typedef long int cost_t;

typedef struct {
    cost_t min_cost;
    char type;
    int visited;
} tile_t;

typedef struct {
    int row;
    int col;
} position_t;

bool operator<(const position_t &lhs, const position_t &rhs) {
    if (lhs.row == rhs.row) {
        return lhs.col < rhs.col;
    }
    return lhs.row < rhs.row;
}

typedef int direction_t;
const int dir_right = 0;
const int dir_down = 1;
const int dir_left = 2;
const int dir_up = 3;

typedef struct {
    position_t position;
    direction_t direction;
} orientation_t;

typedef std::vector<std::vector<tile_t>> maze_t;

void move_forward(orientation_t &o) {
    switch (o.direction) {
    case dir_right:
        o.position.col++;
        break;

    case dir_down:
        o.position.row++;
        break;

    case dir_left:
        o.position.col--;
        break;

    case dir_up:
        o.position.row--;
        break;

    default:
        break;
    }
}

void turn_right(orientation_t &o) {
    o.direction++;
    o.direction &= 0b11;
}

void turn_left(orientation_t &o) {
    o.direction--;
    o.direction &= 0b11;
}

void turn_around(orientation_t &o) {
    // o.direction += 2;
    // o.direction &= 0b11;
    o.direction ^= 0b10;
}

bool search(maze_t &maze, orientation_t orientation, cost_t cost, cost_t target,
            std::set<position_t> &best_seats) {
    int row = orientation.position.row;
    int col = orientation.position.col;
    int dir = orientation.direction;
    tile_t *here = &(maze[row][col]);

    if (here->type == '#') {
        return false;
    }
    if (here->type == 'E') {
        best_seats.insert(orientation.position);
        return cost == target;
    }

    // check if we have visited this tile from this direction
    if (here->visited >> ((dir ^ 0b10) & 0b1)) {
        if (cost > here->min_cost) {
            return false;
        }
    }
    // mark visited from this direction
    here->visited |= 1 << (dir ^ 0b10);
    here->min_cost = std::min(here->min_cost, cost);

    bool on_min_path = false;

    // go to neighbors
    orientation_t forward(orientation);
    move_forward(forward);
    on_min_path |= search(maze, forward, cost + 1, target, best_seats);

    orientation_t right(orientation);
    turn_right(right);
    move_forward(right);
    on_min_path |= search(maze, right, cost + 1001, target, best_seats);

    orientation_t left(orientation);
    turn_left(left);
    move_forward(left);
    on_min_path |= search(maze, left, cost + 1001, target, best_seats);

    orientation_t backward(orientation);
    turn_around(backward);
    move_forward(backward);
    on_min_path |= search(maze, backward, cost + 2001, target, best_seats);

    if (on_min_path) {
        best_seats.insert(orientation.position);
    }

    return on_min_path;
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::vector<std::vector<tile_t>> maze;
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

    std::set<position_t> best_seats;
    search(maze, start, 0, INT_MAX, best_seats);
    cost_t min_cost = maze[end.position.row][end.position.col].min_cost;
    std::cout << "Part 1: " << min_cost << std::endl;

    best_seats.clear();
    for (int row = 0; row < maze.size(); row++) {
        for (int col = 0; col < maze[row].size(); col++) {
            maze[row][col].visited = 0;
        }
    }
    search(maze, start, 0, min_cost, best_seats);
    std::cout << "Part 2: " << best_seats.size() << std::endl;
}

void Solution::part2() {}