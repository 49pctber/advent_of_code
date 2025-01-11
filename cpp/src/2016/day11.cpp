#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <queue>
#include <regex>
#include <set>
#include <string>
#include <vector>

const int hydrogen = 0;  // hydrogen (H 1) for testing
const int lithium = 1;   // lithium (Li 3) for testing
const int strontium = 0; // strontium (Sr 38)
const int ruthenium = 1; // ruthenium (Ru 44)
const int thulium = 2;   // thulium (Tm 69)
const int plutonium = 3; // plutonium (Pu 94)
const int curium = 4;    // curium (Cm 96)
const int elerium = 5;
const int dilithium = 6;

const int max_elements = 7;
const int element_mask = (0b1 << max_elements) - 1;
const int n_floors = 4;

using floors_t = std::vector<int>;

bool valid_floor_state(int state) {
    int microchips = state & element_mask;
    if (microchips == 0) {
        return true;
    }

    int generators = (state >> max_elements) & element_mask;
    if (generators == 0) {
        return true;
    }

    int unpowered_microchips = microchips & ~generators;

    return !(generators != 0 && unpowered_microchips != 0);
}

typedef struct {
    int elevator_floor;
    int n_steps;
    floors_t floors;
} search_state_t;

void print(search_state_t &state) {
    for (int i = n_floors - 1; i >= 0; i--) {
        std::cout << ((i == state.elevator_floor) ? 'E' : '_');
        for (int j = 2 * max_elements - 1; j >= 0; j--) {
            std::cout << ((state.floors[i] >> j) & 0b1);
        }
        std::cout << '\n';
    }
}

bool can_move_up(search_state_t &state, int flip) {
    if (state.elevator_floor >= n_floors - 1) {
        return false;
    }

    return valid_floor_state(state.floors[state.elevator_floor] ^ flip) &&
           valid_floor_state(state.floors[state.elevator_floor + 1] ^ flip);
}

search_state_t move_up(search_state_t state, int flip) {
    if (!can_move_up(state, flip)) {
        exit(EXIT_FAILURE);
    }

    auto next(state);
    next.floors[next.elevator_floor] ^= flip;
    next.elevator_floor++;
    next.n_steps++;
    next.floors[next.elevator_floor] ^= flip;

    return next;
}

bool can_move_down(search_state_t &state, int flip) {
    if (state.elevator_floor <= 0) {
        return false;
    }

    return valid_floor_state(state.floors[state.elevator_floor] ^ flip) &&
           valid_floor_state(state.floors[state.elevator_floor - 1] ^ flip);
}

search_state_t move_down(search_state_t state, int flip) {
    if (!can_move_down(state, flip)) {
        exit(EXIT_FAILURE);
    }

    auto next(state);
    next.floors[next.elevator_floor] ^= flip;
    next.elevator_floor--;
    next.floors[next.elevator_floor] ^= flip;
    next.n_steps++;

    return next;
}

long int hash(search_state_t &state) {
    long int digest = 0; // non cryptographic, obviously
    for (int i = 0; i < n_floors; i++) {
        digest <<= 2 * max_elements;
        digest ^= state.floors[i];
    }
    digest <<= 2;
    digest ^= state.elevator_floor;
    // digest <<= 16;
    // digest ^= state.n_steps;

    return digest;
}

int element_to_index(std::string element) {
    if (element == "strontium") {
        return strontium;
    } else if (element == "ruthenium") {
        return ruthenium;
    } else if (element == "thulium") {
        return thulium;
    } else if (element == "plutonium") {
        return plutonium;
    } else if (element == "curium") {
        return curium;
    } else if (element == "hydrogen") {
        return hydrogen;
    } else if (element == "lithium") {
        return lithium;
    } else if (element == "dilithium") {
        return dilithium;
    } else if (element == "elerium") {
        return elerium;
    } else {
        exit(EXIT_FAILURE);
    }
}

int element_to_bitmask(std::string element) {
    return 1 << element_to_index(element);
}

void Solution::part1() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    int floor = 0;

    search_state_t start;
    start.elevator_floor = 0;
    start.n_steps = 0;
    start.floors.resize(n_floors, 0);

    int end_state = 0;
    std::regex pattern(R"(a (\w+) generator|a (\w+)-compatible microchip)");
    while (std::getline(file, line)) {
        for (std::sregex_iterator it(line.begin(), line.end(), pattern), end;
             it != end; ++it) {
            if ((*it)[1].length() > 0) {
                start.floors[floor] |= element_to_bitmask((*it)[1])
                                       << max_elements;
                end_state |= element_to_bitmask((*it)[1]) << max_elements;
            } else {
                start.floors[floor] |= element_to_bitmask((*it)[2]);
                end_state |= element_to_bitmask((*it)[2]);
            }
        }

        floor++;
    }

    std::queue<search_state_t> search_queue;
    search_queue.push(start);
    std::set<long int>
        seen_states; // maps states to shortest known path to get there

    while (!search_queue.empty()) {
        search_state_t state = search_queue.front();
        search_queue.pop();

        if (state.floors[n_floors - 1] == end_state) {
            std::cout << "Part 1: " << state.n_steps << std::endl;
            break;
        }

        // something like this might introduce subtle bugs
        long int state_hash = hash(state);
        if (seen_states.contains(state_hash)) {
            continue;
        } else {
            seen_states.insert(state_hash);
        }

        // find elements you can move from the current floor
        std::vector<int> options;
        for (int i = 0; i < 2 * max_elements; i++) {
            if ((state.floors[state.elevator_floor] >> i) & 0b1) {
                options.push_back(i);
            }
        }

        // try moving every combination of generator and chip on floor
        for (size_t i = 0; i < options.size(); i++) {

            for (size_t j = i; j < options.size(); j++) {

                int flip = (1 << options[i]) | (1 << options[j]);

                if (can_move_up(state, flip)) {
                    auto next = move_up(state, flip);
                    search_queue.push(next);
                }

                if (can_move_down(state, flip)) {
                    auto next = move_down(state, flip);
                    search_queue.push(next);
                }
            }
        }
    }
}

void Solution::part2() {
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }

    std::string line;
    int floor = 0;

    search_state_t start;
    start.elevator_floor = 0;
    start.n_steps = 0;
    start.floors.resize(n_floors, 0);

    int end_state = 0;
    std::regex pattern(R"(a (\w+) generator|a (\w+)-compatible microchip)");
    while (std::getline(file, line)) {
        for (std::sregex_iterator it(line.begin(), line.end(), pattern), end;
             it != end; ++it) {
            if ((*it)[1].length() > 0) {
                start.floors[floor] |= element_to_bitmask((*it)[1])
                                       << max_elements;
                end_state |= element_to_bitmask((*it)[1]) << max_elements;
            } else {
                start.floors[floor] |= element_to_bitmask((*it)[2]);
                end_state |= element_to_bitmask((*it)[2]);
            }
        }
        floor++;
    }

    start.floors[0] |= element_to_bitmask("elerium") << max_elements;
    end_state |= element_to_bitmask("elerium") << max_elements;
    start.floors[0] |= element_to_bitmask("dilithium") << max_elements;
    end_state |= element_to_bitmask("dilithium") << max_elements;
    start.floors[0] |= element_to_bitmask("elerium");
    end_state |= element_to_bitmask("elerium");
    start.floors[0] |= element_to_bitmask("dilithium");
    end_state |= element_to_bitmask("dilithium");

    std::queue<search_state_t> search_queue;
    search_queue.push(start);
    std::set<long int>
        seen_states; // maps states to shortest known path to get there

    while (!search_queue.empty()) {
        search_state_t state = search_queue.front();
        search_queue.pop();

        if (state.floors[n_floors - 1] == end_state) {
            std::cout << "Part 2: " << state.n_steps << std::endl;
            break;
        }

        long int state_hash = hash(state);
        if (seen_states.contains(state_hash)) {
            continue;
        } else {
            seen_states.insert(state_hash);
        }

        // find elements you can move from the current floor
        std::vector<int> options;
        for (int i = 0; i < 2 * max_elements; i++) {
            if ((state.floors[state.elevator_floor] >> i) & 0b1) {
                options.push_back(i);
            }
        }

        // try moving every combination of generator and chip on floor
        for (size_t i = 0; i < options.size(); i++) {

            for (size_t j = i; j < options.size(); j++) {

                int flip = (1 << options[i]) | (1 << options[j]);

                if (can_move_up(state, flip)) {
                    auto next = move_up(state, flip);
                    search_queue.push(next);
                }

                if (can_move_down(state, flip)) {
                    auto next = move_down(state, flip);
                    search_queue.push(next);
                }
            }
        }
    }
}