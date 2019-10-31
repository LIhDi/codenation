package models


type Answer struct {
    NumeroCasas int `json:"numero_casas"`
	Token string `json:"token"`
	Cifrado string `json:"cifrado"`
	Decifrado string `json:"decifrado"`
	ResumoCriptografico string `json:"resumo_criptografico"`
}