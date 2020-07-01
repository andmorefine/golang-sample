package iteration

const count int = 3

// Repeat return string
func Repeat(chara string) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated = repeated + chara
	}

	return repeated
}
