# advent of code 2021
# https://adventofcode.com/2021
# day 8

from typing import Optional
import json
from collections import Counter

expected = {
    0: 'abcefg',
    1: 'cf',
    2: 'acdeg',
    3: 'acdfg',
    4: 'bcdf',
    5: 'abdfg',
    6: 'abdefg',
    7: 'acf',
    8: 'abcdefg',
    9: 'abcdfg'
}


class Mapping():
    pass


class Display():
    inputs: list[str]
    outputs: list[str]
    digits: dict[int, str] = {
        0: '',
        1: '',
        2: '',
        3: '',
        4: '',
        5: '',
        6: '',
        7: '',
        8: '',
        9: ''
    }
    mapping_options: dict[str, list[list[int]]] = {
        "a": [],
        "b": [],
        "c": [],
        "d": [],
        "e": [],
        "f": [],
        "g": []
    }

    def __init__(self, line: str) -> None:
        split = line.split(" | ")
        self.inputs = split[0].split()
        self.outputs = split[1].split()
        self.digits[1] = self.get_count(2)
        self.digits[4] = self.get_count(4)
        self.digits[7] = self.get_count(3)
        self.digits[8] = self.get_count(7)
        self.apply_nums()

    def get_count(self, count) -> Optional[str]:
        input = [input for input in self.inputs if len(input) == count]
        if len(input) == 1:
            return input[0]

    def __apply_num(self, digit, wires):
        for segment_wire in expected[digit]:
            print(f'{digit:2} | {wires:20} | {segment_wire}')
            if len(wires) > 0:
                self.mapping_options[segment_wire].append(list(wires))

    def apply_nums(self):
        for digit, wires in self.digits.items():
            if wires != '':
                self.__apply_num(digit, wires)

    def get_mappings(self):
        for letter, options in self.mapping_options.items():
            combined_options = []
            for option_list in options:
                combined_options += option_list
            count = Counter(combined_options)
            
            elegible = [item for item, n in count.items() if n == len(options)]
            print(f'{letter} | {len(options)} | {count} | {elegible}')
            self.mapping_options[letter] = elegible
        self.__mappings_recursive()
        

    def __mappings_recursive(self):
        old = self.mapping_options.copy()
        for letter, options in self.mapping_options.items():
            letter:str
            options:list[str]
            if list(self.mapping_options.values()).count(options) == len(options):
                for new_letter, new_options in self.mapping_options.items():
                    if new_options != options:
                        self.mapping_options[new_letter] = [option for option in new_options if not option in options]
            elif len(options) == 1:
                for new_letter, new_options in self.mapping_options.items():
                    if new_options != options:
                        self.mapping_options[new_letter] = [option for option in new_options if not option in options]
        if self.mapping_options != old:
            self.__mappings_recursive()

    def __repr__(self) -> str:
        return f'<DisplayObject mapped={self.digits}'


class SimpleDisplay():
    inputs = []
    outputs = []

    def __init__(self, line: str) -> None:
        split = line.split(" | ")
        self.inputs = split[0].split()
        self.outputs = split[1].split()


class Data():
    displays_s: list[SimpleDisplay] = []
    displays_c: list[Display] = []


def parse_input(lines):
    data = Data()
    for line in lines:
        data.displays_s.append(SimpleDisplay(line))
        data.displays_c.append(Display(line))
        break
    return data


def part1(data: Data):
    output = []
    for display in data.displays_s:
        output += [value for value in display.outputs if len(value) in (2, 4, 3, 7)]
    return len(output)


def part2(data: Data):
    data.displays_c[0].get_mappings()
    print(json.dumps(data.displays_c[0].mapping_options, indent=2, sort_keys=True))
