from async_promises import Promise
#from promise import Promise
from time import sleep

def thing(resolve, reject):
    sleep(10)
    return resolve("CHEESE!")

promices = []
for i in range(100):
    promise = Promise(thing)
    promices.append(promise)

Promise.all(promices).then(print)
