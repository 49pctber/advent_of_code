#include "solution.hpp"

std::filesystem::path this_file(__FILE__);
std::filesystem::path input_directory =
    this_file.parent_path().parent_path().append("input");

int main(int argc, char **argv) {

    Solution sol(argc, argv);
    sol.run();

    return 0;
}