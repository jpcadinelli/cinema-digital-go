package service

import (
	"cinema_digital_go/api/global/enum"
	"cinema_digital_go/api/models"
)

func VerificaPermissaoUsuario(usuario models.UsuarioDTOResponse, permissoes ...string) bool {
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
