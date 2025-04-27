package model

type PaginacaoModel struct {
	TotalItens   int `json:"totalItens"`
	Limite       int `json:"limite"`
	PaginaAtual  int `json:"paginaAtual"`
	TotalPaginas int `json:"totalPaginas"`
}
