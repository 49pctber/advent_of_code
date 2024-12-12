#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <map>
#include <vector>

// typedef std::pair<int, int> coord_t;
typedef struct {
    int row;
    int col;
} coord_t;

void print_coord(coord_t *c) {
    std::cout << "(row: " << c->row << ", col: " << c->col << ") ";
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::map<char, std::vector<coord_t>> locs;
    int row = 0;
    while (std::getline(file, line)) {
        for (int col = 0; col < line.size(); col++) {
            char c = line[col];
            if (c != '.') {
                coord_t coord{row : row, col : col};
                locs[c].push_back(coord);
            }
        }
        row++;
    }

    int n_cols, n_rows;
    n_rows = row;
    n_cols = row; // This is an assumption that isn't true in general.

    std::vector<std::vector<bool>> antinode_locs;
    antinode_locs.resize(n_rows);
    for (int i = 0; i < antinode_locs.size(); i++) {
        antinode_locs[i].resize(n_cols);
    }

    for (auto a : locs) {
        auto antenna_locs = a.second;
        for (int i = 0; i < antenna_locs.size(); i++) {

            coord_t loc1 = antenna_locs[i];
            int row1 = loc1.col;
            int col1 = loc1.row;

            for (int j = i + 1; j < antenna_locs.size(); j++) {

                coord_t loc2 = antenna_locs[j];
                int dcol = loc2.col - loc1.col;
                int drow = loc2.row - loc1.row;

                coord_t loc3{row : loc2.row + drow, col : loc2.col + dcol};
                if (loc3.row >= 0 && loc3.row < n_rows && loc3.col >= 0 &&
                    loc3.col < n_cols) {
                    antinode_locs[loc3.row][loc3.col] = true;
                }

                coord_t loc4{row : loc1.row - drow, col : loc1.col - dcol};
                if (loc4.row >= 0 && loc4.row < n_rows && loc4.col >= 0 &&
                    loc4.col < n_cols) {
                    antinode_locs[loc4.row][loc4.col] = true;
                }
            }
        }
    }

    int count = 0;
    for (auto row : antinode_locs) {
        for (auto col : row) {
            if (col) {
                count++;
            }
        }
    }
    std::cout << "Part 1: " << count << std::endl;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::map<char, std::vector<coord_t>> locs;
    int row = 0;
    while (std::getline(file, line)) {
        for (int col = 0; col < line.size(); col++) {
            char c = line[col];
            if (c != '.') {
                coord_t coord{row : row, col : col};
                locs[c].push_back(coord);
            }
        }
        row++;
    }

    int n_cols, n_rows;
    n_rows = row;
    n_cols = row; // This is an assumption that isn't true in general.

    std::vector<std::vector<bool>> antinode_locs;
    antinode_locs.resize(n_rows);
    for (int i = 0; i < antinode_locs.size(); i++) {
        antinode_locs[i].resize(n_cols);
    }

    for (auto a : locs) {
        auto antenna_locs = a.second;
        for (int i = 0; i < antenna_locs.size(); i++) {

            coord_t loc1 = antenna_locs[i];
            int row1 = loc1.col;
            int col1 = loc1.row;

            for (int j = i + 1; j < antenna_locs.size(); j++) {

                coord_t loc2 = antenna_locs[j];
                int dcol = loc2.col - loc1.col;
                int drow = loc2.row - loc1.row;

                coord_t loc3{row : loc2.row, col : loc2.col};
                while (true) {

                    if (loc3.row >= 0 && loc3.row < n_rows && loc3.col >= 0 &&
                        loc3.col < n_cols) {
                        antinode_locs[loc3.row][loc3.col] = true;
                    } else {
                        break;
                    }

                    loc3.row += drow;
                    loc3.col += dcol;
                }

                coord_t loc4{row : loc1.row, col : loc1.col};
                while (true) {

                    if (loc4.row >= 0 && loc4.row < n_rows && loc4.col >= 0 &&
                        loc4.col < n_cols) {
                        antinode_locs[loc4.row][loc4.col] = true;
                    } else {
                        break;
                    }

                    loc4.row -= drow;
                    loc4.col -= dcol;
                }
            }
        }
    }

    int count = 0;
    for (auto row : antinode_locs) {
        for (auto col : row) {
            if (col) {
                count++;
            }
        }
    }
    std::cout << "Part 2: " << count << std::endl;
}