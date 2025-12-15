package main

import (
	"testing"
)

func TestSignalCreation(t *testing.T) {
	signal := NewSignal("TEST-1", 30, 5, 25)
	if signal.GetSignalId() != "TEST-1" {
		t.Errorf("Expected signal ID to be TEST-1, got %s", signal.GetSignalId())
	}
	if signal.GetCurrentSignal() != RED {
		t.Errorf("Expected initial signal to be RED, got %s", signal.GetCurrentSignal())
	}
	if signal.GetRemainingDuration() != 25 {
		t.Errorf("Expected remaining duration to be 25, got %d", signal.GetRemainingDuration())
	}
}

func TestSignalSwitching(t *testing.T) {
	signal := NewSignal("TEST-2", 30, 5, 25)
	signal.SetWorking(true)

	// Initial state should be RED
	if signal.GetCurrentSignal() != RED {
		t.Errorf("Expected RED, got %s", signal.GetCurrentSignal())
	}

	// Switch to GREEN
	signal.SwitchSignal()
	if signal.GetCurrentSignal() != GREEN {
		t.Errorf("Expected GREEN after switch from RED, got %s", signal.GetCurrentSignal())
	}

	// Switch to YELLOW
	signal.SwitchSignal()
	if signal.GetCurrentSignal() != YELLOW {
		t.Errorf("Expected YELLOW after switch from GREEN, got %s", signal.GetCurrentSignal())
	}

	// Switch back to RED
	signal.SwitchSignal()
	if signal.GetCurrentSignal() != RED {
		t.Errorf("Expected RED after switch from YELLOW, got %s", signal.GetCurrentSignal())
	}
}

func TestSignalUpdate(t *testing.T) {
	signal := NewSignal("TEST-3", 30, 5, 25)
	signal.SetWorking(true)

	initialDuration := signal.GetRemainingDuration()
	signal.UpdateSignal(10)

	if signal.GetRemainingDuration() != initialDuration-10 {
		t.Errorf("Expected remaining duration to be %d, got %d", initialDuration-10, signal.GetRemainingDuration())
	}
}

func TestSignalEmergency(t *testing.T) {
	signal := NewSignal("TEST-4", 30, 5, 25)
	signal.SetWorking(true)

	// Trigger emergency
	signal.HandleEmergency()
	if signal.GetCurrentSignal() != GREEN {
		t.Errorf("Expected GREEN during emergency, got %s", signal.GetCurrentSignal())
	}

	// Clear emergency
	signal.ClearEmergency()
	// Signal should still be GREEN but emergency flag cleared
	if signal.GetCurrentSignal() != GREEN {
		t.Errorf("Expected GREEN after clearing emergency, got %s", signal.GetCurrentSignal())
	}
}

func TestRoadCreation(t *testing.T) {
	road := NewRoad("TEST-ROAD")
	if road.GetRoadId() != "TEST-ROAD" {
		t.Errorf("Expected road ID to be TEST-ROAD, got %s", road.GetRoadId())
	}
	if road.GetSignalCount() != 0 {
		t.Errorf("Expected signal count to be 0, got %d", road.GetSignalCount())
	}
}

func TestRoadAddSignal(t *testing.T) {
	road := NewRoad("TEST-ROAD-2")
	signal := NewSignal("TEST-SIGNAL", 30, 5, 25)

	road.AddSignal(signal)
	if road.GetSignalCount() != 1 {
		t.Errorf("Expected signal count to be 1, got %d", road.GetSignalCount())
	}
}

func TestRoadRemoveSignal(t *testing.T) {
	road := NewRoad("TEST-ROAD-3")
	signal := NewSignal("TEST-SIGNAL-REMOVE", 30, 5, 25)

	road.AddSignal(signal)
	if road.GetSignalCount() != 1 {
		t.Errorf("Expected signal count to be 1, got %d", road.GetSignalCount())
	}

	road.RemoveSignal("TEST-SIGNAL-REMOVE")
	if road.GetSignalCount() != 0 {
		t.Errorf("Expected signal count to be 0 after removal, got %d", road.GetSignalCount())
	}
}

func TestTrafficControllerSingleton(t *testing.T) {
	controller1 := GetInstance("CONTROLLER-1")
	controller2 := GetInstance("CONTROLLER-2")

	if controller1 != controller2 {
		t.Error("Expected singleton pattern to return same instance")
	}
}

func TestTrafficControllerAddRoad(t *testing.T) {
	controller := NewTrafficController("TEST-CONTROLLER")
	road := NewRoad("TEST-ROAD-4")

	controller.AddRoad(road)
	if controller.GetRoadCount() != 1 {
		t.Errorf("Expected road count to be 1, got %d", controller.GetRoadCount())
	}
}

func TestTrafficControllerRemoveRoad(t *testing.T) {
	controller := NewTrafficController("TEST-CONTROLLER-2")
	road := NewRoad("TEST-ROAD-5")

	controller.AddRoad(road)
	controller.RemoveRoad("TEST-ROAD-5")

	if controller.GetRoadCount() != 0 {
		t.Errorf("Expected road count to be 0 after removal, got %d", controller.GetRoadCount())
	}
}

func TestConcurrentSignalUpdates(t *testing.T) {
	signal := NewSignal("CONCURRENT-TEST", 30, 5, 25)
	signal.SetWorking(true)

	done := make(chan bool)
	numGoroutines := 10

	for i := 0; i < numGoroutines; i++ {
		go func() {
			signal.UpdateSignal(1)
			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < numGoroutines; i++ {
		<-done
	}

	// Signal should have been updated
	if signal.GetRemainingDuration() >= 25 {
		t.Errorf("Expected remaining duration to decrease after concurrent updates")
	}
}

func TestFullTrafficSystemIntegration(t *testing.T) {
	// Create controller
	controller := NewTrafficController("INTEGRATION-TEST")

	// Create roads
	road1 := NewRoad("North-South")
	road2 := NewRoad("East-West")

	// Create signals
	signal1 := NewSignal("NS-1", 30, 5, 25)
	signal2 := NewSignal("EW-1", 25, 5, 30)

	// Add signals to roads
	road1.AddSignal(signal1)
	road2.AddSignal(signal2)

	// Add roads to controller
	controller.AddRoad(road1)
	controller.AddRoad(road2)

	// Start signals
	controller.StartAllSignals()

	// Verify all signals are working
	for _, road := range controller.GetRoads() {
		for _, signal := range road.GetSignals() {
			if !signal.GetIsWorking() {
				t.Error("Expected all signals to be working after StartAllSignals")
			}
		}
	}

	// Test emergency handling
	controller.HandleEmergency("North-South", "NS-1")
	if signal1.GetCurrentSignal() != GREEN {
		t.Errorf("Expected signal to be GREEN during emergency, got %s", signal1.GetCurrentSignal())
	}

	// Stop signals
	controller.StopAllSignals()
	for _, road := range controller.GetRoads() {
		for _, signal := range road.GetSignals() {
			if signal.GetIsWorking() {
				t.Error("Expected all signals to be stopped after StopAllSignals")
			}
		}
	}
}

func BenchmarkSignalUpdate(b *testing.B) {
	signal := NewSignal("BENCH-1", 30, 5, 25)
	signal.SetWorking(true)
	for i := 0; i < b.N; i++ {
		signal.UpdateSignal(1)
	}
}

func BenchmarkSignalSwitch(b *testing.B) {
	signal := NewSignal("BENCH-2", 30, 5, 25)
	signal.SetWorking(true)
	for i := 0; i < b.N; i++ {
		signal.SwitchSignal()
	}
}
func TestTrafficSignalSwitchWithGoroutineAndTicker(t *testing.T) {
	road := NewRoad("East-West")
	// 设置持续时间：绿灯5秒，黄灯2秒，红灯4秒
	// 一个完整周期 = 4 + 5 + 2 = 11秒
	greenDuration := 5
	yellowDuration := 2
	redDuration := 4
	signal := NewSignal("EW-1", greenDuration, yellowDuration, redDuration)
	road.AddSignal(signal)
	controller := NewTrafficController("CTRL-TICK")
	controller.AddRoad(road)
	controller.StartAllSignals()

	// 记录信号切换历史
	type SignalState struct {
		signal   SignalType
		duration int
		elapsed  int
	}
	var switchHistory []SignalState

	// 模拟时间流动，逐秒更新
	totalTime := 0
	tickInterval := 1 // 每次代表1秒
	maxTime := 25     // 运行25秒（超过2个完整周期）

	for totalTime < maxTime {
		currentSignal := signal.GetCurrentSignal()
		currentRemaining := signal.GetRemainingDuration()

		// 记录当前状态
		switchHistory = append(switchHistory, SignalState{
			signal:   currentSignal,
			duration: currentRemaining,
			elapsed:  totalTime,
		})

		// 更新信号灯（模拟1秒流逝）
		signal.UpdateSignal(tickInterval)
		totalTime += tickInterval
	}

	// 验证信号切换序列
	if len(switchHistory) == 0 {
		t.Fatal("No signal history recorded")
	}

	// 验证初始状态是 RED
	if switchHistory[0].signal != RED {
		t.Errorf("Expected initial signal to be RED, got %s", switchHistory[0].signal)
	}

	// 查找所有信号切换点
	signalChanges := []SignalState{}
	lastSignal := switchHistory[0].signal
	for _, state := range switchHistory {
		if state.signal != lastSignal {
			signalChanges = append(signalChanges, state)
			lastSignal = state.signal
		}
	}

	// 应该至少有5次信号切换（完成近2个周期：RED->GREEN->YELLOW->RED->GREEN->YELLOW）
	if len(signalChanges) < 5 {
		t.Errorf("Expected at least 5 signal changes, got %d", len(signalChanges))
	}

	// 验证信号切换顺序：RED -> GREEN -> YELLOW -> RED -> GREEN -> YELLOW
	expectedSequence := []SignalType{GREEN, YELLOW, RED, GREEN, YELLOW, RED}
	for i := 0; i < len(signalChanges) && i < len(expectedSequence); i++ {
		if signalChanges[i].signal != expectedSequence[i] {
			t.Errorf("Signal change %d: expected %s, got %s at time %d",
				i, expectedSequence[i], signalChanges[i].signal, signalChanges[i].elapsed)
		}
	}

	// 验证信号持续时间是否正确
	// 第一次切换：RED(4秒) -> GREEN，应该在4秒时发生
	if len(signalChanges) > 0 {
		firstChange := signalChanges[0]
		expectedTime := redDuration
		if firstChange.elapsed != expectedTime {
			t.Errorf("First signal change (RED->GREEN) expected at %ds, got %ds",
				expectedTime, firstChange.elapsed)
		}
		if firstChange.signal != GREEN {
			t.Errorf("First change should be to GREEN, got %s", firstChange.signal)
		}
	}

	// 第二次切换：GREEN(5秒) -> YELLOW，应该在4+5=9秒时发生
	if len(signalChanges) > 1 {
		secondChange := signalChanges[1]
		expectedTime := redDuration + greenDuration
		if secondChange.elapsed != expectedTime {
			t.Errorf("Second signal change (GREEN->YELLOW) expected at %ds, got %ds",
				expectedTime, secondChange.elapsed)
		}
		if secondChange.signal != YELLOW {
			t.Errorf("Second change should be to YELLOW, got %s", secondChange.signal)
		}
	}

	// 第三次切换：YELLOW(2秒) -> RED，应该在4+5+2=11秒时发生
	if len(signalChanges) > 2 {
		thirdChange := signalChanges[2]
		expectedTime := redDuration + greenDuration + yellowDuration
		if thirdChange.elapsed != expectedTime {
			t.Errorf("Third signal change (YELLOW->RED) expected at %ds, got %ds",
				expectedTime, thirdChange.elapsed)
		}
		if thirdChange.signal != RED {
			t.Errorf("Third change should be to RED, got %s", thirdChange.signal)
		}
	}

	// 第四次切换：第二轮，RED -> GREEN，应该在11+4=15秒时发生
	if len(signalChanges) > 3 {
		fourthChange := signalChanges[3]
		expectedTime := (redDuration + greenDuration + yellowDuration) + redDuration
		if fourthChange.elapsed != expectedTime {
			t.Errorf("Fourth signal change (RED->GREEN, cycle 2) expected at %ds, got %ds",
				expectedTime, fourthChange.elapsed)
		}
		if fourthChange.signal != GREEN {
			t.Errorf("Fourth change should be to GREEN, got %s", fourthChange.signal)
		}
	}

	// 第五次切换：GREEN -> YELLOW，应该在15+5=20秒时发生
	if len(signalChanges) > 4 {
		fifthChange := signalChanges[4]
		expectedTime := (redDuration + greenDuration + yellowDuration) + redDuration + greenDuration
		if fifthChange.elapsed != expectedTime {
			t.Errorf("Fifth signal change (GREEN->YELLOW, cycle 2) expected at %ds, got %ds",
				expectedTime, fifthChange.elapsed)
		}
		if fifthChange.signal != YELLOW {
			t.Errorf("Fifth change should be to YELLOW, got %s", fifthChange.signal)
		}
	}

	t.Logf("Signal switch history: %d states recorded, %d signal changes detected",
		len(switchHistory), len(signalChanges))
	t.Logf("Signal durations: RED=%ds, GREEN=%ds, YELLOW=%ds (Cycle=%ds)",
		redDuration, greenDuration, yellowDuration, redDuration+greenDuration+yellowDuration)
	for i, change := range signalChanges {
		t.Logf("  Change %d: -> %s at time %ds", i+1, change.signal, change.elapsed)
	}
}
