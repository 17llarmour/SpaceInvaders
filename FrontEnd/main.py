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


def shootyShoot():
    shoot = requests.post(url="http://localhost/state?place=yes")


def sendPlayer(pos):
    pos = float(pos/30)
    pos = int(round(pos,0))
    player = requests.post(url="http://localhost/player?pos=" + str(pos))



def getStates():
    gridTemp = requests.get("http://localhost/state")
    shootyTemp = requests.get("http://localhost/shootyState")
    grid = json.loads(gridTemp.text)
    shooty = json.loads(shootyTemp.text)
    printGrid(grid)
    printGrid(shooty)


def printGrid(board):
    for i in range(15):
        print(board[i])
    print("----------SPLIT-----------")


if __name__ == '__main__':
    print("running client")
    init()
    width = 900
    height = 450
    screen = display.set_mode((width,height))
    endProgram = False
    newGame()
    while not endProgram:
        for e in event.get():
            if e.type == QUIT:
                endProgram = True

        mouseX, mouseY = mouse.get_pos()
