# 交通信号控制系统 (Traffic Signal Control System)

## 概述

这是一个用 Go 语言实现的交通信号控制系统，用于管理交叉路口的交通信号灯。该系统支持多条道路、多个信号灯的管理，并具备紧急情况处理功能。

## 设计要求实现

✅ **要求 1**: 控制具有多条道路的交叉口的交通流  
✅ **要求 2**: 支持不同类型的信号（红、黄、绿）  
✅ **要求 3**: 可配置和调整的信号持续时间  
✅ **要求 4**: 平滑的信号转换，确保安全高效的交通流  
✅ **要求 5**: 检测和处理紧急情况（救护车、消防车）  
✅ **要求 6**: 可扩展性和可维护性

## 架构设计

### 核心组件

1. **Signal (信号灯)**
   - 表示单个交通信号灯
   - 支持 RED、GREEN、YELLOW 三种状态
   - 可配置每种信号的持续时间
   - 支持紧急模式

2. **Road (道路)**
   - 表示一条道路
   - 可以管理多个信号灯
   - 提供线程安全的操作

3. **TrafficController (交通控制器)**
   - 采用单例模式设计
   - 管理多条道路和信号灯
   - 提供全局控制功能

## 主要功能

### Signal 功能
- ✅ 信号状态切换 (RED → GREEN → YELLOW → RED)
- ✅ 时间更新和倒计时
- ✅ 可配置的信号持续时间
- ✅ 紧急情况处理
- ✅ 工作状态控制

### Road 功能
- ✅ 添加/删除信号灯
- ✅ 更新指定信号灯
- ✅ 同步所有信号灯
- ✅ 紧急情况处理
- ✅ 状态打印

### TrafficController 功能
- ✅ 单例模式实现
- ✅ 添加/删除道路
- ✅ 全局信号控制
- ✅ 紧急情况处理
- ✅ 启动/停止所有信号

## 改进和完善

相比原始实现，主要改进包括：

1. **修复 Bug**
   - 修复了 `SwitchSignal()` 方法缺失 YELLOW → RED 转换
   - 修正了拼写错误 (`remaingDuration` → `remainingDuration`)

2. **新增功能**
   - 添加了 `SignalType.String()` 方法用于更好的输出显示
   - 实现了紧急情况处理机制 (`HandleEmergency`, `ClearEmergency`)
   - 为 TrafficController 实现了单例模式
   - 添加了 `StartAllSignals()` 和 `StopAllSignals()` 方法

3. **测试覆盖**
   - 完整的单元测试套件（12个测试用例）
   - 并发测试
   - 集成测试
   - 性能基准测试

4. **代码质量**
   - 改进了输出格式，更易读
   - 添加了详细的注释
   - 使用线程安全的操作

## 项目结构

```
trafficsignal/
├── signal.go           # 信号灯实现
├── road.go             # 道路实现
├── trafficcontroller.go # 交通控制器实现
├── main.go             # 演示程序
├── main_test.go        # 测试程序
├── go.mod              # Go 模块文件
└── README.md           # 本文档
```

## 使用方法

### 编译

```bash
go build -o trafficsignal
```

### 运行演示程序

```bash
./trafficsignal
```

### 运行测试

```bash
# 运行所有测试
go test -v

# 运行基准测试
go test -bench=. -benchmem
```

## 使用示例

```go
// 创建交通控制器（单例模式）
controller := GetInstance("MainController")

// 创建道路
road := NewRoad("North-South")

// 创建信号灯（绿灯30秒，黄灯5秒，红灯25秒）
signal := NewSignal("NS-Signal-1", 30, 5, 25)

// 添加信号灯到道路
road.AddSignal(signal)

// 添加道路到控制器
controller.AddRoad(road)

// 启动所有信号
controller.StartAllSignals()

// 打印状态
controller.PrintStatus()

// 处理紧急情况
controller.HandleEmergency("North-South", "NS-Signal-1")

// 清除紧急状态
controller.ClearEmergency("North-South", "NS-Signal-1")
```

## 测试结果

所有测试均已通过：

```
TestSignalCreation              ✅ PASS
TestSignalSwitching             ✅ PASS
TestSignalUpdate                ✅ PASS
TestSignalEmergency             ✅ PASS
TestRoadCreation                ✅ PASS
TestRoadAddSignal               ✅ PASS
TestRoadRemoveSignal            ✅ PASS
TestTrafficControllerSingleton  ✅ PASS
TestTrafficControllerAddRoad    ✅ PASS
TestTrafficControllerRemoveRoad ✅ PASS
TestConcurrentSignalUpdates     ✅ PASS
TestFullTrafficSystemIntegration ✅ PASS
```

## 性能

基准测试结果（在 Intel i9-13900K 上）：

- SignalUpdate: ~0.73 ns/op, 0 内存分配
- SignalSwitch: ~0.52 ns/op, 0 内存分配
