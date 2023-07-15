package wav

import (
	"fmt"
	"github.com/youpy/go-wav"
	"math"
	"os"
)

func ReadWavFile(input_file string, number_of_samples uint32) ([]float64, uint16, uint32, uint16) {

	if number_of_samples == 0 {
		number_of_samples = math.MaxInt32
	}

	blockAlign := 2
	file, err := os.Open(input_file)
	if err != nil {
		panic(err)
	}

	reader := wav.NewReader(file)
	wavformat, err_rd := reader.Format()
	if err_rd != nil {
		panic(err_rd)
	}

	if wavformat.AudioFormat != wav.AudioFormatPCM {
		panic("Audio format is invalid ")
	}

	if int(wavformat.BlockAlign) != blockAlign {
		fmt.Println("Block align is invalid ", wavformat.BlockAlign)
	}

	samples, err := reader.ReadSamples(number_of_samples) // must supply num samples w/o defaults to 2048
	//                                                    // just supply a HUGE number then actual num is returned
	wav_samples := make([]float64, 0)

	for _, curr_sample := range samples {
		wav_samples = append(wav_samples, reader.FloatValue(curr_sample, 0))
	}

	return wav_samples, wavformat.BitsPerSample, wavformat.SampleRate, wavformat.NumChannels
}
