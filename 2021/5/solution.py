## advent of code 2021
## https://adventofcode.com/2021
## day 5

from typing import Tuple
from collections import Counter
from itertools import repeat

class direction():
    VERTICAL = 'v'
    HORIZONTAL = 'h'
    DIAGONAL = 'd'



class VentLine():
    def __init__(self, input:str):
        self.input = input
        start = input.split(" -> ")[0].split(",")
        end = input.split(" -> ")[1].split(",")
        self.start_x = int(start[0])
        self.start_y = int(start[1])
        self.end_x = int(end[0])
        self.end_y = int(end[1])

        self.direction = None
        self.coords = [(self.start_x,self.start_y), (self.end_x, self.end_y)]
        if self.start_x == self.end_x:
            self.direction = direction.VERTICAL
        elif self.start_y == self.end_y:
            self.direction = direction.HORIZONTAL
        else:
            self.direction = direction.DIAGONAL


    @property
    def points(self):
        coords = []
        if self.direction == direction.VERTICAL:
            diff = [self.start_y, self.end_y]
            diff.sort()
            for i in range(diff[0], diff[1]+1):
                coords.append((self.start_x, i))
        elif self.direction == direction.HORIZONTAL:
            diff = [self.start_x, self.end_x]
            diff.sort()
            for i in range(diff[0], diff[1]+1):
                coords.append((i, self.start_y))
        elif self.direction == direction.DIAGONAL:
            x_inc = 1 if (self.end_x - self.start_x) > 0 else -1
            y_inc = 1 if (self.end_y - self.start_y) > 0 else -1
            for i in range(x_inc*(self.end_x - self.start_x) + 1):
                coords.append((self.start_x+(i*x_inc), self.start_y+(i*y_inc)))


        return coords
    
    def __repr__(self) -> str:
        return f"<VentLine ({self.start_x}, {self.start_y} -> {self.end_x}, {self.end_y}) direction={self.direction}>"

class VentGrid():
    points:list[Tuple[int,int]] = []
    lines: list[VentLine] = []

    def check(self, line:VentLine):
        for point in line.points:
            if line.direction != direction.DIAGONAL:
                self.points.append(point)
        self.lines.append(line)


    @property
    def point_coords(self) -> Tuple[int, int]:
        return [point.coords for point in self.points]

    @property
    def danger_points(self):
        count = Counter(self.points)
        output = [item for item, n in count.items() if n>=2]
        return len(output)

class DiagVentGrid(VentGrid):
    points:list[Tuple[int,int]] = []
    lines: list[VentLine] = []

    def check(self, line:VentLine):
        for point in line.points:
            self.points.append(point)
        self.lines.append(line)


class Data():
    grid: VentGrid
    lines: list[VentLine] = []
    def __init__(self) -> None:
        self.grid = VentGrid()

def parse_input(lines):
    data = Data()
    for line in lines:
        data.lines.append(VentLine(line))
    return data

def part1(data:Data):
    grid = data.grid
    for line in data.lines:
        grid.check(line)
    return grid.danger_points

def part2(data:Data):
    grid = DiagVentGrid()
    for line in data.lines:
        grid.check(line)
    return grid.danger_points