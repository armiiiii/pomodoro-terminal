package sessionrun

import (
	sessiontemplate "pomodoro/internal/session-template"
	"reflect"
	"testing"
	"testing/synctest"
	"time"
)

func TestSkipBlock_SavedToStats(t *testing.T) {
	session, sessionTemplate := initializeSession(t)

	session.SkipBlock()

	got := session.GetStats()
	want := SessionStatistics{
		skippedBlocks: []sessiontemplate.Block{sessionTemplate.GetBlocks()[0]},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SkipBlock() = %v, want %v", got, want)
	}
}

func TestPause_TotalTimeSavedToStats(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		session, sessionTemplate := initializeSession(t)
		duration := sessionTemplate.GetBlocks()[0].GetDuration()

		session.Start()
		time.Sleep(duration / 2)
		session.Pause()
		time.Sleep(duration / 2)
		session.Resume()

		totalPauseTime := session.GetStats().totalPauseTime

		if totalPauseTime != duration-duration/2 {
			t.Errorf("totalPauseTime = %v, wanted %v", totalPauseTime, duration-duration/2)
		}
	})
}
