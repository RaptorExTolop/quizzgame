package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aquasecurity/table"
)

type Question struct {
	name string
	//	description     string
	possibleAnswers []string
	correctAnswers  []string
}

type Questions []Question

func (qs *Questions) add(name string, possibleAnswers, correctAnswers []string) {
	question := Question{
		name: name,
		//		description:     description,
		possibleAnswers: possibleAnswers,
		correctAnswers:  correctAnswers,
	}

	*qs = append(*qs, question)
}

func (questions *Questions) validateIndex(index int) error {
	if index < 0 || index >= len(*questions) {
		err := errors.New("invalid index")
		return err
	}
	return nil
}

func (questions *Questions) delete(index int) error {
	qs := *questions

	if err := qs.validateIndex(index); err != nil {
		return err
	}

	*questions = append(qs[:index], qs[index+1:]...)

	return nil
}

func (questions *Questions) editName(index int, newName string) error {
	qs := *questions
	if err := qs.validateIndex(index); err != nil {
		return err
	}

	qs[index].name = newName
	return nil
}

func (questions *Questions) editPossibleAnswers(index int, newAnswers []string) error {
	qs := *questions
	if err := qs.validateIndex(index); err != nil {
		return err
	}

	qs[index].possibleAnswers = newAnswers
	return nil
}

func (questions *Questions) editCorrectAnswers(index int, newAnswers []string) error {
	qs := *questions

	if err := qs.validateIndex(index); err != nil {
		return err
	}

	qs[index].correctAnswers = newAnswers
	return nil
}

func (questions *Questions) printQuestions() error {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("Name", "Possible Answers", "Correct Answers")
	for _, qs := range *questions {
		var possibleAnswersStr string
		var correctAnswersStr string

		for _, i := range qs.possibleAnswers {
			possibleAnswersStr = fmt.Sprintf("%v, %v", possibleAnswersStr, i)
		}
		for _, i := range qs.correctAnswers {
			correctAnswersStr = fmt.Sprintf("%v, %v", correctAnswersStr, i)
		}
		possibleAnswersStr = strings.TrimPrefix(possibleAnswersStr, ", ")
		correctAnswersStr = strings.TrimPrefix(correctAnswersStr, ", ")

		table.AddRow(qs.name, possibleAnswersStr, correctAnswersStr)
	}

	table.Render()
	return nil
}
