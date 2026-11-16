package main

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/signal"
	"runtime"
	"time"

	"log/slog"

	"github.com/gen2brain/beeep"

	"pomodoro/config"
	"pomodoro/logger"
	"pomodoro/notifications"
	"pomodoro/ui"
)

var randWorkEmojis = []string{"🛠️", "🎯", "⏰", "🧑🏻‍💻", "🌶️"}
var randRelaxEmojis = []string{"˙✧˖°☕ ༘ ⋆｡", "⊹ ࣪ ﹏𓊝﹏𓂁﹏⊹ ࣪ ˖", "*ੈ✩‧₊˚༺☆༻*ੈ✩‧₊˚"}

func countdown(ctx context.Context, label string, d time.Duration) bool {
	const barWidth = 30
	deadline := time.Now().Add(d)
	ticker := time.NewTicker(100 * time.Millisecond)

	defer ticker.Stop()

	for {
		remaining := time.Until(deadline)
		if remaining <= 0 {
			fmt.Printf("\r%s %s 00:00\n", label, ui.RenderProgress(1, barWidth))
			return true
		}

		elapsed := d - remaining
		progress := float64(elapsed) / float64(d)

		m := uint32(remaining.Minutes())
		s := uint32(remaining.Seconds()) % 60

		fmt.Printf("\r%s %s %02d:%02d\033[K", label, ui.RenderProgress(progress, barWidth), m, s)

		select {
		case <-ticker.C:
		case <-ctx.Done():
			slog.Warn("Timer was cancelled")
			fmt.Printf("\r%s %s cancelled\033[K\n", label, ui.RenderProgress(progress, barWidth))
			return false
		}
	}
}

func main() {

	logger, closeLogger, err := logger.New(true)

	if err != nil {
		log.Fatal(err)
	}

	defer closeLogger()

	slog.SetDefault(logger.With(
		"version", "0.1.0",
		"os", runtime.GOOS,
		"arch", runtime.GOARCH,
	))

	beeep.AppName = "Pomodoro"

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	slog.Info("Application started")

	slog.Debug("Parsing config")
	cfg, err := config.ParseArgs()

	if err != nil {
		slog.Error("Config contains errors")
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	slog.Info("Timer started",
		"work duration", cfg.WorkTime,
		"relax duration", cfg.RelaxTime,
		"iterations", cfg.Iterations)

	for i := 1; i <= int(cfg.Iterations); i++ {
		label := fmt.Sprintf("%s Work [%d/%d]", randWorkEmojis[rand.IntN(len(randWorkEmojis))], i, cfg.Iterations)
		if !countdown(
			ctx,
			label,
			time.Duration(cfg.WorkTime)*time.Minute) {
			return
		}

		notifications.Notify("ХВАТИТ РАБОТАТЬ, ВРЕМЯ ОТДОХНУТЬ")
		notifications.PlaySound()

		if i < int(cfg.Iterations) {
			label := fmt.Sprintf("%s Relax [%d/%d]", randRelaxEmojis[rand.IntN(len(randRelaxEmojis))], i, cfg.Iterations)
			if !countdown(ctx, label, time.Duration(cfg.RelaxTime)*time.Minute) {
				return
			}
		}
	}

	slog.Info("Timer has ended successfully")
	notifications.Notify("Таймер истёк. Можно чиллить!")
	notifications.PlaySound()
}
