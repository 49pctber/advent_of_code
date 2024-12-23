#include "solution.hpp"
#include <algorithm>
#include <fstream>
#include <iostream>
#include <map>
#include <set>
#include <stack>
#include <string>
#include <vector>

struct node_t {
    std::string label;
    int depth;
    int search_state;
    std::set<std::string> edges;

    void mark_discovered() { search_state |= 0b1; }

    bool is_discovered() { return search_state & 0b1; }

    void mark_visited() { search_state |= 0b10; }

    bool is_visited() { return search_state & 0b10; }
};

bool operator<(const node_t &a, const node_t &other) {
    return a.label < other.label;
}

std::ostream &operator<<(std::ostream &os, const node_t &p) {
    os << p.label;
    return os;
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::map<std::string, node_t> nodes;

    std::string line;
    while (std::getline(file, line)) {

        std::string label1 = line.substr(0, 2);
        std::string label2 = line.substr(3, 2);

        if (!nodes.contains(label1)) {
            nodes[label1] = node_t{label : label1, depth : 0, search_state : 0};
        }

        if (!nodes.contains(label2)) {
            nodes[label2] = node_t{label : label2, depth : 0, search_state : 0};
        }

        nodes[label1].edges.insert(label2);
        nodes[label2].edges.insert(label1);
    }

    // DFS
    long int count = 0;
    std::stack<std::string> search_stack;
    std::string start_label = nodes.begin()->first;
    search_stack.push(start_label);
    nodes[start_label].depth = 0;
    nodes[start_label].mark_discovered();

    while (!search_stack.empty()) {
        node_t *current = &nodes[search_stack.top()];
        search_stack.pop();
        current->mark_visited();

        for (auto neighbor_label : current->edges) {
            node_t *neighbor = &nodes[neighbor_label];
            if (!neighbor->is_discovered()) {
                neighbor->mark_discovered();
                neighbor->depth = current->depth + 1;
                search_stack.push(neighbor_label);
            }
        }

        // look for cliques of degree 3
        for (auto neighbor_label : current->edges) {
            if (neighbor_label < current->label) {
                for (auto third_label : current->edges) {
                    if (third_label < neighbor_label) {
                        // check if third_label is neighbor of neighbor_label
                        if (nodes[neighbor_label].edges.contains(third_label)) {
                            if (current->label[0] == 't' ||
                                neighbor_label[0] == 't' ||
                                third_label[0] == 't') {
                                count++;
                            }
                        }
                    }
                }
            }
        }
    }

    std::cout << "Part 1: " << count << std::endl;
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::map<std::string, node_t> nodes;

    std::string line;
    while (std::getline(file, line)) {

        std::string label1 = line.substr(0, 2);
        std::string label2 = line.substr(3, 2);

        if (!nodes.contains(label1)) {
            nodes[label1] = node_t{label : label1, depth : 0, search_state : 0};
        }

        if (!nodes.contains(label2)) {
            nodes[label2] = node_t{label : label2, depth : 0, search_state : 0};
        }

        nodes[label1].edges.insert(label2);
        nodes[label2].edges.insert(label1);
    }

    // DFS
    std::stack<std::string> search_stack;
    std::string start_label = nodes.begin()->first;
    search_stack.push(start_label);
    nodes[start_label].depth = 0;
    nodes[start_label].mark_discovered();

    std::set<std::string> largest_clique;

    while (!search_stack.empty()) {
        node_t *current = &nodes[search_stack.top()];
        search_stack.pop();
        current->mark_visited();

        for (auto neighbor_label : current->edges) {
            node_t *neighbor = &nodes[neighbor_label];
            if (!neighbor->is_discovered()) {
                neighbor->mark_discovered();
                neighbor->depth = current->depth + 1;
                search_stack.push(neighbor_label);
            }
        }

        // find clique
        std::set<std::string> remaining_edges = current->edges;

        while (!remaining_edges.empty()) {
            std::set<std::string> clique;
            clique.insert(current->label);

            std::set<std::string> candidates = remaining_edges;

            if (clique.size() + remaining_edges.size() <=
                largest_clique.size()) {
                break;
            }

            for (auto candidate : candidates) {
                clique.insert(candidate);
                remaining_edges.erase(candidate);

                // remove any candidates that aren't also a neighbor of
                // candidate
                for (auto r : remaining_edges) {
                    if (!nodes[candidate].edges.contains(r))
                        candidates.erase(r);
                }
            }

            // check if largest clique
            if (clique.size() > largest_clique.size()) {
                largest_clique = clique;
            }
        }
    }

    std::string password;

    for (auto label : largest_clique) {
        password += label + ",";
    }
    password = password.substr(0, password.size() - 1);

    std::cout << "Part 2: " << password << std::endl;
}