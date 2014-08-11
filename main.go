package main

import (
	"fmt"
	"math/rand"
	"time"
)

func NewDie(sides int) Die {
	rand.Seed(time.Now().UnixNano())
	return Die{sides, Log{}}
}

type Log struct {
	log []LogEntry
}

func (l *Log) addSimpleEntry(i int) LogEntry {
	entry := simpleEntry(i)
	l.log = append(l.log, entry)
	return entry
}

func simpleEntry(i int) LogEntry {
	return LogEntry{i, 0, []int{}}
}

type LogEntry struct {
	result        int
	numberRolled  int
	droppedValues []int
}

type Die struct {
	Sides int
	Log   Log
}

func (d *Die) Roll() int {
	return rollSingle(d.Sides)
}

func (d *Die) RollMany(num int) (result int, entry LogEntry) {
	result = (rollSingle(d.Sides)) * num
	entry = d.Log.addSimpleEntry(result)
	return result, entry
}

func rollSingle(sides int) int {
	return rand.Intn(sides) + 1
}

//func rollMany(sides int, num int) ( total int, entry LogEntry) {
//for _, i in num
//end
//}

func main() {
	d := D6()
	fmt.Println(d.Roll())
	fmt.Println(d.Roll())
	fmt.Println(d.Roll())
	fmt.Println(d.Roll())
	fmt.Println(d.Roll())
	fmt.Println(d.Roll())
	fmt.Println(d.Roll())
}

func D6() Die {
	return NewDie(6)
}

func D3() Die {
	return NewDie(3)
}

func D4() Die {
	return NewDie(4)
}

func D12() Die {
	return NewDie(12)
}

func D20() Die {
	return NewDie(20)
}
