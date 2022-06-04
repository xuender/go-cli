package {{ .Package }}

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuender/oils/base"
)

// {{ .Name }} TODO.
type {{ .Name }} struct{}

// New{{ .Name }} TODO.
func New{{ .Name }}() *{{ .Name }} {
	return &{{ .Name }}{}
}

// Group config router.
func (p *{{ .Name }}) Group(group *gin.RouterGroup) {
	group.GET(":id", p.GET)
	group.POST("", p.POST)
	group.DELETE(":id", p.DELETE)
	group.PUT(":id", p.PUT)
}

// GET TODO.
func (p *{{ .Name }}) GET(ctx *gin.Context) {
	id := base.Must1(base.ParseInteger[int](ctx.Param("id")))
	// TODO find
	ctx.JSON(http.StatusOK, gin.H{"success": false, "id": id})
}

// POST TODO.
func (p *{{ .Name }}) POST(ctx *gin.Context) {
	// TODO ctx.Bind()
	ctx.JSON(http.StatusOK, "success")
}

// DELETE TODO.
func (p *{{ .Name }}) DELETE(ctx *gin.Context) {
	id := base.Must1(base.ParseInteger[int](ctx.Param("id")))
	// TODO delete
	ctx.JSON(http.StatusOK, gin.H{"success": false, "id": id})
}

// PUT TODO.
func (p *{{ .Name }}) PUT(ctx *gin.Context) {
	id := base.Must1(base.ParseInteger[int](ctx.Param("id")))
	method := ctx.Query("method")
	// TODO method
	ctx.JSON(http.StatusOK, gin.H{"success": false, "id": id, "method": method})
}
