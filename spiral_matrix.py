# Given Matrix = [1  2  3
#                 4  5  6
#                 7  8  9]

import rotate_matrix

mat = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
length = len(mat[0])
c = length*length
row = col = 0
isRowIncr = isColIncr = True
for _ in range(c):
    if row == length-1:
        print(mat[row][col])
    elif col == length-1:
        row += 1
        print(mat[row][col])
    else:
        print(mat[row][col])
        col += 1


# Solution no. 2

# mat = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
# while True:
#     [print(i) for i in mat[0]]
#     del mat[0]
#     if len(mat) == 0:
#         break
#     mat = rotate_matrix.rotate(mat, 1)
