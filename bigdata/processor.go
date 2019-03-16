package bigdata

import (
	"time"
)

func DoComputation() bool {
	time.Sleep(5 * time.Second)
	return true
}

func DoComplexComputation() bool {
	time.Sleep(10 * time.Second)
	return true
}
