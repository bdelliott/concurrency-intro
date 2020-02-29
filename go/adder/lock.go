package adder

import "sync"

func ThreadedLockingAddr(counter *int, num int) {

	var mutex sync.Mutex

	var wg sync.WaitGroup
	num = num / numWorkers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(counter *int, num int, wg *sync.WaitGroup) {
			mutex.Lock()
			Adder(counter, num)
			mutex.Unlock()
			wg.Done()
		}(counter, num, &wg)
	}
	wg.Wait()
}