#/usr/local/bin/python3

def sum_down(n):
    if n == 1:
        return n
    return n + sum_down(n-1)

def sum_down_tail(n, current_sum=0):
    current_sum = current_sum + n
    if n == 1:
        return current_sum
    return sum_down_tail(n-1, current_sum)

def sum_loop(n):
    return sum(range(1, n+1))

def sum_math(n):
    return int((n/2) * (n + 1))

#number = int(input("Number yo: "))
number = 999
result = sum_loop(number)
print(result)
