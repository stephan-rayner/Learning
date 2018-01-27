
def solution(A):
    range_of_A = range(min(A), max(A) + 1)
    return (set(range_of_A) - set(A)).pop()

a = [1, 3, 6, 4, 2, 1]
print(solution(a))    
