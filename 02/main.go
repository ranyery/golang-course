package main

// var b bool
// var c int
// O escopo das variáveis abaixo são globais, pois podem ser acessadas em qualquer função dentro do package main
var (
	b bool    = true      // default is false
	c int     = 0         // default is 0
	d string  = "Ranyery" // default is ""
	e float64 = 1.5       // default is +0.000000e+000
)

func main() {
	// b = true

	// O escopa da variável abaixo é local, visto apenas dentro da função
	const nome string = "Ranyery Coutinho"
	// Outra forma de criar a variável acima seria:
	// nome := "Ranyery Coutinho"

	println(b)
	println(c)
	println(d)
	println(e)
}
