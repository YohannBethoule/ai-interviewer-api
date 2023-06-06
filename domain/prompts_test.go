package domain

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGenerateBasePromptWithoutInterviewerPosition(t *testing.T) {
	want := fmt.Sprintf(promptFormat, "jobTitle", "", 5, "jobDescription")
	got := GenerateBasePrompt("jobTitle", 5, "", "jobDescription")
	assert.Equal(t, want, got)
}

func TestGenerateBasePromptWithInterviewerPosition(t *testing.T) {
	want := fmt.Sprintf(promptFormat, "jobTitle", fmt.Sprintf(promptInterviewerPositionFormat, "interviewerPosition"), 5, "jobDescription")
	got := GenerateBasePrompt("jobTitle", 5, "interviewerPosition", "jobDescription")
	assert.Equal(t, want, got)
}
