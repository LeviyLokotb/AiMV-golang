package main

import (
	"fft/fx"
	"fft/graphdraw"
	"math"
	"math/cmplx"
)

var count = 0

// FFT без complex128 (возвращает отдельно real и imag части)
func FFT(signal []float64) []complex128 {
	count++
	N := len(signal)
	if N == 1 {
		result := []complex128{complex(fx.Furje(signal[0]), 0)}
		return result
	}

	// Рекурсивное разделение на чётные/нечётные
	even := make([]float64, N/2)
	odd := make([]float64, N/2)
	for i := 0; i < N/2; i++ {
		even[i] = signal[2*i]
		odd[i] = signal[2*i+1]
	}

	evenFFT := FFT(even)
	oddFFT := FFT(odd)

	// Комбинирование результатов
	spectrum := make([]complex128, N)
	for k := 0; k < N/2; k++ {
		// Twiddle factor: W_N^k = e^(-2πik/N)
		angle := -2 * math.Pi * float64(k) / float64(N)
		w := cmplx.Exp(complex(0, angle))

		// X[k] = evenFFT[k] + w * oddFFT[k]
		spectrum[k] = evenFFT[k] + w*oddFFT[k]
		// X[k + N/2] = evenFFT[k] - w * oddFFT[k]
		spectrum[k+N/2] = evenFFT[k] - w*oddFFT[k]
	}
	return spectrum
}

// Вычисляем амплитуду чтобы построить график
func Amplitude(val []complex128) []float64 {
	amp := make([]float64, len(val))
	for i := 0; i < len(val); i++ {
		amp[i] = cmplx.Abs(val[i])
	}
	return amp
}

func main() {
	var inputX []float64
	var inputY []float64
	for i := fx.A; i < fx.B; i += fx.H {
		inputX = append(inputX, i)
		inputY = append(inputY, fx.F(i))
	}

	base := graphdraw.DrawFunction(inputX, inputY, "Base")
	graphdraw.Save(base, "BASE")

	recontruct := Amplitude(FFT(inputX))
	fft := graphdraw.DrawFunction(inputX, recontruct, "FFT")
	graphdraw.Save(fft, "FFT")

}
