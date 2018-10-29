package quiz

import (
	"net/http"

	"github.com/labstack/echo"
)

// User holds all the information of the player
type User struct {
	Name       string `json:"name" form:"name" query:"name"`
	Email      string `json:"email" form:"email" query:"email"`
	Difficulty string `json:"difficulty" form:"difficulty" query:"difficulty"`
}

// Quiz holds all the information is needed to start the quiz game.
type Quiz struct {
	User          *User
	QuestionsPack []Question
}

// NewQuiz creates a Quiz filled up with all the questions.
func NewQuiz(user *User) *Quiz {
	q := NewQuestions()
	return &Quiz{
		User:          user,
		QuestionsPack: q.getQuestions(user.Difficulty),
	}
}

// StartQuiz returns the logic to start the quiz game
func StartQuiz(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	quiz := NewQuiz(u)
	return c.JSON(http.StatusOK, quiz)
}
