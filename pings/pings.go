package pings

import (
	"github.com/johnswanson/ttc"
	"math"
)

const IA = int64(16807)
const IM = int64(2147483647)
const FIRSTPING = int64(1184083200)

func nextRan(v int64) int64 {
	return (IA * v) % IM
}

func PingChannel(cfg ttc.Config) chan int64 {
	c := make(chan int64)
	go func() {
		ping := FIRSTPING
		c <- FIRSTPING
		for ran0 := nextRan(cfg.Seed); ; ran0 = nextRan(ran0) {
			ran1 := float64(ran0) / float64(IM)
			exprand := -1 * float64(cfg.Gap) * math.Log(ran1)
			nextPing := ping + int64(exprand+0.5)
			if ping+1 > nextPing {
				ping += 1
			} else {
				ping = nextPing
			}
			c <- ping
		}
	}()
	return c
}
