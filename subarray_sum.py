# Given an unsorted array A of size N of non-negative integers, find a continuous sub-array which adds to a given number S.

# Input:
# The first line of input contains an integer T denoting the number of test cases. Then T test cases follow. Each test case consists of two lines. The first line of each test case is N and S, where N is the size of array and S is the sum. The second line of each test case contains N space separated integers denoting the array elements.

# Output:
# For each testcase, in a new line, print the starting and ending positions(1 indexing) of first such occuring subarray from the left if sum equals to subarray, else print -1.

# Constraints:
# 1 <= T <= 100
# 1 <= N <= 107
# 1 <= Ai <= 1010

# Example:
# Input:
# 2
# 5 12
# 1 2 3 7 5
# 10 15
# 1 2 3 4 5 6 7 8 9 10
# Output:
# 2 4
# 1 5

def cal_sub_array(num_input, arr, x):
    for i, _ in enumerate(arr):
        s = 0
        for j in range(i, len(arr)):
            s = s + arr[j]
            if s == x:
                res = str(i+1) + "  " + str(j+1)
                return res
            if s > x:
                break
    return ""

if __name__ == "__main__":
    num_input = int(input())
    res = []
    for i in range(num_input):
        x = int(input().split(" ")[1])
        arr = list(map(int, input().split(" ")))
        res.append(cal_sub_array(num_input, arr, x))
    for i in res:
        print(i)