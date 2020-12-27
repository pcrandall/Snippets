#!/usr/bin/env python3

import requests
import sys

url_1 = "http://10.136.17.191:8080/lighthouse/login"
url = "http://10.136.17.191:8080/lighthouse-core/services/security/actors/current"

s = requests.Session()

r = s.get(url)

#prints webpage in terminal
#print(r.text)

# grabs the cookie
#print("session here" , s)
#print(dir(s.cookies._cookie_attrs))

# convert the cookie to string
tmp = s.cookies._cookie_attrs.__str__()
# split it up to just grab the cookie
tmp = tmp.split(", ")
tmp = tmp[2].split("'")

print("tmp here", tmp[1], '\n')

r = s.post("http://10.136.17.191:8080/lighthouse-core/services/security/login",
             data = {
                "key":"guest",
                "password":"guest"
                })

print("r_2 here", r_2)

# prints webpage in terminal

s.close()

# from console
# document.cookie = "check webpage for syntax"
