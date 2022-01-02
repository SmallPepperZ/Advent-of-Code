## advent of code 2021
## https://adventofcode.com/2021
## day 1

def parse_input(lines):
    return [int(line) for line in lines]

def part1(data):
    results = [] 
    for index, entry in enumerate(data, start=0):
        results.append(entry > data[index-1])
    return len(list(filter(None, results)))

def part2(data):
    clumps = []
    for index, entry in enumerate(data, start=0):
        try:
            clump = [entry, data[index+1], data[index+2]]
        except IndexError:
            break
        if len(clump) == 3:
            clumps.append(sum(clump))

    return part1(clumps)
    