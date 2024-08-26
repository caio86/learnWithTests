package romannumerals

import "testing"

func BenchmarkConverToArabic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertToArabic("MDCLXVI")
	}
}

func BenchmarkConvertToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertToRoman(3984)
	}
}
