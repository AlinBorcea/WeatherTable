package main

import "fmt"

type TableLine struct {
	nums []int
}

type Table struct {
	lines []TableLine
}

type Data struct {
	tables []Table
}

func main() {
	table := initialTable()
	data := Data{}

	for _, line := range table.lines {
		data.tables = append(data.tables, tableFromLine(line))
	}

	fmt.Println(data.tables)
}

func initialTable() Table {
	return Table{[]TableLine{
		{[]int{3, 6, 9, 12, 15, 18}},
		{[]int{1, 2, 3, 4}},
		{[]int{3, 8, 13, 18}},
	},
	}
}

func tableFromLine(line TableLine) Table {
	t := Table{}

	t.lines = append(t.lines, lineFromLine(line))
	for i := 0; !lineIsAllZero(t.lines[i]); i++ {
		t.lines = append(t.lines, lineFromLine(t.lines[i]))
	}

	return t
}

func lineFromLine(line TableLine) TableLine {
	l := TableLine{}

	for i := 0; i < len(line.nums)-1; i++ {
		l.nums = append(l.nums, line.nums[i+1]-line.nums[i])
	}

	return l
}

func lineIsAllZero(l TableLine) bool {
	for _, x := range l.nums {
		if x != 0 {
			return false
		}
	}
	return true
}
