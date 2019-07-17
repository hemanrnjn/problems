# Given an array of distinct integers. The task is to count all the triplets such that sum of two elements equals the third element.

# Input:
# The first line of input contains an integer T denoting the number of test cases. Then T test cases follow. Each test case consists of two lines. First line of each test case contains an Integer N denoting size of array and the second line contains N space separated elements.

# Output:
# For each test case, print the count of all triplets, in new line. If no such triplets can form, print "-1".

# Constraints:
# 1 <= T <= 100
# 3 <= N <= 105
# 1 <= A[i] <= 106

# Example:
# Input:
# 2
# 4
# 1 5 3 2
# 3
# 3 2 7
# Output:
# 2
# -1

# Explanation:
# Testcase 1: There are 2 triplets: 1 + 2 = 3 and 3 +2 = 5

# Solution 1:
# if __name__ == "__main__":
#     num_input = int(input())
#     res = []
#     for i in range(num_input):
#         x = int(input())
#         arr = list(map(int, input().split(" ")))
#         # arr.sort()
#         for i in range(len(arr)):
#             for j in range(len(arr)):
#                 if i == j:
#                     continue
#                 else:
#                     for k in range(len(arr)):
#                         if k == i or k == j:
#                             continue
#                         else:
#                             if arr[k] == arr[i] + arr[j]:
#                                 a = [arr[i], arr[j], arr[k]]
#                                 a.sort()
#                                 if a in res:
#                                     continue
#                                 res.append(a)
#         print(res)
                
# Solution 2:
if __name__ == "__main__":
    num_input = int(input())
    res = {}
    count = 0
    for i in range(num_input):
        x = int(input())
        arr = list(map(int, input().split(" ")))
        for i in range(len(arr)):
            for j in range(i + 1, len(arr)):
                res[arr[i]+arr[j]] = [arr[i], arr[j]]
        print(res)
        for i in range(len(arr)):
            if arr[i] in res:
                count += 1
    print(count)
            
                    
