package models

import (
	"github.com/programmer-richa/admin_template_slicing/constants"
	"html/template"
	"log"
	"net/http"
)

// Tpl stores the reference to the pointer of the template parser.
// It is used for processing template files used in this project.
var Tpl *template.Template

// struct PageHelper contains data to be presented on the template, url of the page and template file name.
type PageHelper struct {
	Title             string
	FileName          string
	MainLogoFileName  string
	ShortLogoFileName string
	SiteURL           string
	SiteProtocol      string
	SiteName          string
	DeveloperURL      string
	DeveloperCompany  string
	Private           bool
	Css              []string
	Js               []string
	FormNames []string
}

func New() *PageHelper {
	return &PageHelper{}
}

func (p *PageHelper) WithTitle(title string) *PageHelper {
	p.Title = title
	return p
}

func (p *PageHelper) WithFileName(fileName string) *PageHelper {
	p.FileName = fileName
	return p
}

func (p *PageHelper) WithMainLogoFileName(mainLogoFileName string) *PageHelper {
	p.MainLogoFileName = mainLogoFileName
	return p
}

func (p *PageHelper) WithShortLogoFileName(shortLogoFileName string) *PageHelper {
	p.ShortLogoFileName = shortLogoFileName
	return p
}

func (p *PageHelper) WithSiteURL(siteURL string) *PageHelper {
	p.SiteURL = siteURL
	return p
}

func (p *PageHelper) WithSiteProtocol(siteProtocol string) *PageHelper {
	p.SiteProtocol = siteProtocol
	return p
}

func (p *PageHelper) WithSiteName(siteName string) *PageHelper {
	p.SiteName = siteName
	return p
}

func (p *PageHelper) WithDeveloperURL(developerURL string) *PageHelper {
	p.DeveloperURL = developerURL
	return p
}

func (p *PageHelper) WithDeveloperCompany(developerCompany string) *PageHelper {
	p.DeveloperCompany = developerCompany
	return p
}
func (p *PageHelper) WithCss(css string) *PageHelper {
	if p.Css==nil {
		p.Css= []string{}
	}
	err,path:=GetFileUrlWithModifiedTime(css)
	if err == nil {
		p.Css = append(p.Css, path)
	}
	return p
}

func (p *PageHelper) WithJs(js string) *PageHelper {
	if p.Js==nil {
		p.Js= []string{}
	}
	err,path:=GetFileUrlWithModifiedTime(js)
	if err == nil {
		p.Js = append(p.Js, path)
	}
	return p
}

func (p *PageHelper) IsPrivate(private bool) *PageHelper {
	p.Private = private
	return p
}

// Page struct is created to access all features of pagehelper.PageHelper
// It is required as new methods can be attached to the local 'type' (data type) only.
type Page struct {
	*PageHelper
}

// NewPage returns pointer to Page
func NewPage(p *PageHelper) *Page {
	return &Page{p}
}

// SetDefaultData associate logo path, site urls.
// This is refined to pass default page data.
func (p *Page) SetDefaultData() {
	p.MainLogoFileName = constants.MainLogo
	p.ShortLogoFileName = constants.ShortLogo
	p.DeveloperCompany = constants.DeveloperCompany
	p.DeveloperURL = constants.DeveloperSiteUrl
	p.SiteURL = constants.SiteUrl
	p.SiteProtocol = constants.SiteProtocol
	p.SiteName = constants.SiteName
}

// ProcessPage is request handlerFunc for the page
func (p *Page) ProcessPage(w http.ResponseWriter, r *http.Request) {
	p.SetDefaultData()
	err := Tpl.ExecuteTemplate(w, p.FileName, p)
	if err != nil {
		log.Fatalln(err)
	}
}
