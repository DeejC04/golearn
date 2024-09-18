package iteration

func Repeat(character string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return 
}

// Above is diff syntax than was on the site but I think it's a bit nicer instead
// of explicitly definiing the 'repeated' variable above the loop
// Syntax on site is below (commented)

// func Repeat(character string) string {
// 	var repeated string
// 	for i := 0; i < repeatCount; i++ {
// 		repeated += character
// 	}
// 	return repeated
// }