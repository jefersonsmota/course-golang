package interfaces

type IVerificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta IVerificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}
