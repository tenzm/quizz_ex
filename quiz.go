package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type QuestionReader interface {
	ParseQuestions(r io.Reader) ([]Question, error)
}

type Question struct {
	question string
	answer   string
}

func readCsv(filename string) []Question {
	csvfile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	question := make([]Question, len(records))
	for i := range records {
		question[i].question = records[i][0]
		question[i].answer = records[i][1]
	}
	// Допишите код здесь
	return question
}

func main() {
	total := 0
	questions := readCsv("problems.csv")
	for i, _ := range questions {
		fmt.Println(questions[i].question)
		var answ string
		fmt.Scan(&answ)
		if answ == questions[i].answer {
			total += 1
		}
	}
	fmt.Println("You got " + strconv.Itoa(total) + "/" + strconv.Itoa(len(questions)))
	// Пройтись циклом. Вывести вопрос, предложить пользователю ввести ответ.
	// Если ответ правильный, увеличить total.
	// for
}
