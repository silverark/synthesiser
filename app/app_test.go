package app

import (
	"synth/app/chart"
	"synth/app/envelope"
	"synth/app/notes"
	"synth/app/output"
	"synth/app/signal"
	"testing"
)

func TestNewChord(t *testing.T) {

	e5Minor := NewChord(Bar/2, 1.0, notes.E5, Minor)
	g4Major := NewChord(Bar/2, 1.0, notes.G4, Major)

	song := e5Minor.Add(g4Major)

	player := output.NewOtoPlayer(SampleRate) // Write to player
	defer player.Close()
	song.Write(player, 0.5)
}

func TestNewSineWaveWithEnvelope(t *testing.T) {

	c5 := signal.NewSineWave(notes.Freq(notes.C5), Bar, 0.5, SampleRate).Generate()
	c5Sine := signal.NewSineWaveWithEnvelope(notes.Freq(notes.C5), Bar, 0.5, SampleRate, envelope.ShapeRound).Generate()
	//c5CoSine := signal.NewSineWaveWithEnvelope(notes.Freq(notes.C5), Bar, 0.5, SampleRate, envelope.ShapeCosine).Generate()
	c5Linear := signal.NewSineWaveWithEnvelope(notes.Freq(notes.C5), Bar, 0.5, SampleRate, envelope.ShapeLinear).Generate()
	c5ShapeOscillate := signal.NewSineWaveWithEnvelope(notes.Freq(notes.C5), Bar, 0.5, SampleRate, envelope.ShapeOscillate).Generate()
	c5ShapeOscillate20 := signal.NewSineWaveWithEnvelope(notes.Freq(notes.C5), Bar, 0.5, SampleRate, envelope.ShapeOscillate8).Generate()
	c5ShapeOscillate30 := signal.NewSineWaveWithEnvelope(notes.Freq(notes.C5), Bar, 0.5, SampleRate, envelope.ShapeOscillate16).Generate()
	c5ShapeOscillate32 := signal.NewSineWaveWithEnvelope(notes.Freq(notes.C5), Bar, 0.5, SampleRate, envelope.ShapeOscillate32).Generate()
	c5drop := signal.NewSineWaveWithEnvelope(notes.Freq(notes.C5), Bar/2, 0.5, SampleRate, envelope.ShapeDrop).Generate()

	chart.ChartIt("c5-drop", c5drop.Data)

	song := c5.Add(c5Sine).Add(c5Sine).Add(c5Linear).Add(c5ShapeOscillate).Add(c5ShapeOscillate20).Add(c5ShapeOscillate30).Add(c5ShapeOscillate32).Add(c5drop)

	//player := output.NewOtoPlayer(SampleRate) // Write to player
	//defer player.Close()
	//song.Write(player, 0.5)

	output := output.NewWavWriter("test.wav", SampleRate, true) // Write to WAV file
	defer output.Close()
	song.Write(output, 0.5)
}
