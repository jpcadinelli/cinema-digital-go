package model

type Paginacao struct {
	TotalItens   int `json:"totalItens"`
	Limite       int `json:"limite"`
	PaginaAtual  int `json:"paginaAtual"`
	TotalPaginas int `json:"totalPaginas"`
}
