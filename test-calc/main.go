package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var answerSheet = map[int]byte{
	1:  'D',
	2:  'D',
	3:  'B',
	4:  'C',
	5:  'B',
	6:  'B',
	7:  'A',
	8:  'D',
	9:  'A',
	10: 'B',
	11: 'A',
	12: 'A',
	13: 'A',
	14: 'B',
	15: 'C',
	16: 'B',
	17: 'A',
	18: 'A',
	19: 'C',
	20: 'B',
	21: 'A',
	22: 'C',
	23: 'A',
	24: 'B',
	25: 'D',
}

type Result struct {
	Student      string
	Score        uint
	WrongAnswers []uint
}

func main() {
	t := time.Now()
	defer func() {
		fmt.Printf("took %d microsecods\n", time.Since(t).Microseconds())
	}()
	fileNames := filesToRead("tests")
	files := openFiles("tests", fileNames)

	testResults(files)
}

func testResults(files <-chan *os.File) {
	answerFile, err := os.Create("results.csv")
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(answerFile)
	defer w.Flush()

	if err = w.Write([]string{"#", "Student", "Score"}); err != nil {
		panic(err)
	}

	n := 1
	for file := range files {
		go func(n int, f *os.File) {
			result := testResult(f)
			if err = w.Write([]string{strconv.Itoa(n), result.Student, strconv.Itoa(int(result.Score))}); err != nil {
				log.Panicf("IN WRITE: %v", err)
			}
		}(n, file)
		n++
	}
}

func testResult(f *os.File) Result {
	scanner := bufio.NewScanner(f)
	studentName := strings.TrimPrefix(strings.TrimSuffix(f.Name(), ".txt"), "tests/")
	score := 0
	wrongAnswers := make([]uint, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lineElems := strings.Fields(line)
		if len(lineElems) < 2 {
			log.Printf("student %s has invalid line in his file: %s", studentName, line)
			return Result{}
		}
		testNumber, err := strconv.Atoi(lineElems[0])
		if err != nil {
			log.Printf("student %s has invalid test number in line: %s", studentName, line)
			return Result{}
		}
		answer, ok := answerSheet[testNumber]
		if !ok {
			log.Printf("student %s has test number that is not present in answerSheet: %s", studentName, line)
			return Result{}
		}

		if answer == strings.ToUpper(lineElems[1])[0] {
			score++
		} else {
			wrongAnswers = append(wrongAnswers, uint(testNumber))
		}
	}

	return Result{
		Student:      studentName,
		Score:        uint(score),
		WrongAnswers: wrongAnswers,
	}
}

func openFiles(dir string, fileNames <-chan string) <-chan *os.File {
	files := make(chan *os.File, len(fileNames))

	go func() {
		defer close(files)

		for name := range fileNames {
			file, err := os.Open(fmt.Sprintf("%s/%s", dir, name))
			if err != nil {
				panic(err)
			}

			files <- file
		}
	}()

	return files
}

func filesToRead(dir string) <-chan string {
	fileNames := make(chan string, 10)

	go func() {
		defer close(fileNames)

		dirEntries, err := os.ReadDir(dir)
		if err != nil {
			panic(err)
		}

		for _, entry := range dirEntries {
			if !entry.Type().IsRegular() {
				fmt.Println("entry is not a file:", entry)
				continue
			}
			fileNames <- entry.Name()
		}
	}()

	return fileNames
}
