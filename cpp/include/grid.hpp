#pragma once

namespace aoc {

/*
Represents a two-dimensional position
*/
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

/*
Represents two-dimensional directions
*/
typedef int direction_t;
const int dir_right = 0;
const int dir_down = 1;
const int dir_left = 2;
const int dir_up = 3;

/*
Represents orientation on a grid, including both position and direction
*/
typedef struct {
    position_t position;
    direction_t direction;
} orientation_t;

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

void turn_around(orientation_t &o) { o.direction ^= 0b10; }

}; // namespace aoc
