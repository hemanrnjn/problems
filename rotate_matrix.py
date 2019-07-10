def rotate(mat, dirn):
    new_mat = []
    num_row = len(mat)
    if isinstance(mat[0], type):
        num_col = 1
    else:
        num_col = len(mat[0])
    if dirn == 0:
        for i in range(num_col):
            x = []
            for j in range(num_row-1, -1, -1):
                x.append(mat[j][i])
            new_mat.append(x)
    else:
        for i in range(num_col-1, -1, -1):
            x = []
            for j in range(num_row):
                x.append(mat[j][i])
            new_mat.append(x)
    return new_mat

if __name__ == "__main__":
    num_row = int(input("Enter number of rows: \n"))
    num_col = int(input("Enter number of rows; \n"))
    dirn = int(input("Enter 1 for clockwise and 0 for anti-clockwise rotate: \n"))
    print("Enter space seperated integers of all", num_row, " rows")
    mat = []
    for i in range(num_row):
        mat.append(list(map(int, input().split(" "))))
    print(mat)
    print(rotate(mat, dirn))
