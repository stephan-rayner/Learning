
def solution(A):
    N = len(A) + 1
    expected_sum = (N * (N + 1)) / 2
    return int(expected_sum - sum(A))

a = [1, 2, 3, 4, 5, 6, 7, 9]


print(solution(a))

