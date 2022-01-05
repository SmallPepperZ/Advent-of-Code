## advent of code 2021
## https://adventofcode.com/2021
## day 3

def parse_input(lines, by_bit_index=False):
    return lines

def most_common(data, bit_index):
    sorted_data = []
    for index in range(len(data[0])):
        index_data = []
        for line in data:
            index_data.append(int(line[index]))
        sorted_data.append(index_data)
    bit_data:list = sorted_data[bit_index]
    most_common = 1 if bit_data.count(1) >= bit_data.count(0) else 0 
    return most_common

def flip(value:int):
    return 0 if value == 1 else 1

def least_common(data, bit_index):
    sorted_data = []
    for index in range(len(data[0])):
        index_data = []
        for line in data:
            index_data.append(int(line[index]))
        sorted_data.append(index_data)
    bit_data:list = sorted_data[bit_index]
    return 0 if bit_data.count(0) <= bit_data.count(1) else 1

def part1(lines):
    gamma_list = []
    epsilon_list = []
    for bit_index in range(len(lines[0])):
        if most_common(lines, bit_index) == 1:
            gamma_list.append('1')
            epsilon_list.append('0')
        else:
            gamma_list.append('0')
            epsilon_list.append('1')
    gamma_int = int(''.join(gamma_list),2)
    epsilon_int = int(''.join(epsilon_list),2)
    
    return gamma_int * epsilon_int


def line_filter(line, value, bit):
    return int(line[bit]) == value


        


def part2(data:list):
    o2 = data
    co2 = data
    for i in range(len(data[0])):
        majority = most_common(o2, i)
        minority = least_common(co2, i)
        o2 = list(filter(lambda line: line_filter(line, majority,i), o2)) if len(o2) > 1 else o2
        co2 = list(filter(lambda line: line_filter(line, minority,i), co2)) if len(co2) > 1 else co2
    return int(o2[0], 2) * int(co2[0], 2)