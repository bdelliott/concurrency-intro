package adder

import "sync"

func ThreadedAddr(counter *int, num int) {
	var wg sync.WaitGroup
	num = num / numWorkers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(counter *int, num int, wg *sync.WaitGroup) {
			Adder(counter, num)
			wg.Done()
		}(counter, num, &wg)
	}
	wg.Wait()
}