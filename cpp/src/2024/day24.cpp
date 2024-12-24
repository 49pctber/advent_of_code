#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <map>
#include <set>
#include <string>

struct operation_t {
    std::string operand1;
    std::string operand2;
    char op;
};

int evaluate(std::string label, std::map<std::string, operation_t> &operations,
             std::map<std::string, int> &known_values) {
    if (known_values.contains(label)) {
        return known_values[label];
    }

    operation_t o = operations[label];

    int v1 = evaluate(o.operand1, operations, known_values);
    int v2 = evaluate(o.operand2, operations, known_values);

    int result;

    switch (o.op) {
    case 'A':
        result = (v1 & v2) & 0b1;
        break;

    case 'X':
        result = (v1 ^ v2) & 0b1;
        break;

    case 'O':
        result = (v1 | v2) & 0b1;
        break;

    default:
        std::cout << "something is wrong" << std::endl;
        return false;
    }

    known_values[label] = result;
    return result;
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::map<std::string, int> known_values;
    std::map<std::string, operation_t> operations;
    while (std::getline(file, line)) {
        if (line.size() == 0) {
            break;
        }

        std::string label = line.substr(0, 3);
        int value = std::stoi(line.substr(5, 1));

        known_values[label] = value;
    }

    int max_z = -1;
    while (std::getline(file, line)) {
        std::stringstream ss(line);
        std::string operand1, op, operand2, ignore, result;
        ss >> operand1 >> op >> operand2 >> ignore >> result;

        operation_t o{operand1 : operand1, operand2 : operand2, op : op[0]};

        operations[result] = o;

        if (result[0] == 'z') {
            int z = std::stoi(result.substr(1, 2));
            if (z > max_z) {
                max_z = z;
            }
        }
    }

    long int result = 0;
    for (int z = 0; z <= max_z; z++) {
        std::string label;
        if (z < 10) {
            label = "z0" + std::to_string(z);
        } else {
            label = "z" + std::to_string(z);
        }
        result |= long(evaluate(label, operations, known_values)) << z;
    }
    std::cout << "Part 1: " << result << std::endl;
}

int evaluate2(std::string label, std::map<std::string, operation_t> &operations,
              std::map<std::string, int> &known_values, int depth) {

    std::cout << std::string(4 * depth, ' ') << label << ": ";

    if (known_values.contains(label)) {
        std::cout << known_values[label] << '\n';
        return known_values[label];
    }

    operation_t o = operations[label];

    std::cout << o.operand1 << ' ' << o.op << ' ' << o.operand2 << '\n';

    int v1 = evaluate2(o.operand1, operations, known_values, depth + 1);
    int v2 = evaluate2(o.operand2, operations, known_values, depth + 1);

    int result;

    switch (o.op) {
    case 'A':
        result = (v1 & v2) & 0b1;
        break;

    case 'X':
        result = (v1 ^ v2) & 0b1;
        break;

    case 'O':
        result = (v1 | v2) & 0b1;
        break;

    default:
        std::cout << "something is wrong" << std::endl;
        return false;
    }

    known_values[label] = result;
    return result;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    std::map<std::string, int> known_values;
    std::map<std::string, operation_t> operations;
    while (std::getline(file, line)) {
        if (line.size() == 0) {
            break;
        }

        std::string label = line.substr(0, 3);
        int value = std::stoi(line.substr(5, 1));

        known_values[label] = value;
    }

    std::map<std::string, std::string> flip;
    for (int i = 2; i <= argc - 2; i += 2) {
        std::string label1(argv[i]);
        std::string label2(argv[i + 1]);
        flip[label1] = label2;
        flip[label2] = label1;
        std::cout << flip[label1] << " " << flip[label2] << std::endl;
    }

    int max_z = -1;
    while (std::getline(file, line)) {
        std::stringstream ss(line);
        std::string operand1, op, operand2, ignore, result;
        ss >> operand1 >> op >> operand2 >> ignore >> result;

        operation_t o{operand1 : operand1, operand2 : operand2, op : op[0]};

        if (flip.find(result) != flip.end()) {
            result = flip[result];
        }

        operations[result] = o;

        if (result[0] == 'z') {
            int z = std::stoi(result.substr(1, 2));
            if (z > max_z) {
                max_z = z;
            }
        }
    }

    long int input_x = 0;
    long int input_y = 0;
    for (int z = 0; z < max_z - 1; z++) {
        std::string label_x;
        std::string label_y;
        if (z < 10) {
            label_x = "x0" + std::to_string(z);
            label_y = "y0" + std::to_string(z);
        } else {
            label_x = "x" + std::to_string(z);
            label_y = "y" + std::to_string(z);
        }
        input_x |= long(known_values[label_x]) << z;
        input_y |= long(known_values[label_y]) << z;
    }
    long int desired_z = input_x + input_y;

    long int result = 0;
    for (int z = 0; z <= max_z; z++) {
        std::string label;
        if (z < 10) {
            label = "z0" + std::to_string(z);
        } else {
            label = "z" + std::to_string(z);
        }
        result |= long(evaluate2(label, operations, known_values, 0)) << z;
    }

    // std::string part2;
    // for (auto [label, other] : flip) {
    //     part2 += label + ",";
    // }
    // part2 = part2.substr(0, part2.size() - 1);
}