package sessionrun

import (
	"errors"
	sessiontemplate "pomodoro/internal/session-template"
	"time"

	"github.com/google/uuid"
)

var (
	ErrSessionEnded = errors.New("session already ended")
)

type SessionStatistics struct {
	skippedBlocks  []sessiontemplate.Block
	totalPauseTime time.Duration
}

type SessionRun struct {
	id             uuid.UUID
	blocks         []sessiontemplate.Block
	blockStartTime time.Time
	pauseStartTime time.Time
	activeIndex    int
	ended          bool
	stats          SessionStatistics
}

func CreateSessionRun(sessionTemplate sessiontemplate.SessionTemplate) (SessionRun, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return SessionRun{}, err
	}

	var sessionRun SessionRun = SessionRun{
		id:          id,
		activeIndex: 0,
		blocks:      sessionTemplate.GetBlocks(),
	}
	return sessionRun, nil
}

func (s *SessionRun) Start() {
	s.blockStartTime = time.Now()
}

func (s *SessionRun) Pause() {
	s.pauseStartTime = time.Now()
}

func (s *SessionRun) Resume() {
	s.stats.totalPauseTime += time.Since(s.pauseStartTime)
	s.blockStartTime = s.blockStartTime.Add(time.Since(s.pauseStartTime))
	s.pauseStartTime = time.Time{}
}

func (s *SessionRun) GetRemainingTime() (time.Duration, error) {
	block, err := s.GetActiveBlock()
	if err != nil {
		return 0, err
	}

	now := time.Now()

	if !s.pauseStartTime.IsZero() {
		now = s.pauseStartTime
	}

	return block.GetDuration() - now.Sub(s.blockStartTime), nil
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
