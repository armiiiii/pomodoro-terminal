package sessiontemplate

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

type Block struct {
	title string
	d     time.Duration
	icon  string
}

func (b Block) GetTitle() string {
	return b.title
}

func (b Block) GetDuration() time.Duration {
	return b.d
}

func (b Block) GetIcon() string {
	return b.icon
}

type SessionTemplate struct {
	title  string
	blocks []Block
}

func (s SessionTemplate) GetTitle() string {
	return s.title
}

func (s SessionTemplate) GetBlocks() []Block {
	return slices.Clone(s.blocks)
}

func CreateBlock(title string, d time.Duration, icon string) (Block, error) {
	var block Block

	if strings.TrimSpace(title) == "" {
		return block, fmt.Errorf("Title can't be blank for Block")
	}

	if d <= 0 {
		return block, fmt.Errorf("Duration must be greater than zero")
	}

	return Block{
		title: title,
		d:     d,
		icon:  icon,
	}, nil
}

func CreateSessionTemplate(title string, blocks []Block) (SessionTemplate, error) {
	var sessionTemplate SessionTemplate

	if strings.TrimSpace(title) == "" {
		return sessionTemplate, fmt.Errorf("Title can't be blank for Session Template")
	}

	if len(blocks) == 0 {
		return sessionTemplate, fmt.Errorf("Session Template must contain at least one block")
	}

	sessionTemplate = SessionTemplate{
		title:  title,
		blocks: blocks,
	}

	return sessionTemplate, nil
}
