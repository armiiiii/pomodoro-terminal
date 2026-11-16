package ui

import (
	"fmt"
	"strings"
)

func RenderProgress(progress float64, width uint8) string {
	if progress < 0 {
		progress = 0
	}
	if progress > 1 {
		progress = 1
	}

	filled := uint8(progress * float64(width))
	bar := strings.Repeat("█", int(filled)) + strings.Repeat("░", int(width-filled))
	return fmt.Sprintf("[%s] %3.0f%%", bar, progress*100)
}
