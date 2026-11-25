package sessionrun

import (
	sessiontemplate "pomodoro/internal/session-template"
	"pomodoro/internal/testfixtures"
	"testing"
)

func initializeSession(t *testing.T) (SessionRun, sessiontemplate.SessionTemplate) {
	t.Helper()

	sessionTemplate := testfixtures.NewSessionTemplate(t)
	session, err := CreateSessionRun(sessionTemplate)

	if err != nil {
		t.Fatalf("CreateSessionRun() returned unexpected error %v", err)
	}

	return session, sessionTemplate
}

func confirmAllBlocks(
	t *testing.T,
	session *SessionRun,
	sessionTemplate *sessiontemplate.SessionTemplate) {

	t.Helper()

	for range sessionTemplate.GetBlocks() {
		if err := session.ConfirmBlock(); err != nil {
			t.Fatalf("ConfirmBlock(): %v", err)
		}
	}
}

func skipAllBlocks(
	t *testing.T,
	session *SessionRun,
	sessionTemplate *sessiontemplate.SessionTemplate,
) {
	t.Helper()

	for range sessionTemplate.GetBlocks() {
		if err := session.SkipBlock(); err != nil {
			t.Fatalf("SkipBlock() returned error earlier than expected")
		}
	}
}
