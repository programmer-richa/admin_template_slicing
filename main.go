package main

import (
	"fmt"
	"github.com/programmer-richa/admin_template_slicing/constants"
	"github.com/programmer-richa/admin_template_slicing/models"
	"github.com/programmer-richa/admin_template_slicing/routes"
	"html/template"
	"net/http"
)

// Initializer block initialises templates used in the project.
func init() {
	models.Tpl = template.Must(template.New("").ParseGlob(constants.SiteAdminPagesTemplate + "*" + constants.SiteTemplateExtension))
	//models.Tpl.ParseGlob(constants.SiteAdminPagesTemplate + "*." + constants.SiteTemplateExtension)
}



func main() {
	http.Handle("/"+constants.Plugins, http.FileServer(http.Dir(".")))
	http.Handle("/"+constants.Dist, http.FileServer(http.Dir(".")))
	// Routes using defaultServeMux
	routes.SetPages()
	// Listener
	err := http.ListenAndServe(constants.PORT, nil)
	if err != nil {
		fmt.Printf("Could not start application on port 8080. %v", err)
	}

}
