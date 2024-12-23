cmake_minimum_required(VERSION 3.25)

project(AoC)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

set(CMAKE_CXX_FLAGS "-Wall")

include_directories(include)

foreach(year RANGE 2015 2024)
    foreach(day RANGE 1 25)
        set(DAY_FILE src/${year}/day${day}.cpp)

        if(EXISTS ../${DAY_FILE})
            add_executable(${year}_${day} src/main.cpp ${DAY_FILE})
        endif()
    endforeach()
endforeach()
