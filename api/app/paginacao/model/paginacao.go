package model

type Paginacao struct {
	Conteudo     any `json:"conteudo"`
	TotalItens   int `json:"totalItens"`
	Limite       int `json:"limite"`
	PaginaAtual  int `json:"paginaAtual"`
	TotalPaginas int `json:"totalPaginas"`
}
