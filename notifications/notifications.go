package notifications

import (
	"log/slog"

	"github.com/gen2brain/beeep"
)

func Notify(text string) {
	err := beeep.Notify(
		"Pomodoro",
		text,
		"assets/tomato.png",
	)

	if err != nil {
		slog.Warn(
			"Failed to send notification",
			"text", text,
		)
	}
}

func PlaySound() {
	notes := []float64{523.25, 659.25, 783.99}

	for _, note := range notes {
		if err := beeep.Beep(note, 200); err != nil {
			slog.Warn("Failed to play sound")
		}
	}
}
