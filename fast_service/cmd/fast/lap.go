package main

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
)

type Lap struct {
	Bytes int64
	Bps  float64
	PrettyBps float64
	Units string
	
	delta float64
}

func newLap(bytelen int64, delta float64) Lap {
	var bytes float64
	if delta > 0 {
		bytes = float64(bytelen) / delta
	} 
	bps := bytes * 8.0
	prettyBps, unit := humanize.ComputeSI(bps)
	return Lap{
		Bytes:     bytelen,
		Bps:       bps,
		PrettyBps: prettyBps,
		Units:     unit,
		delta:     delta,
	}
}

func (l *Lap) String() string {
	return fmt.Sprintf("%7.2f %s/s", l.PrettyBps, l.Units)
}