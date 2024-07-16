add_executable(
  unit_tests
  tests.cpp
)

set(TEST_INPUT_DIR "${CMAKE_SOURCE_DIR}/tests/input")
target_compile_definitions(unit_tests PRIVATE
  TEST_INPUT_DIR="${TEST_INPUT_DIR}"
)

target_link_libraries(
  unit_tests
  PRIVATE
  ${PROJECT_NAME}_lib
  gtest_main
)

include(GoogleTest)
gtest_discover_tests(unit_tests)
