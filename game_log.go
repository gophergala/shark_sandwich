package main

import (
//"encoding/json"
)

type LogEvent struct {
	Message  string
	Xp       int
	Life     int
	Speed    int
	Power    int
	Ancestry int
}

type GameLog struct {
	recvEvents <-chan LogEvent
}

func (s *GameLog) InitLogEventStream(logEvents <-chan LogEvent) {
	s.recvEvents = logEvents

	go func() {
		for logEvent := range logEvents {
			s.storeLogEvent(logEvent)
		}
	}()
}

func NewGameLog() (*GameLog, error) {
	gameLog := &GameLog{
		recvEvents: make(chan LogEvent),
	}

	return gameLog, nil
}

func (s *GameLog) storeLogEvent(logEvent LogEvent) {

}

func printLog() {
	// todo: loop and print from memory structure
}
