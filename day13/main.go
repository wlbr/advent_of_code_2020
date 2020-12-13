package main

import (
	"fmt"
	"runtime"
	"sync"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func MinBus(v ...int) (min, index int) {
	m := v[0]
	n := -1
	for i, e := range v {
		if e < m {
			m = e
			n = i
		}
	}
	return m, n
}

func calcNextTimeForBus(s, b int) int {
	d := s / b
	n := b*d + b
	//log.Printf("s: %d b: %d ? d: %d  r: %d", s, b, d, n)
	return n
}

func getNextBus(start int, busses []int) (bus, departure, result int) {
	var times []int
	for _, l := range busses {
		times = append(times, calcNextTimeForBus(start, l))
	}
	dep, bus := MinBus(times...)
	return busses[bus], dep, (dep - start) * busses[bus]
}

func busRidesAtTimeStamp(start, ts, bus int) bool {
	if bus == -1 {
		return true
	}
	d := ts % bus
	if d == 0 {
		return true
	} else {
		return false
	}
}

func isValidSequence(ts int, busses []int) bool {
	r := true
	//log.Print("ts: ", ts, "  busses ", busses)
	for i, b := range busses {
		if b != -1 {
			if 0 != ((ts + i) % b) {
				return false
			}
		}
	}
	return r
}

func searchChunk(start, end int, busses []int) (r int) {
	for i := start; i < end; i++ {
		v := isValidSequence(i, busses)
		if v {
			r = i
			break
		}
	}
	return r
}

type worker struct {
	id           int
	inconmigJobs chan *Job
	results      chan int // chan<- int
	wg           *sync.WaitGroup
}

func (w *worker) run() {
	defer w.wg.Done()
	for j := range w.inconmigJobs {
		if j.start == -1 { // catch kill signal
			break
		} else {
			r := searchChunk(j.start, j.end, j.busses)
			//fmt.Printf("Worker: %d  Set: [%d-%d]  r: %d\n", w.id, j.start, j.end, r)
			if r != 0 {
				w.results <- r
			}
		}
	}
}

func stopWorkers(ws []*worker) {
	for _, w := range ws {
		w.inconmigJobs <- &Job{-1, 0, []int{}}
	}
}

type Job struct {
	start, end int
	busses     []int
}

func findBusSequenceMT(busses []int) (result int) {
	var wg sync.WaitGroup
	numJobs := runtime.NumCPU()
	var workers []*worker
	in := make(chan *Job)
	res := make(chan int)

	for i := 0; i < numJobs; i++ {
		w := &worker{i, in, res, &wg}
		workers = append(workers, w)
		wg.Add(1)
		go w.run()
	}

	//catch results
	go func(res chan int) {
		for v := range res {
			if (v != 0 && result == 0) || v < result {
				result = v
			}
		}
	}(res)

	//      1068781
	//      1000001
	step := 10000000000
	for j := 1; j <= MaxInt; j += step {
		jj := &Job{j, j + step, busses}
		if result != 0 {
			stopWorkers(workers)
			break
		}
		in <- jj
	}
	wg.Wait()
	close(in)
	close(res)

	return result
}

func findBusSequenceST(busses []int) (r int) {
	for i := 0; i < MaxInt; i++ {
		v := isValidSequence(i, busses)
		if 0 == i%1000000000 {
			fmt.Printf("i: %d \n", i)
		}

		if v {
			r = i
			return i
		}
	}
	return r
}

func main() {
	input := "input.txt"
	//input = "example.txt"

	s, all, act := readdata(input)
	b, d, r := getNextBus(s, act)
	fmt.Printf("Departure: %d  Bus: %d  Solution: %d\n", d, b, r)

	//fmt.Println("ST ", findBusSequenceST(all))
	fmt.Println("MT ", findBusSequenceMT(all))

}
