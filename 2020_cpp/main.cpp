#include <algorithm>
#include <filesystem>
#include <fstream>
#include <iostream>
#include <regex>
#include <set>
#include <string>
#include <vector>

std::filesystem::path this_file(__FILE__);
std::filesystem::path input_directory = this_file.parent_path().append("input");

class Solution {
  public:
    void run() {
        part1();
        part2();
    }
    virtual void part1() {
        std::cout << "[Part 1 not implemented]" << std::endl;
    };
    virtual void part2() {
        std::cout << "[Part 2 not implemented]" << std::endl;
    };
};

class Day1 : public Solution {
  public:
    const int target = 2020;
    std::filesystem::path input_path = input_directory.append("1.txt");

    void part1() {

        std::ifstream input(input_path);
        if (!input.is_open()) {
            std::cerr << "Failed to open input file" << std::endl;
        }

        std::string line;

        std::vector<int> entries;
        while (std::getline(input, line)) {
            int entry = std::stoi(line);
            entries.push_back(entry);
        }
        input.close();

        // sort to convert O(n^2) to O(n)
        std::sort(entries.begin(), entries.end());

        int low, high;
        low = 0;
        high = entries.size() - 1;

        while (low < high) {
            int sum = entries[low] + entries[high];
            if (sum > target) {
                high--;
            } else if (sum < target) {
                low++;
            } else {
                int product = entries[low] * entries[high];
                std::cout << "Part 1: " << product << std::endl;
                break;
            }
        }
    }

    void part2() {
        std::ifstream input(input_path);
        if (!input.is_open()) {
            std::cerr << "Failed to open input file" << std::endl;
        }

        std::string line;

        std::vector<int> entries;
        std::set<int> entry_set;
        while (std::getline(input, line)) {
            int entry = std::stoi(line);
            entries.push_back(entry);
            entry_set.insert(entry);
        }
        input.close();

        for (int low = 0; low < entries.size() - 2; low++) {
            for (int mid = low + 1; mid < entries.size() - 1; mid++) {
                int diff = target - entries[low] - entries[mid];
                if (entry_set.contains(diff)) {
                    int product = entries[low] * entries[mid] * diff;
                    std::cout << "Part 2: " << product << std::endl;
                    return;
                }
            }
        }
    }
};

class Day2 : public Solution {
  public:
    std::filesystem::path input_path = input_directory.append("2.txt");

    void part1() {
        std::ifstream input(input_path);
        if (!input.is_open()) {
            std::cerr << "Failed to open input file" << std::endl;
        }

        std::string line;
        int n_valid_passwords = 0;
        while (std::getline(input, line)) {
            std::regex pattern(R"((\d+)-(\d+) (\w): (\w+))");
            std::smatch matches;
            if (std::regex_search(line, matches, pattern)) {
                int min = std::stoi(matches[1]);
                int max = std::stoi(matches[2]);
                char c = std::string(matches[3])[0];
                std::string password = matches[4];
                int count = 0;
                for (char d : password) {
                    if (d == c) {
                        count++;
                    }
                }
                if (count <= max && count >= min) {
                    n_valid_passwords++;
                }
            } else {
                std::cout << "No matches found for " << line << '\n';
            }
        }
        input.close();
        std::cout << "Part 1: " << n_valid_passwords << std::endl;
    }

    void part2() {
        std::ifstream input(input_path);
        if (!input.is_open()) {
            std::cerr << "Failed to open input file" << std::endl;
        }

        std::string line;
        int n_valid_passwords = 0;
        while (std::getline(input, line)) {
            std::regex pattern(R"((\d+)-(\d+) (\w): (\w+))");
            std::smatch matches;
            if (std::regex_search(line, matches, pattern)) {
                int pos1 = std::stoi(matches[1]) - 1;
                int pos2 = std::stoi(matches[2]) - 1;
                char c = std::string(matches[3])[0];
                std::string password = matches[4];
                if (password[pos1] == c ^ password[pos2] == c) {
                    n_valid_passwords++;
                }
            } else {
                std::cout << "No matches found for " << line << '\n';
            }
        }
        input.close();
        std::cout << "Part 2: " << n_valid_passwords << std::endl;
    }
};

class Day3 : public Solution {
  public:
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

    void part1() { std::cout << "Part 1: " << trees(3) << std::endl; }

    void part2() {
        long int product = 1;
        product *= trees(1);
        product *= trees(3);
        product *= trees(5);
        product *= trees(7);
        product *= trees(1, 2);
        std::cout << "Part 2: " << product << std::endl;
    }
};

int main(int argc, char **argv) {
    if (argc == 1) {
        std::cerr << "Select a valid day" << std::endl;
        return -1;
    }

    int day = std::atoi(argv[1]);
    Solution *sol;
    switch (day) {
    case 1:
        sol = new Day1();
        break;
    case 2:
        sol = new Day2();
        break;
    case 3:
        sol = new Day3();
        break;
    default:
        std::cerr << "Select a valid day" << std::endl;
        return -1;
    }

    std::cout << "Day " << day << '\n';

    if (argc == 3) {
        int part = std::atoi(argv[2]);
        switch (part) {
        case 1:
            sol->part1();
            break;
        case 2:
            sol->part2();
            break;
        default:
            break;
        }
    } else {
        sol->run();
    }

    return 0;
}