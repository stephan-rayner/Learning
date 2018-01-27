import asyncio
import time
import random

def alpha(x):
    time.sleep(0.2)
    return x + 1 

def bravo(x):
    time.sleep(0.2)
    return random.randint(0, 1000) + x

def charlie(x):
    if x % 2 == 0:
        return x
    raise ValueError(x, 'is odd')

def run():
    
    try:
        number = charlie(bravo(alpha(42)))
    except ValueError as exc:
        print('error:', exc.args[0])
    else:
        print('success:', number)

if __name__ == '__main__':
    run()
