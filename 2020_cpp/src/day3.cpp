#include "solution.hpp"

std::filesystem::path input_path = input_directory.append("3.txt");

int trees(int dx) {

    std::ifstream file(input_path);
    if (!file.is_open()) {
        std::cerr << "Failed to open input file" << std::endl;
    }

    std::string line;
    int x = 0;
    int n_trees_hit = 0;
    while (file >> line) {
        if (line[x] == '#') {
            n_trees_hit++;
        }

        x += dx;
        x %= line.length();
    }
    file.close();
    return n_trees_hit;
}

int trees(int dx, int dy) {

    std::ifstream file(input_path);
    if (!file.is_open()) {
        std::cerr << "Failed to open input file" << std::endl;
    }

    std::string line;
    int x = 0;
    int y = 0;
    int n_trees_hit = 0;
    while (file >> line) {
        if (y % dy == 0) {
            if (line[x] == '#') {
                n_trees_hit++;
            }

            x += dx;
            x %= line.length();
        }
        y++;
    }
    file.close();
    return n_trees_hit;
}

void Solution::part1() { std::cout << "Part 1: " << trees(3) << std::endl; }

void Solution::part2() {
    long int product = 1;
    product *= trees(1);
    product *= trees(3);
    product *= trees(5);
    product *= trees(7);
    product *= trees(1, 2);
    std::cout << "Part 2: " << product << std::endl;
}
