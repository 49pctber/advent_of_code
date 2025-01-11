#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

template <typename T> void print_vector(std::vector<T> &vector) {
    for (auto x : vector) {
        std::cout << x << ' ';
    }
    std::cout << '\n';
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string token;
    std::vector<int> memory;
    while (std::getline(file, token, ',')) {
        int n = std::stoi(token);
        memory.push_back(n);
    }

    memory[1] = 12;
    memory[2] = 2;

    size_t pos = 0;
    while (true) {
        int instruction = memory[pos];

        // print_vector<int>(memory);

        switch (instruction) {
        case 1:
            memory[memory[pos + 3]] =
                memory[memory[pos + 1]] + memory[memory[pos + 2]];
            pos += 4;
            break;

        case 2:
            memory[memory[pos + 3]] =
                memory[memory[pos + 1]] * memory[memory[pos + 2]];
            pos += 4;
            break;

        case 99:
            goto finish;
            break;

        default:
            std::cerr << "This should never happen\n";
            exit(EXIT_FAILURE);
            break;
        }
    }

    // print_vector<int>(memory);

finish:
    std::cout << "Part 1: " << memory[0] << std::endl;
}

int execute(int noun, int verb, std::vector<int> memory) {

    memory[1] = noun;
    memory[2] = verb;

    size_t pos = 0;
    while (true) {
        int instruction = memory[pos];

        switch (instruction) {
        case 1:
            memory[memory[pos + 3]] =
                memory[memory[pos + 1]] + memory[memory[pos + 2]];
            pos += 4;
            break;

        case 2:
            memory[memory[pos + 3]] =
                memory[memory[pos + 1]] * memory[memory[pos + 2]];
            pos += 4;
            break;

        case 99:
            return memory[0];

        default:
            std::cerr << "This should never happen\n";
            exit(EXIT_FAILURE);
            break;
        }
    }
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string token;
    std::vector<int> memory;
    while (std::getline(file, token, ',')) {
        int n = std::stoi(token);
        memory.push_back(n);
    }

    for (int noun = 0; noun < 100; noun++) {
        for (int verb = 0; verb < 100; verb++) {
            if (execute(noun, verb, memory) == 19690720) {
                std::cout << "Part 2: " << 100 * noun + verb << std::endl;
                break;
            }
        }
    }
}
