#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <list>
#include <string>
#include <vector>

typedef std::vector<std::vector<int>> topo_map_t;
typedef struct {
    int row;
    int col;
    // int height;
} location_t;

int search(topo_map_t *map, int row, int col, int height) {
    // check bounds
    if (row < 0 || row >= (*map).size() || col < 0 || col >= (*map).size()) {
        return 0;
    }

    // check height
    if ((*map)[row][col] != height) {
        return 0;
    }

    // check end
    if (height == 9) {
        return 1;
    }

    // recursively search neighbors
    int count = 0;
    count += search(map, row - 1, col, height + 1);
    count += search(map, row + 1, col, height + 1);
    count += search(map, row, col - 1, height + 1);
    count += search(map, row, col + 1, height + 1);
    return count;
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    topo_map_t topo_map;
    while (std::getline(file, line)) {
        std::vector<int> row;
        row.reserve(line.size());
        for (auto height : line) {
            row.push_back(height - '0');
        }
        topo_map.push_back(row);
    }

    int sum = 0;
    for (int row = 0; row < topo_map.size(); row++) {
        for (int col = 0; col < topo_map.size(); col++) {
            if (topo_map[row][col] != 0) {
                continue;
            }

            // BFS
            topo_map_t search_state(topo_map);
            for (int row = 0; row < topo_map.size(); row++) {
                for (int col = 0; col < topo_map[0].size(); col++) {
                    search_state[row][col] = 0;
                }
            }

            std::list<location_t> nodes;
            nodes.push_back(location_t(row, col));

            while (nodes.size() > 0) {

                location_t loc = nodes.front();
                nodes.pop_front();
                search_state[loc.row][loc.col] |= 1 << 1; // mark visited
                int height = topo_map[loc.row][loc.col];

                if (topo_map[loc.row][loc.col] == 9) {
                    sum++;
                    continue;
                }

                if (loc.row > 0) {
                    if (topo_map[loc.row - 1][loc.col] == height + 1 &&
                        search_state[loc.row - 1][loc.col] == 0) {
                        nodes.push_back(location_t(loc.row - 1, loc.col));
                        search_state[loc.row - 1][loc.col] |=
                            1 << 0; // mark discovered
                    }
                }

                if (loc.row < topo_map.size() - 1) {
                    if (topo_map[loc.row + 1][loc.col] == height + 1 &&
                        search_state[loc.row + 1][loc.col] == 0) {
                        nodes.push_back(location_t(loc.row + 1, loc.col));
                        search_state[loc.row + 1][loc.col] |=
                            1 << 0; // mark discovered
                    }
                }

                if (loc.col > 0) {
                    if (topo_map[loc.row][loc.col - 1] == height + 1 &&
                        search_state[loc.row][loc.col - 1] == 0) {
                        nodes.push_back(location_t(loc.row, loc.col - 1));
                        search_state[loc.row][loc.col - 1] |=
                            1 << 0; // mark discovered
                    }
                }

                if (loc.col < topo_map.size() - 1) {
                    if (topo_map[loc.row][loc.col + 1] == height + 1 &&
                        search_state[loc.row][loc.col + 1] == 0) {
                        nodes.push_back(location_t(loc.row, loc.col + 1));
                        search_state[loc.row][loc.col + 1] |=
                            1 << 0; // mark discovered
                    }
                }
            }
        }
    }

    std::cout << "Part 1: " << sum << std::endl;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    topo_map_t topo_map;
    while (std::getline(file, line)) {
        std::vector<int> row;
        row.reserve(line.size());
        for (auto height : line) {
            row.push_back(height - '0');
        }
        topo_map.push_back(row);
    }

    int sum = 0;
    for (int row = 0; row < topo_map.size(); row++) {
        for (int col = 0; col < topo_map.size(); col++) {
            if (topo_map[row][col] != 0) {
                continue;
            }

            sum += search(&topo_map, row, col, 0);
                }
    }

    std::cout << "Part 2: " << sum << std::endl;
}