#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the happyLadybugs function below.
def happyLadybugs(b):
    dct = dict()
    prev = ''
    flag = False
    for ch in b:
        if ch in dct:
            dct[ch] += 1
            if ch != prev:
                flag = True
        else:
            dct[ch] = 1
            prev = ch
    for x, v in dct.items():
        if v < 2 and ord(x) != ord('_'):
            return "NO"
    if flag == False:
        return "YES"
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
