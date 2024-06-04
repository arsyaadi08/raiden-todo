package controllers

import (
	"testraiden/internal/models"

	"github.com/sev-2/raiden"
)

type TodoController struct {
	raiden.ControllerBase
	Http string `path:"/todos" type:"rest"`
	Model models.Todo
}