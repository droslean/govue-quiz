package results

import (
	"testing"

	"github.com/droslean/govue-quiz/pkg/quiz"
)

func TestGetRankOfTotalResults(t *testing.T) {
	testResults := struct {
		id       string
		results  *Results
		expected string
	}{
		id: "test",
		results: &Results{
			Results: []Result{
				{
					Percentage: "40.00",
				},
				{
					Percentage: "21.00",
				},
				{
					Percentage: "34.40",
				},
				{
					Percentage: "73.90",
				},
			},
		},
		expected: "42.33",
	}

	if res := getRankOfTotalResults(testResults.results); res != testResults.expected {
		t.Fatalf("Expected %s and got %s", testResults.expected, res)
	}

}

func TestGetCorrectAnswersAndPercentage(t *testing.T) {
	testResults := struct {
		id                     string
		questions              []quiz.Question
		answers                []string
		expectedCorrectAnswers int
		expectedPercentage     float32
	}{
		id: "test",
		questions: []quiz.Question{
			{
				CorrectAnswer: "correct",
			},
			{
				CorrectAnswer: "correct",
			},
			{
				CorrectAnswer: "correct",
			},
			{
				CorrectAnswer: "correct",
			},
		},
		answers:                []string{"correct", "incorrect", "incorrect", "correct"},
		expectedCorrectAnswers: 2,
		expectedPercentage:     50.00,
	}

	if correctAnswers, perc := getCorrectAnswersAndPercentage(testResults.questions, testResults.answers); correctAnswers != testResults.expectedCorrectAnswers || perc != testResults.expectedPercentage {
		t.Fatalf("Expected %d - %.2f and got %d - %.2f", testResults.expectedCorrectAnswers, testResults.expectedPercentage, correctAnswers, perc)
	}
}
