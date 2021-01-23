from random import randint
from ellib import random_int
from time import time

s = time()
_ = [randint(0, 100) for _ in range(10000)]
print(f"randint(0, 100) -- {time() - s}")

s = time()
_ = [random_int(0, 100) for _ in range(10000)]
print(f"random_int(0, 100) -- {time() - s}")
