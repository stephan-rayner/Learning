from math import ceil

def solution(X, Y, D):
    return ceil( (Y - X) / D )

ans = solution(10, 85, 30)
print(ans)
