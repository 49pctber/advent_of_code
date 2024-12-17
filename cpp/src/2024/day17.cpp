#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

using reg_t = long int;

class VM {
  public:
    VM(std::ifstream &file) {
        reset();
        parse_file(file);
    }

    bool execute() {

        while (instruction_pointer < instructions.size()) {

            int opcode = instructions[instruction_pointer];
            instruction_pointer++;
            reg_t operand = instructions[instruction_pointer];
            instruction_pointer++;

            switch (opcode) {
            case 0:
                adv(operand);
                break;
            case 1:
                bxl(operand);
                break;
            case 2:
                bst(operand);
                break;
            case 3:
                jnz(operand);
                break;
            case 4:
                bxc(operand);
                break;
            case 5:
                out(operand);
                break;
            case 6:
                bdv(operand);
                break;
            case 7:
                cdv(operand);
                break;
            }
        }

        return self_replicating == instructions.size();
    }

    bool execute(reg_t A) {
        reset();
        regA = A;
        return execute();
    }

    void print_state() {
        std::cout << regA << ' ' << regB << ' ' << regC << ": ";
        for (int i = instruction_pointer; i < instructions.size(); i++) {
            std::cout << instructions[i] << ' ';
        }
        std::cout << '\n';
    }

    void set_print(bool val) { print = val; }

  private:
    reg_t regA;
    reg_t regB;
    reg_t regC;
    int instruction_pointer;
    std::vector<int> instructions;
    int self_replicating;
    bool print;

    void parse_file(std::ifstream &file) {
        std::string line;
        std::getline(file, line);
        regA = std::stol(line.substr(12, line.size() - 12));
        std::getline(file, line);
        regB = std::stol(line.substr(12, line.size() - 12));
        std::getline(file, line);
        regC = std::stol(line.substr(12, line.size() - 12));
        std::getline(file, line);

        std::getline(file, line);
        for (int i = 9; i < line.size(); i += 2) {
            instructions.push_back(line[i] - '0');
        }
    }

    void reset() {
        instruction_pointer = 0;
        self_replicating = 0;
        regA = 0;
        regB = 0;
        regC = 0;
        print = false;
    }

    void adv(int operand) { regA = regA / (1 << (get_combo_operand(operand))); }

    void bxl(int operand) { regB ^= operand; }

    void bst(int operand) { regB = get_combo_operand(operand) & 0b111; }

    void jnz(int operand) {
        if (regA == 0) {
            return;
        }

        instruction_pointer = operand;
    }

    void bxc(int operand) { regB ^= regC; }

    void out(int operand) {
        reg_t x = get_combo_operand(operand) & 0b111;
        if (print) {
            std::cout << x << ',';
        }

        if (self_replicating != -1 && self_replicating < instructions.size()) {
            if (x == instructions[self_replicating]) {
                self_replicating++;
            } else {
                self_replicating = -1;
            }
        }
    }

    void bdv(int operand) { regB = regA / (1 << (get_combo_operand(operand))); }

    void cdv(int operand) { regC = regA / (1 << (get_combo_operand(operand))); }

    reg_t get_combo_operand(int operand) {
        switch (operand) {
        case 4:
            return regA;

        case 5:
            return regB;

        case 6:
            return regC;

        case 7:
            std::cout << "This should never happen." << std::endl;
            return -1;

        default:
            return operand;
        }
    }
};

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    VM vm(file);

    vm.set_print(true);
    std::cout << "Part 1: ";
    vm.execute();
    std::cout << std::endl;
}

bool search(long int regA, int depth, std::vector<int> &instructions) {
    if (depth == -1) {
        std::cout << "Part 2: " << regA << std::endl;
        return true;
    }

    int print_next = instructions[depth];
    regA <<= 3;

    for (int i = 0; i < 8; i++) {
        int regB = i ^ 0b101;
        int regC = (regA + i) / (1 << regB);
        regB ^= 0b110 ^ regC;
        int out = regB & 0b111;

        if (out == print_next) {
            if (search(regA + i, depth - 1, instructions)) {
                return true;
            }
        }
    }
    return false;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    // get instructions
    std::string line;
    std::getline(file, line);
    std::getline(file, line);
    std::getline(file, line);
    std::getline(file, line);
    std::getline(file, line);
    std::vector<int> instructions;
    for (int i = 9; i < line.size(); i += 2) {
        instructions.push_back(line[i] - '0');
    }

    long int regA = 0;
    if (sizeof(regA) < 8) {
        std::cerr << "register A is too small\n";
        exit(EXIT_FAILURE);
    }
    search(0, instructions.size() - 1, instructions);
}