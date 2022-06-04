package {{ .Package }}

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuender/oils/base"
	"gorm.io/gorm"
)

// {{ .Name }} TODO.
type {{ .Name }} struct{
	db *gorm.DB
}

// New{{ .Name }} TODO.
func New{{ .Name }}(database *gorm.DB) *{{ .Name }} {
	return &{{ .Name }}{
		db: database,
	}
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
	err := p.db.Transaction(func(tx *gorm.DB) error {
		// TODO call service
		return nil
	})

	ctx.JSON(http.StatusOK, gin.H{"success": err==nil, "error":err})
}

// DELETE TODO.
func (p *{{ .Name }}) DELETE(ctx *gin.Context) {
	id := base.Must1(base.ParseInteger[int](ctx.Param("id")))
	err := p.db.Transaction(func(tx *gorm.DB) error {
		// TODO call service
		return nil
	})

	ctx.JSON(http.StatusOK, gin.H{"success": err==nil, "id": id, "error":err})
}

// PUT TODO.
func (p *{{ .Name }}) PUT(ctx *gin.Context) {
	id := base.Must1(base.ParseInteger[int](ctx.Param("id")))
	method := ctx.Query("method")
	err := p.db.Transaction(func(tx *gorm.DB) error {
		// TODO call service
		return nil
	})

	ctx.JSON(http.StatusOK, gin.H{"success": err==nil, "id": id, "method": method, "error": err})
}
