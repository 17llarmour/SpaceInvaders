import time as t
import requests
import json
from random import *
from pygame import *


# https://www.sitepoint.com/python-multiprocessing-parallel-programming/#:~:text=One%20way%20to%20achieve%20parallelism,multiprocessing%20accomplishes%20process%2Dbased%20parallelism.
# Potentially use that to always update cannon position, have it run in parallel with the rest of the code

def newGame():
    url = "http://localhost/reset" + "?reset=" + "yes"
    r = requests.post(url)
    url = "http://localhost/reset" + "?reset=" + "no"
    r = requests.post(url)


def getStates(pos):
    shoot = requests.post(url="http://localhost/state?place=yes")
    grid = requests.get("http://localhost/state")
    shooty = requests.get("http://localhost/shootyState")


if __name__ == '__main__':
    pass
