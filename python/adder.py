"""Add up some numbers in buckets."""

import time
d = {}


def adder(num: int, num_buckets: int = 10):
    for i in range(num):
        for b in range(num_buckets):
            d.setdefault(b, 0)
            d[b] += 1


if __name__ == '__main__':
    n = 100000

    start = time.perf_counter()
    adder(n)
    end = time.perf_counter()

    for k, v in d.items():
        print(f'd[{k}] = {v}')

    secs = end - start
    print(f'Wall-clock time: {secs:0.2f} secs')
