package results

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/droslean/govue-quiz/pkg/quiz"
	"github.com/labstack/echo"
)

// Results holds the list of all results.
type Results struct {
	Results []Result `json:"results"`
}

// Result holds the information for each result pack.
type Result struct {
	User           *quiz.User      `json:",inline"`
	Questions      []quiz.Question `json:"questions"`
	Answers        []string        `json:"answers"`
	Percentage     string          `json:"percentage"`
	CorrectAnswers string          `json:"correct_answers"`
}

// TotalResults holds the history of all results.
var TotalResults = NewResults()

// NewResults creates an empty Result struct.
func NewResults() *Results {
	return &Results{}
}

// UpdateResults updates all the results in memory and returns the
// required information for the user.
func UpdateResults(c echo.Context) error {
	u := new(Result)
	if err := c.Bind(u); err != nil {
		return err
	}

	correctAnswers, percentage := getCorrectAnswersAndPercentage(u.Questions, u.Answers)
	u.Percentage = fmt.Sprintf("%.2f", percentage)
	u.CorrectAnswers = strconv.Itoa(correctAnswers)

	res := struct {
		CorrectAnswers string `json:"correct_answers"`
		Percentage     string `json:"percentage"`
		Rank           string `json:"rank"`
	}{
		CorrectAnswers: u.CorrectAnswers,
		Percentage:     u.Percentage,
		Rank:           getRankOfTotalResults(TotalResults),
	}

	// Append global results
	TotalResults.Results = append(TotalResults.Results, *u)

	return c.JSON(http.StatusOK, res)
}

func getCorrectAnswersAndPercentage(questions []quiz.Question, answers []string) (int, float32) {
	correctAnswers := 0
	for index, answer := range answers {
		if answer == questions[index].CorrectAnswer {
			correctAnswers++
		}
	}

	if correctAnswers == 0 {
		return 0, 0.00
	}
	return correctAnswers, (float32(correctAnswers) / float32(len(questions))) * 100
}

func getRankOfTotalResults(results *Results) string {
	if len(results.Results) == 0 {
		return "100"
	}

	var totalPercentage float64
	for _, result := range results.Results {
		perc, err := strconv.ParseFloat(result.Percentage, 64)
		if err != nil {
			return "0.00"
		}
		totalPercentage += float64(perc)
	}

	return fmt.Sprintf("%.2f", totalPercentage/float64(len(results.Results)))
}
