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
	ErrUsuarioNaoEncontrado                 = fmt.Errorf("usuário não encontrado")
	ErrCredenciaisInvalidas                 = fmt.Errorf("credenciais inválidas do usuário")
	ErrNaoPodeMudadarDadosDeOutroUsuario    = fmt.Errorf("seu usuário não tem permissão de mudar os dados de outro usuário")
	ErrNaoPodeVisualizarDadosDeOutroUsuario = fmt.Errorf("seu usuário não tem permissão de visualizar os dados de outro usuário")

	ErrPermissaoNaoEncontrada          = fmt.Errorf("permissão não encontrada")
	ErrGrupoDePermissoesNaoEncontradas = fmt.Errorf("grupo de permissões não encontradas")

	ErrFilmeNaoEncontrado = fmt.Errorf("filme não encontrado")

	ErrGeneroNaoEncontrado = fmt.Errorf("gênero não encontrado")
	ErrGeneroInvalido      = fmt.Errorf("nome de gênero inválido")

	ErrSalaNaoEncontrada = fmt.Errorf("sala não encontrada")
	ErrSalaInvalida      = fmt.Errorf("nome de sala inválido")
)
