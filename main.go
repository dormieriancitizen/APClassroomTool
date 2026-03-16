package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
)

type Question struct {
	IsMath     bool       `json:"is_math"`
	Stimulus   string     `json:"stimulus"`
	Options    []Option   `json:"options"`
	Metadata   Metadata   `json:"metadata"`
	Type       string     `json:"type"`
	Validation Validation `json:"validation"`
	ResponseID string     `json:"response_id"`
}

type Option struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type Metadata struct {
	CustomDistractorRationaleResponseLevel []string `json:"custom_distractor_rationale_response_level"`
	ValidResponseCount                     int      `json:"valid_response_count"`
	SheetReference                         string   `json:"sheet_reference"`
	WidgetReference                        string   `json:"widget_reference"`
	Source                                 Source   `json:"source"`
}

type Source struct {
	OrganisationID int `json:"organisation_id"`
}

type Validation struct {
	ValidResponse ValidResponse `json:"valid_response"`
	ScoringType   string        `json:"scoring_type"`
}

type ValidResponse struct {
	Score int      `json:"score"`
	Value []string `json:"value"`
}

type Root struct {
	Data Data `json:"data"`
}

type Data struct {
	ApiActivity ApiActivity `json:"apiActivity"`
}

type ApiActivity struct {
	QuestionsApiActivity QuestionsApiActivity `json:"questionsApiActivity"`
}

type QuestionsApiActivity struct {
	Questions []Question `json:"questions"`
}

func main() {
	jsonBytes, err := os.ReadFile("response.json")
	if err != nil {
		log.Fatal(err)
	}

	var root Root
	err = json.Unmarshal(jsonBytes, &root)

	if err != nil {
		log.Fatal(err)
	}

	for _, question := range root.Data.ApiActivity.QuestionsApiActivity.Questions {
		fmt.Println("Question: " + question.Stimulus)

		answered := false
		for i, option := range question.Options {
			if slices.Contains(question.Validation.ValidResponse.Value, option.Value) {
				fmt.Println("    AnswerID: " + string("ABCD"[i]))
				fmt.Println("    Answer: " + option.Label)
				answered = true
			}
		}

		if !answered {
			fmt.Println("Unknown answer: ", question.Validation.ValidResponse.Value)
		}
		fmt.Println()
		fmt.Scanln()
	}
}
