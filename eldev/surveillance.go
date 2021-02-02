package eldev

import (
	"runtime"
	"time"
)

type Surveillance struct {
	Time         int64
	TimeForHuman string
	NumGoroutine int
	MemStats     *runtime.MemStats
}

// Create a Stats object.
func createSurveillance() *Surveillance {
	mem := new(runtime.MemStats)
	runtime.ReadMemStats(mem)
	now := time.Now()
	return &Surveillance{
		Time:         now.UnixNano(),
		TimeForHuman: now.Format(time.RFC3339Nano),
		NumGoroutine: runtime.NumGoroutine(),
		MemStats:     mem,
	}
}

// Update stats.
func (s *Surveillance) updateSurveillance() {
	now := time.Now()
	s.Time = now.UnixNano()
	s.TimeForHuman = now.Format(time.RFC3339Nano)
	s.NumGoroutine = runtime.NumGoroutine()
	runtime.ReadMemStats(s.MemStats)
}

// Start a goroutine and check system stats with an interval.
func RunSurveillance(d time.Duration, f func(*Surveillance)) {
	go func() {
		s := createSurveillance()
		timer := time.NewTicker(d)
		defer timer.Stop()
		for range timer.C {
			s.updateSurveillance()
			f(s)
		}
	}()
}
