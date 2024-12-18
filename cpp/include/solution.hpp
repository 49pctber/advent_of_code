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
        auto stop1 = std::chrono::high_resolution_clock::now();

        part2();
        auto stop2 = std::chrono::high_resolution_clock::now();

        auto duration = stop2 - start;
        auto duration1 = stop1 - start;
        auto duration2 = stop2 - stop1;

        std::cout
            << "  Time: "
            << std::chrono::duration_cast<std::chrono::microseconds>(duration)
            << " [Part 1: "
            << std::chrono::duration_cast<std::chrono::microseconds>(duration1)
            << ", Part 2: "
            << std::chrono::duration_cast<std::chrono::microseconds>(duration2)
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
