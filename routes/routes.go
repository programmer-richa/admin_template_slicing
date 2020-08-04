package routes

import (
	"github.com/programmer-richa/admin_template_slicing/constants"
	"github.com/programmer-richa/admin_template_slicing/models"

	"net/http"
)

// RequestHandler links a url with its handler
func RequestHandler(url string, method string, handler func(http.ResponseWriter, *http.Request)) {
	if method == http.MethodGet || method == http.MethodPost {
		http.HandleFunc(url, handler)
	}
}

func SetPages() {
	pagehelper := models.New().
		WithTitle("Admin| Login").
		WithFileName("login" + constants.SiteTemplateExtension).
		WithCss(constants.Plugins+"fontawesome-free/css/all.min.css").
		WithCss(constants.Plugins+"icheck-bootstrap/icheck-bootstrap.min.css").
		WithCss(constants.Dist+"css/adminlte.min.css").
		WithJs(constants.Plugins+"jquery/jquery.min.js").
		WithJs(constants.Plugins+"bootstrap/js/bootstrap.bundle.min.js").
		WithJs(constants.Dist+"js/adminlte.min.js")
	login := models.NewPage(pagehelper)
	RequestHandler(constants.AdminLogin, http.MethodGet, login.ProcessPage)
}
