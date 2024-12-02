cmake_minimum_required(VERSION 3.25)

project(AoC)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")

include_directories(include)

add_executable(day1 src/main.cpp src/day1.cpp)
add_executable(day2 src/main.cpp src/day2.cpp)
