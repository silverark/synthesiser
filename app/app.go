package app

import (
	"synth/app/notes"
	"synth/app/output"
)

const (
	SampleRate float64 = 44100.0
	Volume     float64 = 0.5 // 0.5 is the max
	OutputFile string  = "track.wav"
	Bpm                = 120            // beats per minute
	Bar                = 60.0 / Bpm * 4 // seconds
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	output := output.NewOtoPlayer(SampleRate) // Write to player
	// output := output.NewWavWriter(OutputFile, SampleRate, true) // Write to WAV file
	defer output.Close()

	highHat := NewHighHat(1)
	kickDrum := NewKickDrum(2)
	snare := NewSnare(1)

	kickDrumLoop := Sequencer(kickDrum, 4*Bar, Bpm, false)
	snareLoop := Sequencer(snare, 4*Bar, Bpm / 2, true)

	highHatLoopStraight := Sequencer(highHat, 4*Bar, Bpm, true)
	highHatLoopPolyrhythm := Sequencer(highHat, 4*Bar, Bpm / 3, true)
	highHatLoop := highHatLoopStraight.Superpose(highHatLoopPolyrhythm)

	chord1 := NewChord(Bar, 1, notes.A4, Minor)
	chord2 := NewChord(Bar, 1, notes.E5, Minor)
	chord3 := NewChord(Bar, 1, notes.D4, Minor)
	chord4 := NewChord(Bar, 1, notes.G4, Major)

	chordProgression := chord1.Add(chord2).Add(chord3).Add(chord4)

	intro := chordProgression
	verse := chordProgression.Superpose(kickDrumLoop, snareLoop)
	chorus := chordProgression.Superpose(kickDrumLoop, snareLoop, highHatLoop)
	bridge := verse.Superpose(snare)
	chorus2 := chorus.Superpose(highHat)
	outro := chordProgression.Superpose(kickDrumLoop)

	track := intro.Add(verse).Add(chorus).Add(bridge).Add(chorus2).Add(outro)

	track.Write(output, Volume)
}
