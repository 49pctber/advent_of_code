cmake_minimum_required(VERSION 3.25)

project(AoC)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")

include_directories(include)

add_executable(day1 src/main.cpp src/day1.cpp)
add_executable(day2 src/main.cpp src/day2.cpp)
add_executable(day3 src/main.cpp src/day3.cpp)
add_executable(day4 src/main.cpp src/day4.cpp)
add_executable(day5 src/main.cpp src/day5.cpp)
add_executable(day6 src/main.cpp src/day6.cpp)
add_executable(day7 src/main.cpp src/day7.cpp)
add_executable(day8 src/main.cpp src/day8.cpp)
add_executable(day9 src/main.cpp src/day9.cpp)
add_executable(day10 src/main.cpp src/day10.cpp)
add_executable(day11 src/main.cpp src/day11.cpp)
