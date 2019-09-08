# Given an array arr of N integers. Find the contiguous sub-array with maximum sum.

# Input:
# The first line of input contains an integer T denoting the number of test cases. The description of T test cases follows. The first line of each test case contains a single integer N denoting the size of array. The second line contains N space-separated integers A1, A2, ..., AN denoting the elements of the array.

# Output:
# Print the maximum sum of the contiguous sub-array in a separate line for each test case.

# Constraints:
# 1 ≤ T ≤ 110
# 1 ≤ N ≤ 106
# -107 ≤ A[i] <= 107

# Example:
# Input
# 2
# 5
# 1 2 3 -2 5
# 4
# -1 -2 -3 -4
# Output
# 9
# -1

# Explanation:
# Testcase 1: Max subarray sum is 9 of elements (1, 2, 3, -2, 5) which is a contiguous subarray.

# -10 8 1 5 12 -4 1 2
# -10 8 1 5 12 -40 1 2

if __name__ == "__main__":
    num_input = int(input())
    curr_sum = max_sum = startIndex = endIndex = 0
    for i in range(num_input):
        x = int(input())
        arr = list(map(int, input().split(" ")))
        for i in range(len(arr)):
            curr_sum = curr_sum + arr[i]
            if curr_sum < 0:
                curr_sum = 0
                if i is not len(arr)-1:
                    startIndex = i+1
            if curr_sum > max_sum:
                max_sum = curr_sum
                endIndex = i
        print(arr[startIndex], arr[endIndex])

        