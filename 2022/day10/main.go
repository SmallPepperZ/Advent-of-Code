package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Register struct {
	value        int
	cycle        int
	done         bool
	instructions list.List
}

func (register *Register) Increase(value int) {

	register.value += value
}

func (register *Register) AddTask(task *Instruction) {
	register.instructions.PushBack(task)
}

func (register *Register) RunCycle(input *bufio.Scanner) bool {
	register.cycle++

	if register.instructions.Len() != 0 {
		element := register.instructions.Front()
		task := element.Value.(*Instruction)
		task.cycle--
		if task.cycle == 0 {
			register.Increase(task.value)
			register.instructions.Remove(element)
		}
	}
	if register.instructions.Len() == 0 {
		if register.done {
			return false
		}
		if input.Scan() {
			text := input.Text()
			if text == "noop" {
				register.instructions.PushBack(&Instruction{value: 0, cycle: 1})
			} else {
				value, _ := strconv.Atoi(strings.TrimSpace(strings.Replace(text, "addx", "", 1)))
				register.instructions.PushBack(&Instruction{value: value, cycle: 2})
			}
		} else {
			if register.instructions.Len() == 0 {
				register.done = true
			}
		}
	}
	return true
}

type Instruction struct {
	value int
	cycle int
}

func (task Instruction) String() string {
	return fmt.Sprintf("<%v:%v>", task.value, task.cycle)
}

func main() {
	register := &Register{value: 1, cycle: 0}
	file, err := os.Open("input.txt")
	pixels := make([]bool,40*6)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sum := 0
	for register.RunCycle(scanner) {
		if register.cycle > 220 {
			break
		}
		if (register.cycle-20)%40 == 0 {
			fmt.Println(register.cycle, register.value, register.cycle*register.value)
			sum += register.cycle * register.value
		}
	}
	for row:=0;row<6;row++ {
		for col:=0;col<40;col++ {
			pixel := pixels[row*40+col]
			if (pixel) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println(sum)
}

// func runCycle(register *Register, instructions *list.List, text string) {
// 	if text == "noop" {
// 	} else {
// 		value, _ := strconv.Atoi(strings.TrimSpace(strings.Replace(text, "addx", "", 1)))
// 		instructions.PushBack(&Instruction{value: value, cycle: register.cycle + 2})
// 	}

// 	for {
// 		register.cycle++
// 		element := instructions.Front()
// 		if element != nil {
// 			isDone := element.Value.(*Instruction).Run(register)
// 			if isDone {
// 				instructions.Remove(element)
// 			}
// 		} else {
// 			break
// 		}

// 		fmt.Println(register.cycle, register.value)
// 	}

// }
