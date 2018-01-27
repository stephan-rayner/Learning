# Rotate an array to the right by a given number of steps.

a = [1, 2, 3, 4, 5, 6, 7, 8, 9]

list.pop(a)

def shift(A, K):
    if len(A) == 0:
        return A
     
    n = K % len(A) 
    for i in range(n):
        new_head = [A.pop()]
        new_head.extend(A)
        A = new_head
    return A

shift_number = int(input("number please: "))

print("Original array: ", a)
print("Shifted array: ", shift(a, shift_number))

def test():
    start = [3, 8, 9, 7, 6]
    finish = shift(start, 3)
    assert finish == [9, 7, 6, 3, 8]
test()
