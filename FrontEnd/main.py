import time as t
import requests
import json
from random import *
from pygame import *

def newGame():
    url = "http://localhost/reset" + "?reset=" + "yes"
    r = requests.post(url)
    url = "http://localhost/reset" + "?reset=" + "no"
    r = requests.post(url)


if __name__ == '__main__':
    print_hi('PyCharm')

