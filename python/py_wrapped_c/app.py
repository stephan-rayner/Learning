import ctypes

class POINT(ctypes.Structure):
    _fields_ = ("x", ctypes.c_int), ("y", ctypes.c_int)

lib = ctypes.CDLL("lib2wrap.so")

lib.dbl_int.restype = ctypes.c_int
c_dbl = lib.dbl_int(100)
print("C Result: ", c_dbl)

lib.hello.restype = ctypes.c_char_p
c_hello = lib.hello().decode("utf-8") # This comes out as a binary and I wanted a string
print(c_hello, "World :)")

# lib.make_point.argtypes = (ctypes.c_int, ctypes.POINTER(ctypes.c_int))
lib.make_point.argtypes = (ctypes.c_int, ctypes.c_int)
lib.make_point.restype = POINT
c_struct = lib.make_point(3, 4)
print("c_struct.x:{} and c_struct.y:{}".format(c_struct.x, c_struct.y))