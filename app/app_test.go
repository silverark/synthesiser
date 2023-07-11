package app

import (
	"synth/app/output"
	"testing"
)

func TestNewAMinorChord(t *testing.T) {
	c := NewAMinorChord(Bar, 1.0)
	c2 := NewAMinorChord2(Bar, 1.0)

	both := c.Add(c2)

	player := output.NewOtoPlayer(SampleRate) // Write to player
	defer player.Close()
	both.Write(player, 0.5)
}
