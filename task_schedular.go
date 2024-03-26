package dsa

import (
	"fmt"
	"math"
)

func reduceCoolDown(taskCoolDown map[byte]int) {
	for key, val := range taskCoolDown {
		taskCoolDown[key] = val - 1
	}
}

func schedulableTask(taskCoolDown, taskCount map[byte]int) byte {
	minVal := math.MaxInt
	minKey := byte('0')
	for key, val := range taskCoolDown {
		if val <= 0 && taskCount[key] > 0 && val < minVal {
			minKey = key
		}
	}
	return minKey
}

func leastInterval(tasks []byte, n int) int {
	steps := 0
	taskCount := make(map[byte]int)
	taskCoolDown := make(map[byte]int)
	for _, e := range tasks {
		taskCount[e]++
		taskCoolDown[e] = 0
	}
	// fmt.Println(taskCoolDown)
	i := 0
	for i < len(tasks) {
		taskToSchedule := schedulableTask(taskCoolDown, taskCount)
		// fmt.Println(taskToSchedule, i, taskCoolDown)
		reduceCoolDown(taskCoolDown)
		steps++
		if taskToSchedule != byte('0') {
			i++
			taskCoolDown[taskToSchedule] = n
			taskCount[taskToSchedule]--
			fmt.Print(string(taskToSchedule), "->")
			continue
		}
		fmt.Print("idle", "->")
	}
	return steps
}
