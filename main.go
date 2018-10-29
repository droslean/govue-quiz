package main

import (
	"flag"

	"github.com/droslean/govue-quiz/pkg/quiz"
	"github.com/droslean/govue-quiz/pkg/results"

	"github.com/labstack/echo"
)

type flags struct {
	port string
}

func parseFlags() *flags {
	flags := flags{}
	flag.StringVar(&flags.port, "port", "3000", "Assigning the port that the server should listen on.")
	flag.Parse()
	return &flags
}

func main() {
	flags := parseFlags()
	e := echo.New()

	e.File("/", "static/index.html")
	e.Static("/static", "static")

	// Routes
	e.POST("/quiz", quiz.StartQuiz)
	e.POST("/results", results.UpdateResults)

	e.Start(":" + flags.port)
}
