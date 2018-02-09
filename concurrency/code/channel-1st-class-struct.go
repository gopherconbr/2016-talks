type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

