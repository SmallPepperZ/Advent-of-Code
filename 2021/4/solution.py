## advent of code 2021
## https://adventofcode.com/2021
## day 4

from typing import Any, Optional

class bcolors:
    CHECKED = '\033[92m' #GREEN
    UNMARKED = '\033[93m' #YELLOW
    RESET = '\033[0m'


class BingoNumber():
    checked:bool
    def __init__(self, number:int):
        self.checked = False
        self.num = number

    def check(self, number:int) -> bool:
        if number == self.num:
            self.checked = True
        return self.checked
    
    def __repr__(self) -> str:
        if self.checked:
            return str(bcolors.CHECKED + str(self.num) + bcolors.RESET)
        else:
            return str(bcolors.UNMARKED + str(self.num) + bcolors.RESET)

class BingoRow():
    nums:list[BingoNumber]
    def __init__(self, *, numbers:list[int]=None, bingo_numbers:list[BingoNumber] = None) -> None:
        nums = []
        if numbers:
            for num in numbers:
                nums.append(BingoNumber(num))
        elif bingo_numbers:
            for num in bingo_numbers:
                nums.append(num)
        self.nums = nums    
    
    @property
    def completed(self):
        return all([num.checked for num in self.nums])

    def check(self, number:int) -> bool:
        for num in self.nums:
            num.check(number)
        return self.completed

    def __iter__(self):
        return self.nums

    def __repr__(self) -> str:
        return ' '.join([f'{str(num):11}' for num in self.nums])

class BingoBoard():
    rows:list[BingoRow]
    numbers:list[BingoNumber]
    def __init__(self, numbers:list[list[int]]) -> None:
        rows = []
        for numlist in numbers:
            rows.append(BingoRow(numbers=numlist))
        self.numbers = []
        for row in rows:
            for number in row.nums:
                self.numbers.append(number)

        self.rows = rows

    @property
    def completed(self) -> bool:
        conditions = [
            any([row.completed for row in self.rows]),
            any([column.completed for column in self.columns]),
            # all([
            #     self.rows[0].nums[0].checked,
            #     self.rows[1].nums[1].checked,
            #     self.rows[2].nums[2].checked,
            #     self.rows[3].nums[3].checked,
            #     self.rows[4].nums[4].checked
            # ]),
            # all([
            #     self.rows[0].nums[4].checked,
            #     self.rows[1].nums[3].checked,
            #     self.rows[2].nums[2].checked,
            #     self.rows[3].nums[1].checked,
            #     self.rows[4].nums[0].checked
            # ])
        ]
        return any(conditions)

    def check(self, number:int):
        for row in self.rows:
            row.check(number)
        
    @property
    def columns(self) -> list[BingoRow]:
        columns:list[list[BingoNumber]] = []
        for i in range(len(self.rows[0].nums)):
            index_nums = []
            for row in self.rows:
                index_nums.append(row.nums[i])
            columns.append(index_nums)
        for i, column in enumerate(columns):
            columns[i] = BingoRow(bingo_numbers=column)
        return columns

    @property
    def unmarked(self) -> list[BingoNumber]:
        return [num.num for num in filter(lambda x: not x.checked, self.numbers)]

    def __repr__(self) -> str:
        return '\n'.join([str(row) for row in self.rows])

class Data():
    rolled:list[int]= []
    boards:list[BingoBoard] = []
    
    @property
    def winner(self) -> Optional[BingoBoard]:
        for board in self.boards:
            if board.completed:
                return board

    @property
    def winners(self) -> list[Optional[BingoBoard]]:
        boards= []
        for board in self.boards:
            if board.completed:
                boards.append(board)
        return boards

    def check_boards(self, number:int):
        for board in self.boards:
            board.check(number)
        return self.winner

    

def parse_input(lines:list[str]):
    input = Data()
    input.rolled = [int(val) for val in lines[0].split(",")]
    board_lines = lines[1:]
    for board in zip(*[iter(board_lines)]*6):
        board_data = []
        for line in board:
            board_data.append([int(num) for num in line.split()])
        board_data.pop(0)
        print(board_data)
        input.boards.append(BingoBoard(board_data))
    print(input.rolled)
    return input


def part1(data:Data):
    for number in data.rolled:
        if data.check_boards(number):
            return sum(data.winner.unmarked) * number
    return "No winning board found"

def part2(data:Data):
    for number in data.rolled:
        
        winner = data.check_boards(number)
        if winner:
            if len(data.boards) == 1:
                print(data.boards[0])
                return sum(data.boards[0].unmarked) * number
            else:
                for board in data.winners:
                    print(board)
                    print('\n')
                    data.boards.pop(data.boards.index(board))
        print(f'Roll: {number:10} | Boards: {len(data.boards)}')
        print(data.boards[1 if len(data.boards)>1 else 0].columns[2].nums)
        
    return "None found"