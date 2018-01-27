collection = [3, 1, 2, 4, 3]

def solution(A):
    left = 0
    right = sum(A)
    minimal_split_point = A[0]

    for val in A:
        left = left + val
        right = right - val
        current = abs(left - right)
        if current < minimal_split_point:
            minimal_split_point = current
    
    return minimal_split_point

print(solution(collection))

