#include "solution.hpp"
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

void Solution::part1() {
    // get input string
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }
    std::string input;
    std::getline(file, input);

    // parse input
    long int len = 0;
    std::vector<int> disk;
    disk.reserve(input.size());
    for (char c : input) {
        int n = c - '0';
        len += n;
        disk.push_back(n);
    }

    // calculate the amount of used space
    long int stop_pos = 0;
    for (int i = 0; i < disk.size(); i += 2) {
        stop_pos += disk[i];
    }

    // initialize checksum computation
    long int checksum = 0;
    long int i = 0; // front location
    long int front_file_no = 0;

    long int mode_countdown = disk[i];
    bool back_mode = false;

    long int j = disk.size() - 1; // back location
    long int back_file_no = disk.size() / 2;
    long int back_countdown = disk[j];

    for (long int pos = 0; pos < stop_pos; pos++) {

        if (mode_countdown == 0) {
            if (!back_mode) {
                front_file_no++;
            }
            i++;
            mode_countdown = disk[i];
            back_mode = !back_mode;
            pos--;
            continue;
        } else if (mode_countdown < 0) {
            std::cout << "this should never happen" << std::endl;
            exit(EXIT_FAILURE);
        }

        if (back_mode) {
            checksum += pos * back_file_no;
            back_countdown--;
            if (back_countdown == 0) {
                j -= 2;
                back_countdown = disk[j];
                back_file_no--;
            }
        } else {
            checksum += pos * front_file_no;
        }

        mode_countdown--;
    }

    std::cout << "Part 1: " << checksum << std::endl;
}

typedef struct {
    long int number;
    int file_size;
    int space_after;
} file_descriptor_t;

typedef std::vector<file_descriptor_t> disk_map_t;

long int disk_checksum(disk_map_t *disk_map) {
    long int checksum = 0;
    int pos = 0;

    for (auto file : *disk_map) {
        for (int i = 0; i < file.file_size; i++) {
            checksum += pos * file.number;
            pos++;
        }
        pos += file.space_after;
    }

    return checksum;
}

void Solution::part2() {
    // get input string
    std::ifstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "error opening file\n";
        exit(EXIT_FAILURE);
    }
    std::string input;
    std::getline(file, input);

    // parse input
    disk_map_t disk_map;
    int file_no = 0;
    for (int i = 0; i < input.size(); i += 2) {
        disk_map.push_back(file_descriptor_t{
            number : file_no,
            file_size : input[i] - '0',
            space_after : (i + 1 < input.size()) ? (input[i + 1] - '0') : 0,
        });
        file_no++;
    }

    for (int fn = file_no - 1; fn >= 0; fn--) {
        int j;
        for (int k = 0; k < disk_map.size(); k++) {
            if (disk_map[k].number == fn) {
                j = k;
                break;
            }
        }

        for (int i = 0; i < j; i++) {
            if (disk_map[j].file_size <= disk_map[i].space_after) {

                file_descriptor_t new_descriptor{
                    number : disk_map[j].number,
                    file_size : disk_map[j].file_size,
                    space_after : disk_map[i].space_after -
                        disk_map[j].file_size,
                };

                if (i < j - 1) {

                    disk_map[i].space_after = 0;
                    disk_map[j - 1].space_after +=
                        disk_map[j].file_size + disk_map[j].space_after;
                    disk_map.erase(disk_map.begin() + j);
                    disk_map.insert(disk_map.begin() + i + 1, new_descriptor);
                } else { // i == j - 1
                    disk_map[j].space_after += disk_map[i].space_after;
                    disk_map[i].space_after = 0;
                }

                break;
            }
        }
    }

    std::cout << "Part 2: " << disk_checksum(&disk_map) << std::endl;
}
