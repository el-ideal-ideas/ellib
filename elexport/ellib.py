# -- Imports --------------------------------------------------------------------------

from ctypes import *
from pathlib import Path

# -------------------------------------------------------------------------- Imports --

# -- Variables --------------------------------------------------------------------------

ELLIB_SO: str = str(Path(__file__).parent.joinpath('ellib-export.so'))
ellib = cdll.LoadLibrary(ELLIB_SO)

# -------------------------------------------------------------------------- Variables --

# -- Wrapper --------------------------------------------------------------------------

ellib.similarity.restype = c_double
ellib.randomInt.restype = c_int


def similarity(s: str, t: str) -> float:
    """
    Calculate similarity between two strings.
    The result is between 0.0 and 1.0
    If the result is 0, two strings will be same.
    If the result is 1, two strings will be different completely.
    """
    return float(ellib.similarity(c_char_p(s.encode('utf-8')), c_char_p(t.encode('utf-8'))))


def random_int(min_value: int, max_value: int) -> int:
    """
    Please use `random.randint`
    This function is slower than it.
    Return a random int [min_value, max_value)
    """
    return int(ellib.randomInt(min_value, max_value))

# -------------------------------------------------------------------------- Wrapper --
