package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tealeg/xlsx"
)

// 从Excel里读取所有顶点及其相邻顶点
func readExcel(filePath string) ([][]string, error) {
	file, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	var rowValue [][]string
	for _, sheet := range file.Sheets { // 遍历所有工作表 本题中只有一个工作表
		for _, row := range sheet.Rows { // 遍历所有行
			var cellValue []string
			for _, cell := range row.Cells { // 遍历所有单元格
				if cell.Value != "" {
					cellValue = append(cellValue, cell.Value)
				}
			}
			rowValue = append(rowValue, cellValue)
		}
	}
	return rowValue, nil
}

// 随机挑选两个顶点u,v
func selectTwoVertices(allVertices [][]string) (string, string) {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	i := r.Intn(len(allVertices))
	u := allVertices[i][0]
	j := r.Intn(len(allVertices[i])-1) + 1
	v := allVertices[i][j]
	return u, v
}

// 把v变为u
func replaceVertice(u, v string, allVertices [][]string) [][]string {
	for i, row := range allVertices {
		for j := 1; j < len(row); j++ {
			if allVertices[i][j] == v {
				allVertices[i][j] = u
			}
		}
	}
	return allVertices
}

// 把和v相连的顶点和u相连,并删除顶点v
func connectVertices(u, v string, allVertices [][]string) [][]string {
	var adjencyV []string
	var newAllVertices [][]string
	for _, row := range allVertices {
		if row[0] == v {
			adjencyV = append(adjencyV, row[1:]...)
		}
		if row[0] != v {
			newAllVertices = append(newAllVertices, row)
		}
	}
	for i, row := range newAllVertices {
		if row[0] == u {
			newAllVertices[i] = append(newAllVertices[i], adjencyV...)
		}
	}
	return newAllVertices
}

// 删除self loop
func deleteSelfLoop(allVertices [][]string) [][]string {
	var newAllVertices [][]string
	for _, row := range allVertices {
		var newRow []string
		newRow = append(newRow, row[0])
		for _, vertice := range row {
			if vertice != row[0] {
				newRow = append(newRow, vertice)
			}
		}
		newAllVertices = append(newAllVertices, newRow)
	}
	return newAllVertices
}

// 计算min cut
func countMinCut(filePath string) int {
	allVertices, err := readExcel(filePath)
	if err != nil {
		fmt.Println("err---", err)
	}
	for {
		u, v := selectTwoVertices(allVertices)
		allVertices = replaceVertice(u, v, allVertices)
		allVertices = connectVertices(u, v, allVertices)
		allVertices = deleteSelfLoop(allVertices)
		if len(allVertices) == 2 {
			break
		}
	}
	return len(allVertices[0]) - 1
}

func main() {
	var min int
	for i := 0; i < 100; i++ {
		cut := countMinCut("/Users/yangziyun/Desktop/mincut.xlsx")
		if i == 0 {
			min = cut
		}
		if cut < min {
			min = cut
		}
		fmt.Println("cut--", cut)
	}
	fmt.Println()
	fmt.Println("min---", min)
}
