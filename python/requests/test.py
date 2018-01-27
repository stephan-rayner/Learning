"""
import requests

r = requests.get("http://www.reddit.com/r/babyelephants.json")
print(r.status_code)
"""

import urllib.request
import json
with urllib.request.urlopen("http://www.reddit.com/r/babyelephants.json") as response:
    stuff = response.read()
    print(json.loads(stuff)['data']['children'])
