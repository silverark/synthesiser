package app

import (
	"fmt"
	"synth/app/envelope"
	"synth/app/notes"
	"synth/app/output"
	"synth/app/reverb"
	"synth/app/signal"
	"synth/app/wav"
	"testing"
)

func TestNewChord(t *testing.T) {

	e5Minor := NewChord(Bar/2, 1.0, notes.E5, Minor)
	g4Major := NewChord(Bar/2, 1.0, notes.G4, Major)

	song := e5Minor.Add(g4Major).Add(e5Minor).Add(g4Major)

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

	//chart.ChartIt("c5-drop", c5drop.Data)

	song := c5.Add(c5Sine).Add(c5Sine).Add(c5Linear).Add(c5ShapeOscillate).Add(c5ShapeOscillate20).Add(c5ShapeOscillate30).Add(c5ShapeOscillate32).Add(c5drop)

	//player := output.NewOtoPlayer(SampleRate) // Write to player
	//defer player.Close()
	//song.Write(player, 0.5)

	output := output.NewWavWriter("test.wav", SampleRate, true) // Write to WAV file
	defer output.Close()
	song.Write(output, 0.5)
}

func TestReverb(t *testing.T) {

	c5 := signal.NewSineWaveWithEnvelope(notes.Freq(notes.C5), Bar/4, 0.5, SampleRate, envelope.ShapeRound).Generate()
	e5 := signal.NewSineWaveWithEnvelope(notes.Freq(notes.E5), Bar/4, 0.5, SampleRate, envelope.ShapeRound).Generate()

	songClean := c5.Add(e5).Add(c5).Add(e5).Add(c5).Add(e5).Add(c5).Add(e5)
	songReverb := c5.Add(e5).Add(c5).Add(e5).Add(c5).Add(e5).Add(c5).Add(e5)

	songReverb.Data = reverb.ApplyReverb(songReverb.Data, 30.00, 0.25, 100, SampleRate)

	track := songClean.Add(songReverb)

	player := output.NewOtoPlayer(SampleRate) // Write to player
	defer player.Close()
	track.Write(player, 0.5)
}

func TestReverb_withAudio(t *testing.T) {
	input_audio := "tests/vocal.wav"
	audio_samples, bits_per_sample, input_audio_sample_rate, num_channels := wav.ReadWavFile(input_audio, 0)

	fmt.Println("num samples ", len(audio_samples)/int(num_channels))
	fmt.Println("bit depth   ", bits_per_sample)
	fmt.Println("sample rate ", input_audio_sample_rate)
	fmt.Println("num channels", num_channels)

	reverbAudio1 := reverb.ApplyReverb(audio_samples, 300.00, 0.45, 50, SampleRate)
	reverbAudio2 := reverb.ApplyReverb(audio_samples, 600.00, 0.45, 50, SampleRate)

	//player := output.NewOtoPlayer(SampleRate) // Write to player
	player := output.NewWavWriter("reverb.wav", SampleRate, true)
	defer player.Close()
	player.Write(audio_samples)
	player.Write(reverbAudio1)
	player.Write(reverbAudio2)
}
