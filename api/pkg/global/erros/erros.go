package erros

import (
	"errors"
	"fmt"
)

var (
	ErrTokenInexistente = errors.New("token inexistente, acesso não autorizado")
	ErrTokenInvalido    = errors.New("token inválido, acesso não autorizado")

	ErrUsuarioNaoTemPermissao = fmt.Errorf("usuário não tem permissão")
)

var (
	ErrUsuarioNaoEncontrado              = fmt.Errorf("usuário não encontrado")
	ErrCredenciaisInvalidas              = fmt.Errorf("credenciais inválidas do usuário")
	ErrNaoPodeMudadarDadosDeOutroUsuario = fmt.Errorf("seu usuário não tem permissão de mudar os dados de outro usuário")

	ErrPermissaoNaoEncontrada = fmt.Errorf("permissão não encontrada")

	ErrProdutoNaoEncontrado = fmt.Errorf("produto não encontrado")

	ErrEstoqueNaoEncontrado   = fmt.Errorf("cadastro no estoque não encontrado")
	ErrQuantidadeInsuficiente = "quantidade insuficiente, a disponibilidade do produto %v no estoque é de %v"

	ErrEnderecoNaoEncontrado = fmt.Errorf("endereço não encontrado")

	ErrClienteNaoEncontrado = fmt.Errorf("cliente não encontrado")

	ErrVendaNaoEncontrada = fmt.Errorf("venda não encontrada")
)
