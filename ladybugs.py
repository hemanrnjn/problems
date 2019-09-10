#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the happyLadybugs function below.
def happyLadybugs(b):
    s = ''.join(sorted(b))
    dct = {}
    prev = ''
    for ch in s:
        if ch in dct:
            dct[ch] += 1
        else:
            dct[ch] = 1
            if prev != '' and ord(prev) != ord('_') and dct[prev] < 2:
                return "NO"
            prev = ch
    if len(dct) > 0 and ord(prev) != ord('_') and dct[prev] < 2:
        return "NO"
    if '_' in dct:
        return "YES"
    else:
        return "NO"



if __name__ == '__main__':
    # fptr = open(os.environ['OUTPUT_PATH'], 'w')

    g = int(input())

    for g_itr in range(g):
        n = int(input())

        b = input()

        result = happyLadybugs(b)
        print(result)

    #     fptr.write(result + '\n')

    # fptr.close()
