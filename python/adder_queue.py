"""Lock-less version of a safe threaded adder."""

import queue
import threading
import time

result_queue = queue.Queue()


def adder(num: int, num_buckets: int = 10):
    d = {}

    for i in range(num):
        for b in range(num_buckets):
            d.setdefault(b, 0)
            d[b] += 1

    result_queue.put(d)


def threaded_adder(num: int, num_workers: int = 8, num_buckets: int = 10) -> dict:
    """Farm adder work up as work units over the queue."""
    d = {}  # local!
    num = int(num / num_workers)

    threads = []
    for i in range(num_workers):
        thread = threading.Thread(target=adder, args=(num,), kwargs={'num_buckets': num_buckets})
        threads.append(thread)

    for thread in threads:
        thread.start()

    for i in range(num_workers):
        e = result_queue.get()

        for k in e:
            d.setdefault(k, 0)
            d[k] += e[k]

    return d


if __name__ == '__main__':
    n = 100000

    start = time.perf_counter()
    d = threaded_adder(n)
    end = time.perf_counter()

    for k, v in d.items():
        print(f'd[{k}] = {v}')

    secs = end - start
    print(f'Wall-clock time: {secs:0.2f} secs')
