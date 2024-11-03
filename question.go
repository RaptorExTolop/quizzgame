package main

import (
	"errors"
	"os"

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
		table.AddRow(qs.name, "testing A", "testing B")
	}

	table.Render()
	return nil
}
