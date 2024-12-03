#pragma once

#include <chrono>
#include <filesystem>
#include <fstream>
#include <iostream>

extern std::filesystem::path this_file;
extern std::filesystem::path input_directory;

class Solution {
  public:
    void run() {
        auto start = std::chrono::high_resolution_clock::now();

        part1();
        part2();

        auto stop = std::chrono::high_resolution_clock::now();
        auto duration = stop - start;

        std::cout << "["
                  << std::chrono::duration_cast<std::chrono::microseconds>(
                         duration)
                  << "]" << std::endl;
    }

    void part1();

    void part2();

    Solution(int argc, char **argv) {
        this->argc = argc;
        this->argv = argv;
    }

  private:
    int argc;
    char **argv;
};
