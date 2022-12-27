import time as t
import requests
import json
from random import *
from pygame import *


# https://www.sitepoint.com/python-multiprocessing-parallel-programming/#:~:text=One%20way%20to%20achieve%20parallelism,multiprocessing%20accomplishes%20process%2Dbased%20parallelism.
# Potentially use that to always update cannon position, have it run in parallel with the rest of the code
lastCannonPos = int


def newGame():
    url = "http://localhost/reset" + "?reset=" + "yes"
    r = requests.post(url)
    url = "http://localhost/reset" + "?reset=" + "no"
    r = requests.post(url)


def shootyShoot():
    url = "http://localhost/shoot?place=yes"
    shoot = requests.post(url)
    

def sendPlayer(pos):
    global lastCannonPos
    pos = float(pos/60)
    pos = int(round(pos, 0))
    if pos > 29:
        pos = 29
    if pos == lastCannonPos:
        return
    lastCannonPos = pos
    print(pos)
    url = "http://localhost/playerPos?pos=" + str(pos)
    print(url)
    player = requests.post(url)



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
    newGame()
    init()
    width = 1800
    height = 900
    screen = display.set_mode((width,height))
    endProgram = False
    newGame()
    while not endProgram:
        for e in event.get():
            if e.type == QUIT:
                endProgram = True
            if e.type == MOUSEBUTTONDOWN:
                shootyShoot()

        mouseX, mouseY = mouse.get_pos()
        sendPlayer(mouseX)
        #getStates()