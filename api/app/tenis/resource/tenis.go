package resource

import (
	"cinema_digital_go/api/app/tenis/model"
	"cinema_digital_go/api/pkg/global/enum"
	"cinema_digital_go/api/pkg/global/erros"
	"cinema_digital_go/api/pkg/middleware"
	"cinema_digital_go/api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func Criar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoTenisCriar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var t model.Tenis
	if err = ginctx.ShouldBindJSON(&t); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	t.CreatedAt = time.Now()
	t.Id = uuid.New()

	//if g.Nome == "" {
	//ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(erros.ErrGeneroInvalido, nil))
	//return
	//}

	//if err = repository.NewGeneroRepository(dbConection.DB).Create(&g); err != nil {
	//ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
	//return
	//}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, t))
}
