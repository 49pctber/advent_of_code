#include "solution.hpp"
#include <regex>

#define INVALID_PASSPORT_FIELD (-1)

std::filesystem::path input_path = input_directory.append("4.txt");

class Passport {
  public:
    Passport()
        : byr(INVALID_PASSPORT_FIELD), iyr(INVALID_PASSPORT_FIELD),
          eyr(INVALID_PASSPORT_FIELD), cid(INVALID_PASSPORT_FIELD) {}

    void parseString(std::string line) {
        std::regex pattern(R"((\w{3}):(#?\w+))");
        std::smatch matches;

        auto matches_begin =
            std::sregex_iterator(line.begin(), line.end(), pattern);
        auto matches_end = std::sregex_iterator();

        for (std::sregex_iterator i = matches_begin; i != matches_end; ++i) {
            std::smatch match = *i;
            if (match[1] == "byr") {
                byr = std::stoi(match[2]);
            } else if (match[1] == "iyr") {
                iyr = std::stoi(match[2]);
            } else if (match[1] == "eyr") {
                eyr = std::stoi(match[2]);
            } else if (match[1] == "hcl") {
                hcl = match[2];
            } else if (match[1] == "hgt") {
                hgt = match[2];
            } else if (match[1] == "ecl") {
                ecl = match[2];
            } else if (match[1] == "pid") {
                pid = match[2];
            } else if (match[1] == "cid") {
                cid = std::stol(match[2]);
            } else {
                std::cout << "AHH! ERRORRORORORO" << std::endl;
            }
        }
    }

    bool isValid() {
        return byr != INVALID_PASSPORT_FIELD && iyr != INVALID_PASSPORT_FIELD &&
               eyr != INVALID_PASSPORT_FIELD && pid.length() != 0 &&
               hgt.length() != 0 && hcl.length() != 0 && ecl.length() != 0;
    }

    bool isValidStrict() {
        if (!isValid()) {
            return false;
        }

        // byr (Birth Year) - four digits; at least 1920 and at most 2002.
        if (byr < 1920 || byr > 2002) {
            return false;
        }

        // iyr (Issue Year) - four digits; at least 2010 and at most 2020.
        if (iyr < 2010 || iyr > 2020) {
            return false;
        }

        // eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
        if (eyr < 2020 || eyr > 2030) {
            return false;
        }

        // hgt (Height) - a number followed by either cm or in:
        // If cm, the number must be at least 150 and at most 193.
        // If in, the number must be at least 59 and at most 76.
        if (hgt.length() < 3) {
            return false;
        }
        std::string unit = hgt.substr(hgt.length() - 2);
        if (unit == "cm") {
            int x = std::stoi(hgt.substr(0, 3));
            if (x < 150 || x > 193) {
                return false;
            }
        } else if (unit == "in") {
            int x = std::stoi(hgt.substr(0, 2));
            if (x < 59 || x > 76) {
                return false;
            }
        } else {
            return false;
        }

        // hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
        if (hcl[0] != '#') {
            return false;
        }
        if (hcl.length() != 7) {
            return false;
        }
        for (int i = 1; i < 7; i++) {
            if (hcl[i] >= '0' && hcl[i] <= '9') {
                continue;
            } else if (hcl[i] >= 'a' && hcl[i] <= 'f') {
                continue;
            } else {
                return false;
            }
        }

        // ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
        if (ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" &&
            ecl != "grn" && ecl != "hzl" && ecl != "oth") {
            return false;
        }

        // pid (Passport ID) - a nine-digit number, including leading zeroes.
        if (pid.length() != 9) {
            return false;
        }

        // cid (Country ID) - ignored, missing or not.
        return true;
    }

    std::string string() {
        std::stringstream ss;
        ss << byr << ' ' << iyr << ' ' << eyr << ' ' << hgt << ' ' << hcl << ' '
           << ecl << ' ' << pid << ' ' << cid;
        return ss.str();
    }

  private:
    int byr;         // (Birth Year)
    int iyr;         // (Issue Year)
    int eyr;         // (Expiration Year)
    std::string hgt; // (Height)
    std::string hcl; // (Hair Color)
    std::string ecl; // (Eye Color)
    std::string pid; // (Passport ID)
    long int cid;    // (Country ID)
};

void Solution::part1() {
    std::ifstream file(input_path);
    std::string line;
    std::unique_ptr<Passport> passport = std::make_unique<Passport>();
    int n_valid = 0;

    while (std::getline(file, line)) {
        if (line == "") {
            if (passport.get()->isValid()) {
                n_valid++;
            }
            passport.reset(new Passport());
        } else {
            passport.get()->parseString(line);
        }
    }

    if (passport.get()->isValid()) {
        n_valid++;
    }

    std::cout << "Part 1: " << n_valid << std::endl;
}

void Solution::part2() {
    std::ifstream file(input_path);
    std::string line;
    std::unique_ptr<Passport> passport = std::make_unique<Passport>();
    int n_valid = 0;

    while (std::getline(file, line)) {
        if (line == "") {
            if (passport.get()->isValidStrict()) {
                n_valid++;
            }
            passport.reset(new Passport());
        } else {
            passport.get()->parseString(line);
        }
    }

    if (passport.get()->isValidStrict()) {
        n_valid++;
    }

    std::cout << "Part 2: " << n_valid << std::endl;
}
