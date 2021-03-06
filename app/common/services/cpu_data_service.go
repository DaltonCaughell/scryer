package services

import (
	"fmt"
	"math/rand"
	"time"
)

type CPUDataService struct {
}

func (cpuDataService *CPUDataService) GetCPUDataWeekly() [604800]map[time.Time]int {
	var CPUTimeSeries [604800]map[time.Time]int
	currentCPUUsage := 0

	now := time.Now()
	lastWeek := now.AddDate(0, 0, -7)
	timeIteration := lastWeek

	for i := 0; i < 604800; i++ {
		timeIteration = timeIteration.Add(time.Second)
		snapShot := make(map[time.Time]int)
		if rand.Intn(1000) == 5 {
			currentCPUUsage = rand.Intn(101)
		}
		snapShot[timeIteration] = currentCPUUsage
		CPUTimeSeries[i] = snapShot
		fmt.Println(currentCPUUsage)
	}
	return CPUTimeSeries
}
