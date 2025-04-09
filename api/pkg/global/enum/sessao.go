package enum

const (
	SessaoDisponivel = 1
	SessaoCancelada  = 2
)

var mapperSessao = map[int]string{
	SessaoDisponivel: "Disponível",
	SessaoCancelada:  "Cancelada",
}

func GetSessaoByEnum(enum int) string {
	return mapperSessao[enum]
}
