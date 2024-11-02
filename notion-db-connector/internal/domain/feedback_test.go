package domain

import "testing"

func TestNewFeedbackCompanyName(t *testing.T) {
	expected := "xx株式会社"
	companyName := expected
	contractType := "contract"
	plan := "S1"
	createdBy := "田中太郎"
	content := "サンプル文章"

	result, err := NewFeedback(companyName, contractType, plan, createdBy, content)

	if err != nil {
		t.Fatalf("expected no error, but got %q", err)
	}

	if result.CompanyName != expected {
		t.Fatalf("expected %q, but got %q", expected, result.CompanyName)
	}
}

func TestNewFeedbackCompanyNameEmpty(t *testing.T) {
	companyName := ""
	contractType := "contract"
	plan := "S1"
	createdBy := "田中太郎"
	content := "サンプル文章"

	result, err := NewFeedback(companyName, contractType, plan, createdBy, content)

	if err == nil {
		t.Fatalf("expected error, but got %+v", result)
	}
}

func TestNewFeedbackContent(t *testing.T) {
	expected := "サンプル文章"
	companyName := "xx株式会社"
	contractType := "contract"
	plan := "S1"
	createdBy := "田中太郎"
	content := expected

	result, err := NewFeedback(companyName, contractType, plan, createdBy, content)

	if err != nil {
		t.Fatalf("expected no error, but got %q", err)
	}

	if result.Content != expected {
		t.Fatalf("expected %q, but got %q", expected, result.Content)
	}
}

func TestNewFeedbackContentEmpty(t *testing.T) {
	companyName := "xx株式会社"
	contractType := "contract"
	plan := "S1"
	createdBy := "田中太郎"
	content := ""

	result, err := NewFeedback(companyName, contractType, plan, createdBy, content)

	if err == nil {
		t.Fatalf("expected error, but got %+v", result)
	}
}

func TestNewFeedbackCreatedBy(t *testing.T) {
	expected := "田中太郎"
	companyName := "xx株式会社"
	contractType := "contract"
	plan := "S1"
	createdBy := expected
	content := "サンプル文章"

	result, err := NewFeedback(companyName, contractType, plan, createdBy, content)

	if err != nil {
		t.Fatalf("expected no error, but got %q", err)
	}

	if result.CreatedBy != expected {
		t.Fatalf("expected %q, but got %q", expected, result.CreatedBy)
	}
}

func TestNewFeedbackCreatedByEmpty(t *testing.T) {
	companyName := "xx株式会社"
	contractType := "contract"
	plan := "S1"
	createdBy := ""
	content := "サンプル文章"

	result, err := NewFeedback(companyName, contractType, plan, createdBy, content)

	if err == nil {
		t.Fatalf("expected error, but got %+v", result)
	}
}
