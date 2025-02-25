package resource

import (
	"cinema_digital_go/api/app/genero/model"
	"cinema_digital_go/api/app/genero/repository"
	"cinema_digital_go/api/pkg/global/enum"
	"cinema_digital_go/api/pkg/global/erros"
	"cinema_digital_go/api/pkg/middleware"
	service2 "cinema_digital_go/api/pkg/utils"
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeneroResource struct {
	repo repository.GeneroRepository
}

func NewGeneroResource(repo repository.GeneroRepository) *GeneroResource {
	return &GeneroResource{repo}
}

func (h *GeneroResource) Create(c *gin.Context) {
	usuarioLogado, err := service2.GetUsuarioLogado(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioCriar) {
		c.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var genero model.Genero
	if err := c.ShouldBindJSON(&genero); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	genero.ID = uuid.New()

	if err := h.repo.Create(&genero); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar o gênero"})
		return
	}
	c.JSON(http.StatusCreated, genero)
}

func (h *GeneroResource) GetAll(c *gin.Context) {
	usuarioLogado, err := service2.GetUsuarioLogado(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioVisualizar) {
		c.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	genero, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar gêneros"})
		return
	}
	c.JSON(http.StatusOK, genero)
}

func (h *GeneroResource) Update(c *gin.Context) {
	id := c.Param("id")

	idUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var genero model.Genero
	if err := c.ShouldBindJSON(&genero); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	genero.ID = idUUID

	if err := h.repo.Update(&genero); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o gênero"})
		return
	}

	c.JSON(http.StatusOK, genero)
}

func (h *GeneroResource) Delete(c *gin.Context) {
	id := c.Param("id")

	idUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	usuarioLogado, err := service2.GetUsuarioLogado(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioDeletar) {
		c.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	if err := h.repo.Delete(idUUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar o gênero"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
