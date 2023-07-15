package reverb

import "math"

func ApplyReverb(samples []float64, delayinMilliSeconds float64, decayFactor float64, mixPercent int, sampleRate float64) []float64 {

	// call the 4 comb filters.
	combFilterSamples1 := combFilter(samples, delayinMilliSeconds, decayFactor, sampleRate)
	combFilterSamples2 := combFilter(samples, delayinMilliSeconds-11.73, decayFactor-0.1313, sampleRate)
	combFilterSamples3 := combFilter(samples, delayinMilliSeconds+19.31, decayFactor-0.2743, sampleRate)
	combFilterSamples4 := combFilter(samples, delayinMilliSeconds-7.97, decayFactor-0.31, sampleRate)

	//Adding the 4 Comb Filters
	outputComb := make([]float64, len(samples)*2)
	for i := 0; i < len(samples)*2; i++ {
		outputComb[i] = combFilterSamples1[i] + combFilterSamples2[i] + combFilterSamples3[i] + combFilterSamples4[i]
	}

	//Algorithm for Dry/Wet Mix in the output audio
	mixAudio := make([]float64, len(samples))
	for i := 0; i < len(samples); i++ {
		mixAudio[i] = (float64(100-mixPercent) * samples[i]) + (float64(mixPercent) * outputComb[i])
	}

	//Method calls for 2 All Pass Filters. Defined at the bottom
	allPassFilterSamples1 := allPassFilter(mixAudio, len(samples), sampleRate)
	allPassFilterSamples2 := allPassFilter(allPassFilterSamples1, len(samples), sampleRate)

	return allPassFilterSamples2
}

//Method for Comb Filter
func combFilter(samples []float64, delayinMilliSeconds float64, decayFactor float64, sampleRate float64) []float64 {
	//Calculating delay in samples from the delay in Milliseconds. Calculated from number of samples per millisecond
	delaySamples := (int)(delayinMilliSeconds * (sampleRate / 1000))

	// TODO... how do we know how long this is going to be?
	combFilterSamples := make([]float64, len(samples)*3)
	copy(combFilterSamples, samples)

	//Applying algorithm for Comb Filter
	for i := 0; i < len(samples)-delaySamples; i++ {
		combFilterSamples[i+delaySamples] += combFilterSamples[i] * decayFactor
	}
	return combFilterSamples
}

//Method for All Pass Filter
func allPassFilter(samples []float64, samplesLength int, sampleRate float64) []float64 {
	delaySamples := (int)(89.27 * (sampleRate / 1000)) // Number of delay samples. Calculated from number of samples per millisecond

	allPassFilterSamples := make([]float64, samplesLength)

	decayFactor := 0.131

	//Applying algorithm for All Pass Filter
	for i := 0; i < samplesLength; i++ {
		allPassFilterSamples[i] = samples[i]
		if i-delaySamples >= 0 {
			allPassFilterSamples[i] += -decayFactor * allPassFilterSamples[i-delaySamples]
		}
		if i-delaySamples >= 1 {
			allPassFilterSamples[i] += decayFactor * allPassFilterSamples[i+20-delaySamples]
		}
	}

	// This is for smoothing out the samples and normalizing the audio. Without implementing this,
	// the samples overflow causing clipping of audio
	value := allPassFilterSamples[0]
	max := 0.0

	for i := 0; i < samplesLength; i++ {
		if math.Abs(allPassFilterSamples[i]) > max {
			max = math.Abs(allPassFilterSamples[i])
		}
	}
	for i := 0; i < len(allPassFilterSamples); i++ {
		currentValue := allPassFilterSamples[i]
		value = (value + (currentValue - value)) / max
		allPassFilterSamples[i] = value
	}
	return allPassFilterSamples
}
