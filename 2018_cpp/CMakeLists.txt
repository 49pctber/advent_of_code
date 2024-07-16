cmake_minimum_required(VERSION 3.22)
project(aoc2018 VERSION 0.0.0)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set(INPUT_DIR "${CMAKE_SOURCE_DIR}/input")
set(LIB_SOURCES src/day1.cpp)

add_library(${PROJECT_NAME}_lib STATIC ${LIB_SOURCES})
target_include_directories(${PROJECT_NAME}_lib PUBLIC include)
target_compile_definitions(${PROJECT_NAME}_lib PRIVATE
    INPUT_DIR="${INPUT_DIR}"
)

add_executable(
    ${PROJECT_NAME}
    src/main.cpp
)
target_link_libraries(${PROJECT_NAME} PRIVATE ${PROJECT_NAME}_lib)

add_subdirectory(extern/googletest)
enable_testing()
add_subdirectory(tests)
