package utils

import (
	"cinema_digital_go/api/app/usuario/model"
	"cinema_digital_go/api/pkg/global/enum"
)

func VerificaPermissaoUsuario(usuario model.UsuarioDTOResponse, permissoes ...string) bool {
	for _, permissao := range usuario.Permissoes {
		if permissao.Nome == enum.PermissaoSistemaAdmin {
			return true
		}

		for _, p := range permissoes {
			if p == permissao.Nome {
				return true
			}
		}
	}

	return false
}
