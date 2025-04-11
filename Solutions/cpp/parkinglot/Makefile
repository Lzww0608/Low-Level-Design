CXX = g++
CXXFLAGS = -std=c++11 -Wall -Wextra
TARGET = parking_lot_demo
SRCS = demo.cc ParkingLot.cc ParkingSpot.cc Vehicle.cc
OBJS = $(SRCS:.cc=.o)

# 默认目标
all: $(TARGET)

# 编译目标文件
%.o: %.cc
	$(CXX) $(CXXFLAGS) -c $< -o $@

# 链接目标文件生成可执行文件
$(TARGET): $(OBJS)
	$(CXX) $(CXXFLAGS) -o $@ $(OBJS)

# 清理生成的文件
clean:
	rm -f $(OBJS) $(TARGET)

# 运行程序
run: $(TARGET)
	./$(TARGET)

.PHONY: all clean run 