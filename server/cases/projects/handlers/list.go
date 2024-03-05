package handlers

import (
	"github.com/Quero-Freela/system/server/utils"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {

	projects := []string{"Project 1", "Project 2", "Project 3"}

	utils.Success(w, r, projects)
}
