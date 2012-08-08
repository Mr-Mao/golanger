package web

type Config struct {
	SupportSession            bool                   "SupportSession"
	AutoGenerateHtml          bool                   "AutoGenerateHtml"
	AutoGenerateHtmlCycleTime int64                  "AutoGenerateHtmlCycleTime"
	AutoJumpToHtml            bool                   "AutoJumpToHtml"
	Debug                     bool                   "Debug"
	StaticDirectory           string                 "StaticDirectory"
	ThemeDirectory            string                 "ThemeDirectory"
	Theme                     string                 "Theme"
	StaticCssDirectory        string                 "StaticCssDirectory"
	StaticJsDirectory         string                 "StaticJsDirectory"
	StaticImgDirectory        string                 "StaticImgDirectory"
	HtmlDirectory             string                 "HtmlDirectory"
	TemplateDirectory         string                 "TemplateDirectory"
	TemplateGlobalDirectory   string                 "TemplateGlobalDirectory"
	TemplateGlobalFile        string                 "TemplateGlobalFile"
	TemporaryDirectory        string                 "TemporaryDirectory"
	UploadDirectory           string                 "UploadDirectory"
	IndexDirectory            string                 "IndexDirectory"
	IndexPage                 string                 "IndexPage"
	SiteRoot                  string                 "SiteRoot"
	Environment               map[string]string      "Environment"
	Database                  map[string]string      "Database"
	M                         map[string]interface{} "Custom"
}
