package adder

func ChanneledAdder(counter *int, num int) {

	num = num / numWorkers
	channel := make(chan int)

	for i := 0; i < numWorkers; i++ {
		go worker(num, channel)
	}

	for i := 0; i < numWorkers; i++ {
		*counter += <-channel // increment counter by the value of a message received on the channel
	}
}

func worker(num int, channel chan int) {
	counter := 0
	Adder(&counter, num)
	channel <- num
}
