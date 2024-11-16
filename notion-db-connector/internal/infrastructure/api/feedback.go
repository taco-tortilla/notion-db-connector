package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/taco-tortilla/notion-db-connector/config"
	"github.com/taco-tortilla/notion-db-connector/internal/domain"
	"github.com/taco-tortilla/notion-db-connector/internal/domain/repository"
	"github.com/taco-tortilla/notion-db-connector/internal/interfaces/request"
)

type feedbackAPI struct{}

func NewFeedbackAPI() repository.FeedbackRepository {
	return &feedbackAPI{}
}

func (fa feedbackAPI) Insert(fb *domain.Feedback) error {

	parent := request.Parent{DatabaseID: config.DbId}
	properties := request.Properties{
		CompanyName: request.CompanyName{
			Title: []request.Title{
				{
					Text: request.Text{
						Content: fb.CompanyName,
					},
				},
			},
		},
		ContractType: request.ContractType{
			Select: request.Select{
				Name: fb.ContractType.Value,
			},
		},
		Plan: request.Plan{
			Select: request.Select{
				Name: fb.Plan,
			},
		},
		Read: request.Read{
			Checkbox: false,
		},
		CreatedBy: request.CreatedBy{
			RichText: []request.RichText{
				{
					Type: "text", Text: request.Text{
						Content: fb.CreatedBy,
					},
				},
			},
		},
	}

	children := []request.Children{
		{
			Object: "block", Type: "paragraph", Paragraph: &request.Paragraph{
				RichText: []request.RichText{
					{
						Type: "text",
						Text: request.Text{Content: fb.Content},
					},
				},
			},
		},
	}

	request := request.NotionDBFeedbackRequest{
		Parent:     parent,
		Properties: properties,
		Children:   children,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", config.NotionURL+"pages", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", config.NotionVersion)
	req.Header.Set("Authorization", "Bearer "+config.NotionToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("API error: status %d, body: %s", res.StatusCode, string(body))
	}

	return nil
}
