package sessiontemplate

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Block struct {
	title string
	d     time.Duration
	icon  string
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
	id     uuid.UUID
	title  string
	blocks []Block
}

func CreateSessionTemplate(title string, blocks []Block) (SessionTemplate, error) {
	var sessionTemplate SessionTemplate

	if strings.TrimSpace(title) == "" {
		return sessionTemplate, fmt.Errorf("Title can't be blank for Session Template")
	}

	if len(blocks) == 0 {
		return sessionTemplate, fmt.Errorf("Session Template must contain at least one block")
	}

	id, err := uuid.NewUUID()

	if err != nil {
		return SessionTemplate{}, err
	}

	sessionTemplate = SessionTemplate{
		id:     id,
		title:  title,
		blocks: blocks,
	}

	return sessionTemplate, nil
}

func (s SessionTemplate) ID() uuid.UUID {
	return s.id
}

func (s SessionTemplate) GetTitle() string {
	return s.title
}

func (s SessionTemplate) GetBlocks() []Block {
	return slices.Clone(s.blocks)
}
