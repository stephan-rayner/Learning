import time
import asyncio
import random

async def alpha(x):
    await asyncio.sleep(0.2)
    return x + 1 

async def bravo(x):
    await asyncio.sleep(0.2)
    return random.randint(0, 1000) + x

async def charlie(x):
    if x % 2 == 0:
        return x
    raise ValueError(x, 'is odd')

async def run():
    try:
        number = await charlie(
            await bravo(
                await alpha(42)
                )
            )
        
    except ValueError as exc:
        print('error:', exc.args[0])
    else:
        print('success:', number)

if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(run())
    loop.close()
