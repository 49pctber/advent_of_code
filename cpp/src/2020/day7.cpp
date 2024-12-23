#include "solution.hpp"
#include <fstream>
#include <map>
#include <regex>
#include <set>

typedef std::string bag_t;

class Rules {
  public:
    Rules(std::filesystem::path input) {
        std::ifstream file(input);
        std::string line;

        std::regex pattern(R"((\w+ \w+) bags? contain)");
        std::regex pattern2(R"((\d+) (\w+ \w+))");
        std::regex pattern3(R"(no other bags\.)");

        while (std::getline(file, line)) {

            std::smatch matches;
            if (std::regex_search(line, matches, pattern)) {
                bag_t key = matches[1];

                if (std::regex_search(line, matches, pattern3)) {
                    continue;
                }

                auto matches_begin =
                    std::sregex_iterator(line.begin(), line.end(), pattern2);
                auto matches_end = std::sregex_iterator();

                for (std::sregex_iterator i = matches_begin; i != matches_end;
                     ++i) {
                    std::smatch match = *i;
                    bag_t vk = match[2];
                    int vv = std::stoi(match[1]);
                    rules[key][vk] = vv;
                }
            }
        }
    }

    int countBagsWithShinyGold() {
        int count = 0;
        for (auto rule : rules) {
            if (containsShinyGold(rule.first)) {
                count++;
            }
        }
        return count;
    }

    int countBagsInside(bag_t bag) {
        int count = 0;
        for (auto [next, c] : rules[bag]) {
            count += c * (1 + countBagsInside(next));
        }
        return count;
    }

  private:
    std::map<bag_t, std::map<bag_t, int>> rules;
    std::map<bag_t, bool> contains_shiny_gold;

    bool containsShinyGold(bag_t bag) {

        for (const auto &[new_bag, count] : rules[bag]) {

            if (new_bag == "shiny gold") {
                contains_shiny_gold[bag] = true;
                return true;
            }

            // check cache
            if (contains_shiny_gold.find(new_bag) !=
                contains_shiny_gold.end()) {
                if (contains_shiny_gold[new_bag]) {
                    contains_shiny_gold[bag] = true;
                    return true;
                }
            }

            // look in the other bags
            if (containsShinyGold(new_bag)) {
                contains_shiny_gold[bag] = true;
                return true;
            }
        }

        contains_shiny_gold[bag] = false;
        return false;
    }
};

void Solution::part1() {
    Rules rules(argv[1]);
    std::cout << "Part 1: " << rules.countBagsWithShinyGold() << std::endl;
}

void Solution::part2() {
    Rules rules(argv[1]);
    std::cout << "Part 2: " << rules.countBagsInside("shiny gold") << std::endl;
}
