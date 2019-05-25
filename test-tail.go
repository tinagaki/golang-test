package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {

	var (
		line_num = flag.Int("n", 10, "line number")
	)
	flag.Parse()
	args := flag.Args()
	// 複数ファイル受ける対応
	for _, file := range args {
		err := mytail(file, *line_num)
		if err != nil {
			os.Exit(0)
		}
	}
}
func mytail(file string, line_num int) (err error) {
	f, err := os.Open(file)
	if err != nil {
		return errors.New("file not exist")
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if len(lines) >= line_num {
			// 要素カット
			lines = append(lines[:0], lines[1:]...)
		}
		lines = append(lines, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		return errors.New("line scan error")
	}
	fmt.Println("=====", file, "=====")
	for i := range lines {
		fmt.Println(lines[i])
	}
	return nil
}
