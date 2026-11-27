package sessionrun

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestGetRemainingTime(t *testing.T) {
	tests := []struct {
		name     string
		scenario func(s *SessionRun, d time.Duration)
		want     func(d time.Duration) time.Duration
	}{
		{
			name: "after elapsed time",
			scenario: func(s *SessionRun, d time.Duration) {
				time.Sleep(d / 2)
			},
			want: func(d time.Duration) time.Duration {
				return d - d/2
			},
		},
		{
			name: "during pause",
			scenario: func(s *SessionRun, d time.Duration) {
				time.Sleep(d / 2)
				s.Pause()
				time.Sleep(d / 2)
			},
			want: func(d time.Duration) time.Duration {
				return d - d/2
			},
		},
		{
			name: "after resume",
			scenario: func(s *SessionRun, d time.Duration) {
				time.Sleep(d / 2)
				s.Pause()
				time.Sleep(d / 2)
				s.Resume()
				time.Sleep(d / 4)
			},
			want: func(d time.Duration) time.Duration {
				return d - d/2 - d/4
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				session, template := initializeSession(t)
				duration := template.GetBlocks()[0].GetDuration()

				session.Start()

				tt.scenario(&session, duration)

				got, err := session.GetRemainingTime()
				if err != nil {
					t.Fatal(err)
				}

				want := tt.want(duration)

				if got != want {
					t.Errorf("got %v, want %v", got, want)
				}
			})
		})
	}
}

func TestConfirmBlock_ResetsRemainingTime(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		session, template := initializeSession(t)
		duration := template.GetBlocks()[0].GetDuration()

		session.Start()

		time.Sleep(duration)

		// User confirms the block
		if err := session.ConfirmBlock(); err != nil {
			t.Fatal(err)
		}

		block, err := session.GetActiveBlock()
		if err != nil {
			t.Fatal(err)
		}

		expectedBlock := template.GetBlocks()[1]
		if block != expectedBlock {
			t.Errorf("block wasn't confirmed, block %v, expectedBlock %v", block, expectedBlock)
		}

		got, err := session.GetRemainingTime()
		if err != nil {
			t.Fatal(err)
		}
		want := expectedBlock.GetDuration()
		if got != want {
			t.Errorf("GetRemainingTime()=%v, want %v", got, want)
		}
	})
}
