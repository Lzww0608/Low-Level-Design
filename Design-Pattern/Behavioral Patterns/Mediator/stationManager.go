package main

type StationManager struct {
	isPlatformFree bool 
	trainQueue []Train
}


func newStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
		trainQueue: []Train{},
	}
}

func (s *StationManager) canArrive(train Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, train)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true	
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}