# Find value that occurs in odd number of elements.

initial_list = [1, 1, 2, 3, 4, 4, 4]
test_list = [9, 3, 9, 3, 9, 7, 9]
initial_list = test_list

counts = {}

for i in initial_list:
    if i in counts:
        counts[i] = counts[i] + 1
    else:
        counts[i] = 1
odd_values = []
for j in counts:
    if counts[j] % 2 != 0:
      odd_values.append(j)

print(odd_values) 
