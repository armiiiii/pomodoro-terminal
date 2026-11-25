package sessionrun

import (
	"errors"
	sessiontemplate "pomodoro/internal/session-template"
)

var (
	ErrSessionEnded = errors.New("session already ended")
)

type SessionStatistics struct {
	skippedBlocks []sessiontemplate.Block
}

type SessionRun struct {
	stats       SessionStatistics
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

func (s *SessionRun) ConfirmBlock() error {
	if s.ended {
		return ErrSessionEnded
	}

	s.advanceBlocks()

	return nil
}

func (s *SessionRun) SkipBlock() error {
	block, err := s.GetActiveBlock()
	if err != nil {
		return ErrSessionEnded
	}

	s.stats.skippedBlocks = append(s.stats.skippedBlocks, block)

	s.advanceBlocks()

	return nil
}

func (s *SessionRun) GetActiveBlock() (sessiontemplate.Block, error) {
	if s.ended {
		return sessiontemplate.Block{}, ErrSessionEnded
	}
	return s.blocks[s.activeIndex], nil
}

func (s *SessionRun) advanceBlocks() {
	if s.isLastBlock() {
		s.ended = true
		return
	}
	s.activeIndex++
}

func (s *SessionRun) isLastBlock() bool {
	return s.activeIndex >= len(s.blocks)-1
}

func (s *SessionRun) IsEnded() bool {
	return s.ended
}

func (s *SessionRun) GetStats() SessionStatistics {
	return s.stats
}
