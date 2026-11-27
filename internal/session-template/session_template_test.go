package sessiontemplate

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateBlock(t *testing.T) {
	tests := []struct {
		name    string
		title   string
		d       time.Duration
		icon    string
		want    Block
		wantErr bool
	}{
		{
			name:    "rejects no title",
			wantErr: true,
			title:   "",
			d:       time.Duration(15) * time.Second,
			icon:    "🎯",
		},
		{
			name:    "rejects zero duration time",
			wantErr: true,
			title:   "Learning Go",
			d:       time.Duration(0) * time.Second,
			icon:    "🎯",
		},
		{
			name:  "icon optional",
			title: "Learning Go",
			d:     time.Duration(15) * time.Second,
			want:  Block{title: "Learning Go", d: time.Duration(15) * time.Second},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateBlock(tt.title, tt.d, tt.icon)

			if (err != nil) != tt.wantErr {
				t.Fatalf("CreateBlock() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			if got != tt.want {
				t.Errorf("CreateBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSessionTemplate(t *testing.T) {

	blocks := []Block{
		{
			title: "Recalling previous material",
			d:     15 * time.Minute,
		},
		{
			title: "Learning new topic",
			d:     25 * time.Minute,
		},
		{
			title: "Relax",
			d:     5 * time.Minute,
		},
		{
			title: "Practics",
			d:     35 * time.Minute,
		},
		{
			title: "Q&A with no hints",
			d:     10 * time.Minute,
		},
	}

	tests := []struct {
		name    string
		title   string
		blocks  []Block
		want    SessionTemplate
		wantErr bool
	}{
		{
			name:    "rejects blank title",
			title:   "",
			blocks:  blocks,
			wantErr: true,
		},
		{
			name:    "rejects zero blocks",
			title:   "Interview Prep",
			wantErr: true,
		},
		{
			name:   "creates successfully",
			title:  "Interview prep",
			blocks: blocks,
			want: SessionTemplate{
				title:  "Interview prep",
				blocks: blocks,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSessionTemplate(tt.title, tt.blocks)

			if (err != nil) != tt.wantErr {
				t.Fatalf("CreateSessionTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			if got.ID() == uuid.Nil {
				t.Error("CreateSessionTemplate() returned a zero UUID")
			}

			if got.GetTitle() != tt.want.GetTitle() {
				t.Errorf("title = %q, want %q", got.GetTitle(), tt.want.GetTitle())
			}

			if !reflect.DeepEqual(got.GetBlocks(), tt.want.GetBlocks()) {
				t.Errorf("blocks = %v, want %v", got.GetBlocks(), tt.want.GetBlocks())
			}
		})
	}
}
