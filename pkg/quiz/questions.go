package quiz

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Questions hold all the pack of questions.
type Questions struct {
	Questions QuestionsPacks `json:"questions"`
}

// QuestionsPacks holds the question packs depending the difficulty.
type QuestionsPacks struct {
	Easy   []Question `json:"easy"`
	Medium []Question `json:"medium"`
	Hard   []Question `json:"hard"`
}

// Question holds all the information for each question.
type Question struct {
	Category      string   `json:"category"`
	Type          string   `json:"type"`
	Difficulty    string   `json:"difficulty"`
	CorrectAnswer string   `json:"correct_answer"`
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
}

// NewQuestions reads all the questions from a file, and generates data.
func NewQuestions() *Questions {
	var data Questions

	jsonFile, err := os.Open("static/quiz.json")
	if err != nil {
		return nil
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil
	}

	json.Unmarshal(byteValue, &data)
	return &data
}

func (q *Questions) getQuestions(difficulty string) []Question {
	var finalQuestions []Question

	switch difficulty {
	case "Easy":
		finalQuestions = q.Questions.Easy
	case "Medium":
		finalQuestions = q.Questions.Medium
	case "Hard":
		finalQuestions = q.Questions.Hard
	default:
		return nil
	}
	return finalQuestions
}
