package iteration

func Repeat(character string, repeatedCount int) string {
	if repeatedCount == 0 {
		repeatedCount = 5
	}
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += character
	}
	return repeated
}
