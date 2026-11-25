package sessionrun

import (
	sessiontemplate "pomodoro/internal/session-template"
	"reflect"
	"testing"
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
