## advent of code 2021
## https://adventofcode.com/2021
## day 7

from statistics import median

def cost(distance:int) -> int:
    return (distance**2+distance)/2



def avg(data:list[int]):
    return sum(data)/len(data)

def parse_input(lines):
    return [int(val) for val in lines[0].split(",")]

def part1(data):
    target = int(median(data))
    fuel = 0
    for val in data:
        fuel += abs(val-target)
    return int(fuel)


def part2(data): # Very janky, not sure if it works in all situations
    targetA = int(avg(data))
    targetB = int(avg(data))+1
    fuelA = 0
    fuelB = 0
    for val in data:
        fuelA += cost(abs(val-targetA))
        fuelB += cost(abs(val-targetB))
    return int(min(fuelA, fuelB))