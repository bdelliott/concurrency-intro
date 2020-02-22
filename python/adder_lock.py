"""Demonstrate thread safety in critical sections."""

import threading
import time

d = {}
lock = threading.RLock()


def adder(num: int, num_buckets: int = 10):
	for i in range(num):
		for b in range(num_buckets):
			with lock:  # one at a time
				d.setdefault(b, 0)
				d[b] += 1


def threaded_adder(num: int, num_workers: int = 8, num_buckets: int = 10):
	"""Break python task up among threads."""
	num = int(num / num_workers)

	threads = []
	for i in range(num_workers):
		thread = threading.Thread(target=adder, args=(num,), kwargs={'num_buckets': num_buckets})
		threads.append(thread)

	for thread in threads:
		thread.start()

	for thread in threads:
		thread.join()


if __name__ == '__main__':
	n = 100000

	start = time.perf_counter()
	threaded_adder(n)
	end = time.perf_counter()

	for k, v in d.items():
		print(f'd[{k}] = {v}')

	secs = end - start
	print(f'Wall-clock time: {secs:0.2f} secs')
