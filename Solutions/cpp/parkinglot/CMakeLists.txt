cmake_minimum_required(VERSION 3.10)
project(ParkingLotSystem)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# 添加所有源文件
set(SOURCES
    demo.cc
    ParkingLot.cc
    ParkingSpot.cc
    Vehicle.cc
)

# 添加可执行文件
add_executable(parking_lot_demo ${SOURCES})

# 添加头文件目录
target_include_directories(parking_lot_demo PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}) 