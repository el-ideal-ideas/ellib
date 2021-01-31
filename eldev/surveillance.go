package eldev

import (
	"runtime"
	"time"
)

type Stats struct {
	Time         int64
	TimeForHuman string
	NumGoroutine int
	MemStats     *runtime.MemStats
}

// Create a Stats object.
func CurrentStats() *Stats {
	mem := new(runtime.MemStats)
	runtime.ReadMemStats(mem)
	now := time.Now()
	return &Stats{
		Time:         now.UnixNano(),
		TimeForHuman: now.Format(time.RFC3339Nano),
		NumGoroutine: runtime.NumGoroutine(),
		MemStats:     mem,
	}
}

// Update stats.
func (s *Stats) Update() {
	now := time.Now()
	s.Time = now.UnixNano()
	s.TimeForHuman = now.Format(time.RFC3339Nano)
	s.NumGoroutine = runtime.NumGoroutine()
	runtime.ReadMemStats(s.MemStats)
}

// Start a goroutine and check system stats with interval.
func Start(d time.Duration, f func(*Stats)) {
	go func() {
		s := CurrentStats()
		timer := time.NewTicker(d)
		defer timer.Stop()
		for range timer.C {
			s.Update()
			f(s)
		}
	}()
}
