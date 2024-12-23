#include "solution.hpp"
#include <algorithm>
#include <vector>

class Seat {
  public:
    void parseString(std::string input) {
        row = 0;
        col = 0;

        if (input.length() != 10) {
            return;
        }

        for (int i = 0; i < 7; i++) {
            row <<= 1;
            if (input[i] == 'B') {
                row |= 0b1;
            }
        }

        for (int i = 7; i < 10; i++) {
            col <<= 1;
            if (input[i] == 'R') {
                col |= 0b1;
            }
        }
    };

    int seatId() { return row * 8 + col; }

  private:
    int row;
    int col;
};

void Solution::part1() {
    std::ifstream input(argv[1]);
    std::string line;
    int max = INT32_MIN;

    while (input >> line) {
        Seat seat;
        seat.parseString(line);
        int sid = seat.seatId();
        if (sid > max) {
            max = sid;
        }
    }

    std::cout << "Part 1: " << max << std::endl;
}

void Solution::part2() {
    std::ifstream input(argv[1]);
    std::string line;
    std::vector<int> seat_ids;

    while (input >> line) {
        Seat seat;
        seat.parseString(line);
        int sid = seat.seatId();
        seat_ids.push_back(sid);
    }

    std::sort(seat_ids.begin(), seat_ids.end());
    int min = seat_ids[0];

    for (int i = 0; i < seat_ids.size(); i++) {
        if (seat_ids[i] != min + i) {
            std::cout << "Part 2: " << min + i << std::endl;
            return;
        }
    }
}
