#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the gridSearch function below.
def gridSearch(G, P):
    for index, gVal in enumerate(G):
        if (len(G)-index) < len(P):
            return "NO"
        pIndex = 0
        found = [[i, i+len(P[0])] for i in range(len(gVal)) if gVal.startswith(P[pIndex], i)]
        # found = [[x.start(), x.end()] for x in re.finditer(P[pIndex], gVal)]
        print(found)
        if len(found) != 0:
            for val in found:
                lst = [G[index][val[0]:val[1]]]
                for i in range(len(P) - 1):
                    lst.append(G[index+i+1][val[0]:val[1]])
                if lst == P:
                    return "YES"
            pIndex += 1

    # pValIndex = 0       
    # fromIndex = -1
    # for index, gVal in enumerate(G):
    #     print(index, gVal, pValIndex)
    #     found = gVal.find(P[pValIndex])
    #     if found != -1:
    #         if fromIndex != -1:
    #             print(gVal[fromIndex:fromIndex+len(P[0])], P[pValIndex])
    #             if gVal[fromIndex:fromIndex+len(P[0])] != P[pValIndex]:
    #                 if len(G)-index == len(P):
    #                     return "NO"
    #                 continue
    #         else:
    #             fromIndex = found
    #         if pValIndex == len(P)-1:
    #             return "YES"
    #         else:
    #             pValIndex += 1
    #     else:
    #         if len(G)-index < len(P):
    #             return "NO"



if __name__ == '__main__':

    t = int(input())

    for t_itr in range(t):
        RC = input().split()

        R = int(RC[0])

        C = int(RC[1])

        G = []

        for _ in range(R):
            G_item = input()
            G.append(G_item)

        rc = input().split()

        r = int(rc[0])

        c = int(rc[1])

        P = []

        for _ in range(r):
            P_item = input()
            P.append(P_item)

        result = gridSearch(G, P)

        print(result + '\n')
