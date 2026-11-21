package sessionrun

import (
	"errors"
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

func TestCreateSessionRun_ActiveBlockIsFirstInBlocks(t *testing.T) {

	session, sessionTemplate := initializeSession(t)

	want := sessionTemplate.GetBlocks()[0]
	got, err := session.GetActiveBlock()

	if err != nil {
		t.Fatalf("GetActiveBlock() = %v. Expected to not return error", err)
	}

	if got != want {
		t.Errorf("First active block is %v, want %v", got, want)
	}
}

func TestConfirmBlock_ChangesAfterConfirm(t *testing.T) {
	session, sessionTemplate := initializeSession(t)

	err := session.ConfirmBlock()
	if err != nil {
		t.Fatalf("ConfirmBlock() = %v. returned error to quickly", err)
	}

	want := sessionTemplate.GetBlocks()[1]
	got, err := session.GetActiveBlock()

	if err != nil {
		t.Fatalf("GetActiveBlock() = %v. Expected to not return error", err)
	}

	if got != want {
		t.Errorf("Active block is %v, want %v", got, want)
	}
}

func TestConfirmBlock_IsEndedAfterLastBlock(t *testing.T) {
	session, sessionTemplate := initializeSession(t)
	confirmAllBlocks(t, &session, &sessionTemplate)

	if !session.IsEnded() {
		t.Errorf("IsEnded() expected to be true, got false")
	}
}

func TestGetActiveBlock_AfterSessionEnded(t *testing.T) {
	session, sessionTemplate := initializeSession(t)
	confirmAllBlocks(t, &session, &sessionTemplate)

	_, err := session.GetActiveBlock()

	if !errors.Is(err, ErrSessionEnded) {
		t.Errorf("GetActiveBlock() = %v, want %v", err, ErrSessionEnded)
	}

}

func TestConfirmBlock_RejectsAfterLastBlock(t *testing.T) {
	session, sessionTemplate := initializeSession(t)
	confirmAllBlocks(t, &session, &sessionTemplate)

	err := session.ConfirmBlock()

	if !errors.Is(err, ErrSessionEnded) {
		t.Errorf("ConfirmBlock() = %v, want %v", err, ErrSessionEnded)
	}
}
