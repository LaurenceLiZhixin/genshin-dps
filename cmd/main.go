package main

import (
	"github.com/alibaba/ioc-golang"

	"github.com/laurencelizhixin/genshin-dps/internal/simulator"
)

func main() {
	if err := ioc.Load(); err != nil {
		panic(err)
	}
	(&simulator.DamageSimulator{}).Run()
}
