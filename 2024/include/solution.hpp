#pragma once

#include <filesystem>
#include <fstream>
#include <iostream>

extern std::filesystem::path this_file;
extern std::filesystem::path input_directory;

class Solution {
  public:
    void run() {
        part1();
        part2();
    }
    void part1();
    void part2();
};
