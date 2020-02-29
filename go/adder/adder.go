package adder

const numWorkers = 8

func Adder(counter *int, num int) {
	for i := 0; i < num; i++ {
		*counter = *counter + 1
	}
}
