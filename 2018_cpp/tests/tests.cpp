#include "aoc.hpp"
#include <gtest/gtest.h>

static std::filesystem::path dir(TEST_INPUT_DIR);

TEST(ExampleTest, 1_1_0) { EXPECT_EQ(day1part1(dir / "1_1_0.txt"), 3); }

TEST(ExampleTest, 1_1_1) { EXPECT_EQ(day1part1(dir / "1_1_1.txt"), 3); }

TEST(ExampleTest, 1_1_2) { EXPECT_EQ(day1part1(dir / "1_1_2.txt"), 0); }

TEST(ExampleTest, 1_1_3) { EXPECT_EQ(day1part1(dir / "1_1_3.txt"), -6); }

TEST(ExampleTest, 1_2_0) { EXPECT_EQ(day1part2(dir / "1_2_0.txt"), 2); }

TEST(ExampleTest, 1_2_1) { EXPECT_EQ(day1part2(dir / "1_2_1.txt"), 0); }

TEST(ExampleTest, 1_2_2) { EXPECT_EQ(day1part2(dir / "1_2_2.txt"), 10); }

TEST(ExampleTest, 1_2_3) { EXPECT_EQ(day1part2(dir / "1_2_3.txt"), 5); }

TEST(ExampleTest, 1_2_4) { EXPECT_EQ(day1part2(dir / "1_2_4.txt"), 14); }
