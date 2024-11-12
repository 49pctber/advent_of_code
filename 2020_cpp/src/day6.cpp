#include "solution.hpp"
#include <map>
#include <set>

std::filesystem::path input_path = input_directory.append("6.txt");

void Solution::part1() {
    std::ifstream input(input_path);
    std::string line;
    std::set<char> yeses;

    int sum = 0;

    while (std::getline(input, line)) {
        if (line.length() == 0) {
            sum += yeses.size();
            yeses.clear();
        }
        for (char c : line) {
            yeses.insert(c);
        }
    }

    sum += yeses.size();

    std::cout << "Part 1: " << sum << std::endl;
}

void Solution::part2() {
    std::ifstream input(input_path);
    std::string line;
    std::map<char, int> yeses;

    int sum = 0;
    int n_rows = 0;

    while (std::getline(input, line)) {
        if (line.length() == 0) {
            for (auto const &[key, val] : yeses) {
                if (val == n_rows) {
                    sum++;
                }
            }
            yeses.clear();
            n_rows = 0;
        } else {
            n_rows++;
            for (char c : line) {
                yeses[c] += 1;
            }
        }
    }
    std::cout << n_rows << '\n';
    for (auto const &[key, val] : yeses) {
        std::cout << key << ' ' << val << '\n';
        if (val == n_rows) {
            sum++;
        }
    }

    std::cout << "Part 2: " << sum << std::endl;
}
