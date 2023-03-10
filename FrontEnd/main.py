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
    url = "http://localhost/shoot?place=yes"
    shoot = requests.post(url)


def sendPlayer(pos):
    global lastCannonPos
    pos = float(pos / 60)
    pos = int(round(pos, 0))
    if pos > 29:
        pos = 29
    if pos == lastCannonPos:
        return
    lastCannonPos = pos
    # print(pos)
    url = "http://localhost/playerPos?pos=" + str(pos)
    # print(url)
    player = requests.post(url)


def getStates():
    gridTemp = requests.get("http://localhost/state")
    shootyTemp = requests.get("http://localhost/shootyState")
    grid = json.loads(gridTemp.text)
    shooty = json.loads(shootyTemp.text)
    # printGrid(grid)
    # printGrid(shooty)
    drawing(grid, shooty)


def getInfo():
    global lives, score
    r = requests.get("http://localhost/info")
    info = json.loads(r.text)
    lives, score = int(info[0]), int(info[1])


def printGrid(board):
    for i in range(15):
        print(board[i])
    print("----------SPLIT-----------")


def drawing(grid, shooty):
    global lives, score
    screen.fill((0, 0, 0))
    drawGrid(grid)
    drawShooty(shooty)
    writeScreen(lives, score)
    display.flip()


def drawGrid(grid):
    for y in range(15):
        for x in range(30):
            invaderImage = None
            if grid[y][x] == "6":
                invaderImage = image.load("redShip.png")
            elif x % 2 == 1:
                if grid[y][x] == "5" or grid[y][x] == "4":
                    invaderImage = image.load("invader3Odd.png").convert()
                elif grid[y][x] == "3" or grid[y][x] == "2":
                    invaderImage = image.load("invader2Odd.png").convert()
                elif grid[y][x] == "1":
                    invaderImage = image.load("invader1Odd.png").convert()
                elif grid[y][x] == "0":
                    invaderImage = image.load("cannonDraw.png").convert()
            else:
                if grid[y][x] == "5" or grid[y][x] == "4":
                    invaderImage = image.load("invader3Even.png").convert()
                elif grid[y][x] == "3" or grid[y][x] == "2":
                    invaderImage = image.load("invader2Even.png").convert()
                elif grid[y][x] == "1":
                    invaderImage = image.load("invader1Even.png").convert()
                elif grid[y][x] == "0":
                    invaderImage = image.load("cannonDraw.png").convert()
            if invaderImage != None:
                screen.blit(invaderImage, (x * 60, y * 60 + 60))


def drawShooty(grid):
    for y in range(15):
        for x in range(30):
            shootyImage = None
            if grid[y][x] == "4":
                draw.rect(screen, (0, 255, 0), (x * 60, y * 60 + 60, 60, 60))
            if grid[y][x] == "3":
                draw.rect(screen, (255, 255, 0), (x * 60, y * 60 + 60, 60, 60))
            if grid[y][x] == "2":
                draw.rect(screen, (255, 165, 0), (x * 60, y * 60 + 60, 60, 60))
            if grid[y][x] == "1":
                draw.rect(screen, (255, 0, 0), (x * 60, y * 60 + 60, 60, 60))
            if grid[y][x] == "p1":
                shootyImage = image.load("bullet1.png").convert()
                #draw.rect(screen, (255, 255, 255), (x * 60, y * 60 + 60, 60, 60))
            if grid[y][x] == "p2":
                #draw.rect(screen, (255, 0, 0), (x * 60, y * 60 + 60, 60, 60))
                shootyImage = image.load("bullet2.png").convert()
            if grid[y][x] == "p3":
                #draw.rect(screen, (255, 0, 0), (x * 60, y * 60 + 60, 60, 60))
                shootyImage = image.load("bullet3.png").convert()
            if grid[y][x] == "y":
                #draw.rect(screen, (255, 0, 0), (x * 60, y * 60 + 60, 60, 60))
                shootyImage = image.load("bullet4.png").convert()
            if shootyImage != None:
                screen.blit(shootyImage, (x * 60, y * 60 + 60))


def writeScreen(lives, score):
    fontObj = font.SysFont("Comic Sans MS", 50)
    img = fontObj.render("Score = " + str(score), True, (0, 0, 255))
    screen.blit(img, (0, 0))
    img = fontObj.render("Lives", True, (0, 0, 255))
    screen.blit(img, (1400, 0))
    x = 0
    for i in range(lives):
        img = image.load("cannonDraw.png").convert()
        screen.blit(img, (1560 + (i * 60 + x), 0))
        x += 20


def endScreen():
    screen.fill((255, 255, 255))
    fontObj = font.SysFont("Comic Sans MS", 120)
    img = fontObj.render("Your final score was " + str(score), True, (0, 0, 0))
    screen.blit(img, (0, 480))
    display.flip()
    t.sleep(3)


if __name__ == '__main__':
    lastCannonPos = int
    lives = 3
    score = 0
    print("running client")
    newGame()
    sendPlayer(0)
    init()
    width = 1800
    height = 960
    screen = display.set_mode((width, height))
    endProgram = False
    while not endProgram:
        # t.sleep(0.25)
        for e in event.get():
            if e.type == QUIT:
                endProgram = True
            if e.type == MOUSEBUTTONDOWN:
                shootyShoot()

        mouseX, mouseY = mouse.get_pos()
        sendPlayer(mouseX)
        getStates()
        getInfo()
        #print(lives, score)
        if lives == 0:
            endProgram = True

    endScreen()
    print("Closing Client")
