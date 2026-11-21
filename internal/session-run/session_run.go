package sessionrun

import (
	"errors"
	sessiontemplate "pomodoro/internal/session-template"
)

var (
	ErrSessionEnded = errors.New("session alreay ende")
)

type SessionRun struct {
	blocks      []sessiontemplate.Block
	activeIndex int
	ended       bool
}

func CreateSessionRun(sessionTemplate sessiontemplate.SessionTemplate) (SessionRun, error) {
	return SessionRun{
		activeIndex: 0,
		blocks:      sessionTemplate.GetBlocks(),
	}, nil
}

func (s *SessionRun) isLastBlock() bool {
	return s.activeIndex >= len(s.blocks)-1
}

func (s *SessionRun) GetActiveBlock() (sessiontemplate.Block, error) {
	if s.ended {
		return sessiontemplate.Block{}, ErrSessionEnded
	}
	return s.blocks[s.activeIndex], nil
}

func (s *SessionRun) ConfirmBlock() error {

	if s.ended {
		return ErrSessionEnded
	}

	if s.isLastBlock() {
		s.ended = true
		return nil
	}
	s.activeIndex++
	return nil
}

func (s *SessionRun) IsEnded() bool {
	return s.ended
}
