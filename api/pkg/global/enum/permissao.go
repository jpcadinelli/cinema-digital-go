package enum

var (
	ListaPermissoes = []string{
		PermissaoSistemaAdmin,

		PermissaoPermissaoCriar,
		PermissaoPermissaoVisualizar,
		PermissaoPermissaoListar,
		PermissaoPermissaoDropdown,
		PermissaoPermissaoAtualizar,
		PermissaoPermissaoDeletar,

		PermissaoUsuarioVisualizar,
		PermissaoUsuarioListar,
		PermissaoUsuarioDropdown,
		PermissaoUsuarioAtualizar,
		PermissaoUsuarioDeletar,

		PermissaoUsuarioAtribuirPermissao,
		PermissaoUsuarioRemoverPermissao,

		PermissaoFilmeCriar,
		PermissaoFilmeAtualizar,
		PermissaoFilmeDeletar,

		PermissaoGeneroCriar,
		PermissaoGeneroVisualizar,
		PermissaoGeneroListar,
		PermissaoGeneroDropdown,
		PermissaoGeneroAtualizar,
		PermissaoGeneroDeletar,

		PermissaoSalaCriar,
		PermissaoSalaVisualizar,
		PermissaoSalaListar,
		PermissaoSalaDropdown,
		PermissaoSalaAtualizar,
		PermissaoSalaDeletar,
	}
	GrupoN1Permissoes = []string{
		PermissaoUsuarioVisualizar,
		PermissaoUsuarioAtualizar,
	}
)

const (
	PermissaoSistemaAdmin = "SISTEMA_ADMIN"

	PermissaoPermissaoCriar      = "PERMISSAO_CRIAR"
	PermissaoPermissaoVisualizar = "PERMISSAO_VISUALIZAR"
	PermissaoPermissaoListar     = "PERMISSAO_LISTAR"
	PermissaoPermissaoDropdown   = "PERMISSAO_DROPDOWN"
	PermissaoPermissaoAtualizar  = "PERMISSAO_ATUALIZAR"
	PermissaoPermissaoDeletar    = "PERMISSAO_DELETAR"

	PermissaoUsuarioVisualizar = "USUARIO_VISUALIZAR"
	PermissaoUsuarioListar     = "USUARIO_LISTAR"
	PermissaoUsuarioDropdown   = "USUARIO_DROPDOWN"
	PermissaoUsuarioAtualizar  = "USUARIO_ATUALIZAR"
	PermissaoUsuarioDeletar    = "USUARIO_DELETAR"

	PermissaoUsuarioAtribuirPermissao = "USUARIO_ATRIBUIR_PERMISSAO"
	PermissaoUsuarioRemoverPermissao  = "USUARIO_REMOVER_PERMISSAO"

	PermissaoFilmeCriar     = "FILME_CRIAR"
	PermissaoFilmeAtualizar = "FILME_ATUALIZAR"
	PermissaoFilmeDeletar   = "FILME_DELETAR"

	PermissaoGeneroCriar      = "GENERO_CRIAR"
	PermissaoGeneroVisualizar = "GENERO_VISUALIZAR"
	PermissaoGeneroListar     = "GENERO_LISTAR"
	PermissaoGeneroDropdown   = "GENERO_DROPDOWN"
	PermissaoGeneroAtualizar  = "GENERO_ATUALIZAR"
	PermissaoGeneroDeletar    = "GENERO_DELETAR"

	PermissaoSalaCriar      = "SALA_CRIAR"
	PermissaoSalaVisualizar = "SALA_VISUALIZAR"
	PermissaoSalaListar     = "SALA_LISTAR"
	PermissaoSalaDropdown   = "SALA_DROPDOWN"
	PermissaoSalaAtualizar  = "SALA_ATUALIZAR"
	PermissaoSalaDeletar    = "SALA_DELETAR"
)
