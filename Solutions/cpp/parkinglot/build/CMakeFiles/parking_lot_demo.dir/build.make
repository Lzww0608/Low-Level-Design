# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.28

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:

#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:

# Disable VCS-based implicit rules.
% : %,v

# Disable VCS-based implicit rules.
% : RCS/%

# Disable VCS-based implicit rules.
% : RCS/%,v

# Disable VCS-based implicit rules.
% : SCCS/s.%

# Disable VCS-based implicit rules.
% : s.%

.SUFFIXES: .hpux_make_needs_suffix_list

# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:
.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/build

# Include any dependencies generated for this target.
include CMakeFiles/parking_lot_demo.dir/depend.make
# Include any dependencies generated by the compiler for this target.
include CMakeFiles/parking_lot_demo.dir/compiler_depend.make

# Include the progress variables for this target.
include CMakeFiles/parking_lot_demo.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/parking_lot_demo.dir/flags.make

CMakeFiles/parking_lot_demo.dir/demo.cc.o: CMakeFiles/parking_lot_demo.dir/flags.make
CMakeFiles/parking_lot_demo.dir/demo.cc.o: /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/demo.cc
CMakeFiles/parking_lot_demo.dir/demo.cc.o: CMakeFiles/parking_lot_demo.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/parking_lot_demo.dir/demo.cc.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/parking_lot_demo.dir/demo.cc.o -MF CMakeFiles/parking_lot_demo.dir/demo.cc.o.d -o CMakeFiles/parking_lot_demo.dir/demo.cc.o -c /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/demo.cc

CMakeFiles/parking_lot_demo.dir/demo.cc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing CXX source to CMakeFiles/parking_lot_demo.dir/demo.cc.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/demo.cc > CMakeFiles/parking_lot_demo.dir/demo.cc.i

CMakeFiles/parking_lot_demo.dir/demo.cc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling CXX source to assembly CMakeFiles/parking_lot_demo.dir/demo.cc.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/demo.cc -o CMakeFiles/parking_lot_demo.dir/demo.cc.s

CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.o: CMakeFiles/parking_lot_demo.dir/flags.make
CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.o: /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/ParkingLot.cc
CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.o: CMakeFiles/parking_lot_demo.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building CXX object CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.o -MF CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.o.d -o CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.o -c /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/ParkingLot.cc

CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing CXX source to CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/ParkingLot.cc > CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.i

CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling CXX source to assembly CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/ParkingLot.cc -o CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.s

CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.o: CMakeFiles/parking_lot_demo.dir/flags.make
CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.o: /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/ParkingSpot.cc
CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.o: CMakeFiles/parking_lot_demo.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Building CXX object CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.o -MF CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.o.d -o CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.o -c /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/ParkingSpot.cc

CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing CXX source to CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/ParkingSpot.cc > CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.i

CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling CXX source to assembly CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/ParkingSpot.cc -o CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.s

CMakeFiles/parking_lot_demo.dir/Vehicle.cc.o: CMakeFiles/parking_lot_demo.dir/flags.make
CMakeFiles/parking_lot_demo.dir/Vehicle.cc.o: /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/Vehicle.cc
CMakeFiles/parking_lot_demo.dir/Vehicle.cc.o: CMakeFiles/parking_lot_demo.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_4) "Building CXX object CMakeFiles/parking_lot_demo.dir/Vehicle.cc.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/parking_lot_demo.dir/Vehicle.cc.o -MF CMakeFiles/parking_lot_demo.dir/Vehicle.cc.o.d -o CMakeFiles/parking_lot_demo.dir/Vehicle.cc.o -c /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/Vehicle.cc

CMakeFiles/parking_lot_demo.dir/Vehicle.cc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing CXX source to CMakeFiles/parking_lot_demo.dir/Vehicle.cc.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/Vehicle.cc > CMakeFiles/parking_lot_demo.dir/Vehicle.cc.i

CMakeFiles/parking_lot_demo.dir/Vehicle.cc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling CXX source to assembly CMakeFiles/parking_lot_demo.dir/Vehicle.cc.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/Vehicle.cc -o CMakeFiles/parking_lot_demo.dir/Vehicle.cc.s

# Object files for target parking_lot_demo
parking_lot_demo_OBJECTS = \
"CMakeFiles/parking_lot_demo.dir/demo.cc.o" \
"CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.o" \
"CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.o" \
"CMakeFiles/parking_lot_demo.dir/Vehicle.cc.o"

# External object files for target parking_lot_demo
parking_lot_demo_EXTERNAL_OBJECTS =

parking_lot_demo: CMakeFiles/parking_lot_demo.dir/demo.cc.o
parking_lot_demo: CMakeFiles/parking_lot_demo.dir/ParkingLot.cc.o
parking_lot_demo: CMakeFiles/parking_lot_demo.dir/ParkingSpot.cc.o
parking_lot_demo: CMakeFiles/parking_lot_demo.dir/Vehicle.cc.o
parking_lot_demo: CMakeFiles/parking_lot_demo.dir/build.make
parking_lot_demo: CMakeFiles/parking_lot_demo.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --bold --progress-dir=/home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_5) "Linking CXX executable parking_lot_demo"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/parking_lot_demo.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/parking_lot_demo.dir/build: parking_lot_demo
.PHONY : CMakeFiles/parking_lot_demo.dir/build

CMakeFiles/parking_lot_demo.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/parking_lot_demo.dir/cmake_clean.cmake
.PHONY : CMakeFiles/parking_lot_demo.dir/clean

CMakeFiles/parking_lot_demo.dir/depend:
	cd /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/build && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/build /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/build /home/haider/Work/Lzww/practice/Low-Level-Design/Solutions/cpp/parkinglot/build/CMakeFiles/parking_lot_demo.dir/DependInfo.cmake "--color=$(COLOR)"
.PHONY : CMakeFiles/parking_lot_demo.dir/depend

