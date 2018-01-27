
# Find longest sequence of zeros in binary representation of an integer.

from random import randint

number = randint(0, 2147483647)

#number = 1041 #This is the test it should yield 5

binary_string = "{0:b}".format(number)
gap_list = binary_string.split("1")[:-1]
zero_lengths = [ len(x) for x in gap_list ]

print("Number: ", number)
print("Binary Number: ", binary_string)
print("Maximum Length of Zero Sequence: ",  max(zero_lengths))

