## advent of code 2021
## https://adventofcode.com/2021
## day 6
from collections import Counter
import json

def parse_input(lines):
    return [int(num) for num in lines[0].split(",")]

def part1(data):
    days = 1
    new_fish = []
    while days <= 80:
        for fish in data:
            if fish == 0:
                new_fish.append(6)
                new_fish.append(8)
            else:
                new_fish.append(fish-1)
        data = new_fish
        new_fish = []
        days += 1
    return len(data)


def part2(data):
    days = 1
    data = Counter(data)
    new_fish = {
            0:0,
            1:0,
            2:0,
            3:0,
            4:0,
            5:0,
            6:0,
            7:0,
            8:0,
        }
    while days <= 256:
        for timer, instances in data.items():
            if timer == 0:
                new_fish[6] += instances
                new_fish[8] += instances
            else:
                new_fish[timer-1] += instances
        data = new_fish
        new_fish = {
            0:0,
            1:0,
            2:0,
            3:0,
            4:0,
            5:0,
            6:0,
            7:0,
            8:0,
        }
        days += 1
    
    return(sum(data.values()))