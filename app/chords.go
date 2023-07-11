package app

import (
	"math"
	"synth/app/notes"
	"synth/app/signal"
)

type ChordType int

const (
	Major ChordType = iota
	Minor
)

// NewChord creates a chord based on the base note and if you want Major or Minor
// Major Triads have 4 semitones then 3 semitones
// Minor triads have 3 semitones then 4 semitones
func NewChord(duration, volume float64, baseNote notes.Note, chordType ChordType) *signal.Signal {
	//for major chords lets set the steps to 4 and 3
	secondStep := 4
	thirdStep := 7
	//for minor chords lets set the steps to 3 and 4
	if chordType == Minor {
		secondStep = 3
		thirdStep = 7
	}

	first := notes.Freq(baseNote)
	second := notes.Freq(baseNote.Add(secondStep))
	third := notes.Freq(baseNote.Add(thirdStep))

	// TODO: If the NewSineWave accepted multiple frequencies it could generate the chord in one call.
	outputSignal := signal.NewSineWave(first, duration, volume, SampleRate).Generate()
	outputSignal = outputSignal.Superpose(
		signal.NewSineWave(second, duration, volume, SampleRate).Generate(),
		signal.NewSineWave(third, duration, volume, SampleRate).Generate(),
	)
	return outputSignal
}

func NewEMinorChord(duration, volume float64) *signal.Signal {
	eFreq := 440 * math.Pow(2, 7.0/12)  // E5
	gFreq := 440 * math.Pow(2, 10.0/12) // G5
	bFreq := 440 * math.Pow(2, 14.0/12) // B5

	outputSignal := signal.NewSineWave(eFreq, duration, volume, SampleRate).Generate()
	outputSignal = outputSignal.Superpose(
		signal.NewSineWave(gFreq, duration, volume, SampleRate).Generate(),
		signal.NewSineWave(bFreq, duration, volume, SampleRate).Generate(),
	)
	return outputSignal
}
