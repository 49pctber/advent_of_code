#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <list>
#include <map>
#include <set>
#include <string>
#include <vector>

const int right = 0b00;
const int below = 0b01;
const int left = 0b10;
const int above = 0b11;

struct Plot {
    char type;
    int row;
    int col;
    std::map<int, Plot *> neighbor;
    bool state; // for storing search state

    bool operator<(const Plot &other) const {
        if (row == other.row) {
            return col < other.col;
        }
        return row < other.row;
    }
};

typedef struct Plot plot_t;

typedef struct {
    char type;
    std::set<plot_t> plots;
    plot_t *start_plot;
    int perimeter;
    int area;

    int n_sides() { return n_corners(); }

    int n_corners() {
        int n = 0;

        std::set<std::pair<int, int>> locs;

        for (auto plot : plots) {
            locs.insert(std::pair<int, int>(plot.row, plot.col));
        }

        for (auto loc : locs) {
            // exterior corner
            if (!locs.contains(
                    std::pair<int, int>(loc.first + 1, loc.second)) &&
                !locs.contains(
                    std::pair<int, int>(loc.first, loc.second + 1))) {
                n++;
            }
            if (!locs.contains(
                    std::pair<int, int>(loc.first - 1, loc.second)) &&
                !locs.contains(
                    std::pair<int, int>(loc.first, loc.second + 1))) {
                n++;
            }
            if (!locs.contains(
                    std::pair<int, int>(loc.first + 1, loc.second)) &&
                !locs.contains(
                    std::pair<int, int>(loc.first, loc.second - 1))) {
                n++;
            }
            if (!locs.contains(
                    std::pair<int, int>(loc.first - 1, loc.second)) &&
                !locs.contains(
                    std::pair<int, int>(loc.first, loc.second - 1))) {
                n++;
            }

            // interior corner
            if (locs.contains(std::pair<int, int>(loc.first + 1, loc.second)) &&
                locs.contains(std::pair<int, int>(loc.first, loc.second + 1)) &&
                !locs.contains(
                    std::pair<int, int>(loc.first + 1, loc.second + 1))) {
                n++;
            }
            if (locs.contains(std::pair<int, int>(loc.first - 1, loc.second)) &&
                locs.contains(std::pair<int, int>(loc.first, loc.second + 1)) &&
                !locs.contains(
                    std::pair<int, int>(loc.first - 1, loc.second + 1))) {
                n++;
            }
            if (locs.contains(std::pair<int, int>(loc.first + 1, loc.second)) &&
                locs.contains(std::pair<int, int>(loc.first, loc.second - 1)) &&
                !locs.contains(
                    std::pair<int, int>(loc.first + 1, loc.second - 1))) {
                n++;
            }
            if (locs.contains(std::pair<int, int>(loc.first - 1, loc.second)) &&
                locs.contains(std::pair<int, int>(loc.first, loc.second - 1)) &&
                !locs.contains(
                    std::pair<int, int>(loc.first - 1, loc.second - 1))) {
                n++;
            }
        }

        return n;
    }

    int n_exterior_sides() {
        // only gives the number of sides on the exterior
        // if this region surrounds other regions, the number of sides inside
        // need to be added
        plot_t *current = start_plot;
        int dir = right;
        int n = 0;

        while (true) {
            // check in dir-1
            // corner
            if (current->neighbor.find((dir - 1) & 0b11) !=
                current->neighbor.end()) {
                dir = (dir - 1) & 0b11;
                n++;
                current = current->neighbor[dir];
                continue;
            }

            // check in dir
            // no corner
            if (current->neighbor.find(dir) != current->neighbor.end()) {
                current = current->neighbor[dir];
                continue;
            }

            n++;
            dir++;
            dir &= 0b11;

            if (current == start_plot && dir == right) {
                break;
            }
        }

        return n;
    }
} region_t;

typedef std::vector<std::vector<plot_t>> farm_t;

int get_fence_cost(region_t &region) { return region.area * region.perimeter; }

int get_discount_cost(region_t &region) {
    return region.area * region.n_sides();
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    farm_t farm;
    std::string line;
    while (std::getline(file, line)) {
        std::vector<plot_t> row;
        row.reserve(line.size());
        int col_no = 0;
        for (auto c : line) {
            row.push_back(plot_t{
                type : c,
                row : int(farm.size()),
                col : col_no,
                state : 0
            });
            col_no++;
        }
        farm.push_back(row);
    }

    // find regions
    std::list<region_t> regions;
    for (int row = 0; row < farm.size(); row++) {
        for (int col = 0; col < farm[0].size(); col++) {

            if (farm[row][col].state != 0) {
                continue; // already added to a region
            }

            // create new region
            region_t region;
            region.type = farm[row][col].type;
            region.area = 0;
            region.perimeter = 0;

            // initialize search queue
            std::list<plot_t *> search_queue;
            search_queue.push_back(&(farm[row][col]));
            farm[row][col].state = 0b1; // mark discovered
            region.start_plot = &(farm[row][col]);

            // BFS for region plots
            while (search_queue.size() > 0) {

                // get next discovered plot
                plot_t *current_plot = search_queue.front();
                search_queue.pop_front();
                current_plot->state |= 0b10; // mark visited;

                // add to region
                region.plots.insert(*current_plot);
                region.area++;
                region.perimeter += 4;

                // look for new plots
                if (current_plot->row > 0) {
                    plot_t *candidate =
                        &(farm[current_plot->row - 1][current_plot->col]);
                    if (candidate->type == region.type) {
                        region.perimeter--;
                        current_plot->neighbor[above] = candidate;
                        if (candidate->state == 0) {
                            search_queue.push_back(candidate);
                            candidate->state |= 0b1;
                        }
                    }
                }

                if (current_plot->row < farm.size() - 1) {
                    plot_t *candidate =
                        &(farm[current_plot->row + 1][current_plot->col]);
                    if (candidate->type == region.type) {
                        region.perimeter--;
                        current_plot->neighbor[below] = candidate;
                        if (candidate->state == 0) {
                            search_queue.push_back(candidate);
                            candidate->state |= 0b1;
                        }
                    }
                }

                if (current_plot->col > 0) {
                    plot_t *candidate =
                        &(farm[current_plot->row][current_plot->col - 1]);
                    if (candidate->type == region.type) {
                        region.perimeter--;
                        current_plot->neighbor[left] = candidate;
                        if (candidate->state == 0) {
                            search_queue.push_back(candidate);
                            candidate->state |= 0b1;
                        }
                    }
                }

                if (current_plot->col < farm[0].size() - 1) {
                    plot_t *candidate =
                        &(farm[current_plot->row][current_plot->col + 1]);
                    if (candidate->type == region.type) {
                        region.perimeter--;
                        current_plot->neighbor[right] = candidate;
                        if (candidate->state == 0) {
                            search_queue.push_back(candidate);
                            candidate->state |= 0b1;
                        }
                    }
                }
            }

            regions.push_back(region);
        }
    }

    // compute cumulative cost
    long int cost = 0;
    long int discount = 0;
    for (auto region : regions) {
        cost += get_fence_cost(region);
        discount += get_discount_cost(region);
    }

    std::cout << "Part 1: " << cost << std::endl;
    std::cout << "Part 2: " << discount << std::endl;
}

void Solution::part2() {}
