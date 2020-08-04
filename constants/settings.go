package constants

const (
	SitePagesTemplate      = "templates/"
	SiteAdminPagesTemplate = "templates/webmaster/pages/"
	SiteTemplateExtension  = ".gohtml"
)

// Site Constants
const (
	// Domain. It is used in session management.
	Domain = "localhost"
	// PORT. It is used in session management.
	PORT = ":8080"
	// SiteUrl Web URL of the website
	SiteUrl = Domain + PORT + "/"
	// SiteProtocol contains the protocol supported by the domain.
	SiteProtocol = "http://"
	// SiteName. It is used in emails.
	SiteName = "mywebsite.com"
	// DeveloperSiteUrl contains Web Developer Site. It is used in footer and mailers.
	DeveloperSiteUrl = "https://www.abc.com/"
	// Web Developer Company
	DeveloperCompany = "Abc developers"
	// standard image references
	// TODO: Change shortLogo when its ready
	// Logo file name used in main site
	MainLogo = "logo-8.png"
	// Logo icon file name used in main site
	ShortLogo = "logo-8.png"
)


const AdminFolderPath = "webmaster/"
const AdminFolder = "/"+AdminFolderPath

const (
	AdminLogin = AdminFolder + "login"
	Dist       = AdminFolderPath + "dist/"
	Plugins    = AdminFolderPath + "plugins/"
)
