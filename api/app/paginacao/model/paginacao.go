package model

type PaginacaoModel struct {
	TotalItens   int `json:"total_itens"`
	Limite       int `json:"limite"`
	PaginaAtual  int `json:"pagina_atual"`
	TotalPaginas int `json:"total_paginas"`
}
