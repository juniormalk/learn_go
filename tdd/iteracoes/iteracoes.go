package iteracao

func Repetir(caractere string) string {
	var repeticoes string
	for i := 0; i < 5; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
