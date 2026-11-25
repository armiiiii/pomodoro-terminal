package sessionrun

import (
	"errors"
	"testing"
)

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

func TestSkipBlock_ActivatesNextBlock(t *testing.T) {
	session, sessionTemplate := initializeSession(t)

	session.SkipBlock()

	got, err := session.GetActiveBlock()
	want := sessionTemplate.GetBlocks()[1]
	if err != nil {
		t.Fatalf("GetActiveBlock() = %v. returned error to quickly", err)
	}

	if got != want {
		t.Errorf("SkipBlock() = %v, want %v", got, want)
	}
}

func TestSkipBlock_LastBlockEndsSession(t *testing.T) {
	session, sessionTemplate := initializeSession(t)

	skipAllBlocks(t, &session, &sessionTemplate)

	if !session.IsEnded() {
		t.Errorf("SkipBlock() expected to end the session")
	}
}

func TestSkipBlock_RejectsSkippingAfterLastBlock(t *testing.T) {
	session, sessionTemplate := initializeSession(t)
	skipAllBlocks(t, &session, &sessionTemplate)

	err := session.SkipBlock()

	if !errors.Is(err, ErrSessionEnded) {
		t.Errorf("SkipBlock() = %v, expected %v", err, ErrSessionEnded)
	}
}
