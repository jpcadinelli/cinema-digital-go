package model

type Paginacao struct {
	Conteudo     any `json:"conteudo"`
	TotalItens   int `json:"totalItens"`
	LimiteItens  int `json:"limiteItens"`
	PaginaAtual  int `json:"paginaAtual"`
	TotalPaginas int `json:"totalPaginas"`
}
