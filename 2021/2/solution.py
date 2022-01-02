## advent of code 2021
## https://adventofcode.com/2021
## day 2


class Direction():
    def __init__(self, line:str):
        split = line.split(" ")
        self.direction = split[0]
        self.amount = int(split[1])
        self.vertical = 0
        self.horizontal = 0
        if self.direction == "forward":
            self.horizontal = self.amount
        elif self.direction == "up":
            self.vertical = -self.amount
        elif self.direction == "down":
            self.vertical = self.amount
        else:
            raise ValueError(f"direction '{self.direction}' not recognized")

class AimSubmarine():
    def __init__(self):
        self.aim = 0
        self.vertical = 0
        self.horizontal = 0
        

    def action(self, direction:str, amount:int):
        amount = int(amount)
        if direction == "forward":
            self.move(amount)
        elif direction == "up":
            self.aim -= amount
        elif direction == "down":
            self.aim += amount
        else:
            raise ValueError(f"direction '{direction}' not recognized")

    def move(self, amount):
        self.horizontal += amount
        self.vertical += self.aim * amount
    

def parse_input(lines):
    return lines

def part1(data):
    vertical = 0
    horizontal = 0
    for line in data:
        line = Direction(line)
        vertical += line.vertical
        horizontal += line.horizontal
    return vertical * horizontal

def part2(data):
    submarine = AimSubmarine()
    for line in data:
        split = line.split(" ")
        submarine.action(split[0], split[1])
    return submarine.vertical * submarine.horizontal