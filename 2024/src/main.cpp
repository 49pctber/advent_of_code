#include "solution.hpp"

std::filesystem::path this_file(__FILE__);
std::filesystem::path input_directory =
    this_file.parent_path().parent_path().append("input");

int main(int argc, char **argv) {

    Solution sol;
    switch (argc) {
    case 1:
        sol.run();
        break;
    case 2:
        int part = std::atoi(argv[1]);
        switch (part) {
        case 1:
            sol.part1();
            break;
        case 2:
            sol.part2();
            break;
        default:
            break;
        }
    }

    return 0;
}