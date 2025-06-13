package main

import (
	"github.com/jcoelho93/gotop/internal/uptime"
)

func main() {
	uptimeReader := uptime.NewUptimeReader("/proc/uptime")
	uptime, err := uptimeReader.Read()
	if err != nil {
		panic(err)
	}
	println("Uptime in seconds:", uptime.Seconds)
	println("Idle time in seconds:", uptime.Idle)
}
