package testfixtures

import (
	sessiontemplate "pomodoro/internal/session-template"
	"testing"
	"time"
)

func NewBlocks(t testing.TB) []sessiontemplate.Block {

	t.Helper()

	var blocks []sessiontemplate.Block

	input := []struct {
		title string
		d     time.Duration
		icon  string
	}{
		{
			title: "Recalling previous material",
			d:     15 * time.Minute,
			icon:  "",
		},
		{
			title: "Learning new topic",
			d:     25 * time.Minute,
			icon:  "",
		},
		{
			title: "Relax",
			d:     5 * time.Minute,
			icon:  "",
		},
		{
			title: "Practics",
			d:     35 * time.Minute,
			icon:  "",
		},
		{
			title: "Q&A with no hints",
			d:     10 * time.Minute,
			icon:  "",
		},
	}

	for _, rawBlock := range input {
		block, err := sessiontemplate.CreateBlock(rawBlock.title, rawBlock.d, rawBlock.icon)

		if err != nil {
			t.Fatalf("CreateBlock() returned an error %v", err)
		}

		blocks = append(blocks, block)
	}

	return blocks
}

func NewSessionTemplate(t testing.TB) sessiontemplate.SessionTemplate {

	t.Helper()

	var sessionTemplate sessiontemplate.SessionTemplate

	blocks := NewBlocks(t)

	sessionTemplate, err := sessiontemplate.CreateSessionTemplate("Interview Prep", blocks)

	if err != nil {
		t.Fatalf("CreateSessionTemplate() returned an error %v", err)
	}

	return sessionTemplate
}
