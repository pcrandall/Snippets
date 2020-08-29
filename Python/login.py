#!/usr/bin/env python3

import requests

url = "http://"

s = requests.Session()

r = s.post(url, data = {
    "login": "guest",
    "password": "guest",
    })
# prints webpage in terminal
print(r.text)

# grabs the cookie
print(s.cookies)

s.close()

# from console
# document.cookie = "check webpage for syntax"
