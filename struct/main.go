package main

import "github.com/davecgh/go-spew/spew"

type Sample struct {
	WebApp struct {
		Servlet []struct {
			ServletName  string `json:"servlet-name"`
			ServletClass string `json:"servlet-class"`
			InitParam    struct {
				ConfigGlossaryInstallationAt string `json:"configGlossary:installationAt"`
				ConfigGlossaryAdminEmail     string `json:"configGlossary:adminEmail"`
				ConfigGlossaryPoweredBy      string `json:"configGlossary:poweredBy"`
				ConfigGlossaryPoweredByIcon  string `json:"configGlossary:poweredByIcon"`
				ConfigGlossaryStaticPath     string `json:"configGlossary:staticPath"`
				TemplateProcessorClass       string `json:"templateProcessorClass"`
				TemplateLoaderClass          string `json:"templateLoaderClass"`
				TemplatePath                 string `json:"templatePath"`
				TemplateOverridePath         string `json:"templateOverridePath"`
				DefaultListTemplate          string `json:"defaultListTemplate"`
				DefaultFileTemplate          string `json:"defaultFileTemplate"`
				UseJSP                       bool   `json:"useJSP"`
				JspListTemplate              string `json:"jspListTemplate"`
				JspFileTemplate              string `json:"jspFileTemplate"`
				CachePackageTagsTrack        int64  `json:"cachePackageTagsTrack"`
				CachePackageTagsStore        int64  `json:"cachePackageTagsStore"`
				CachePackageTagsRefresh      int64  `json:"cachePackageTagsRefresh"`
				CacheTemplatesTrack          int64  `json:"cacheTemplatesTrack"`
				CacheTemplatesStore          int64  `json:"cacheTemplatesStore"`
				CacheTemplatesRefresh        int64  `json:"cacheTemplatesRefresh"`
				CachePagesTrack              int64  `json:"cachePagesTrack"`
				CachePagesStore              int64  `json:"cachePagesStore"`
				CachePagesRefresh            int64  `json:"cachePagesRefresh"`
				CachePagesDirtyRead          int64  `json:"cachePagesDirtyRead"`
				SearchEngineListTemplate     string `json:"searchEngineListTemplate"`
				SearchEngineFileTemplate     string `json:"searchEngineFileTemplate"`
				SearchEngineRobotsDB         string `json:"searchEngineRobotsDb"`
				UseDataStore                 bool   `json:"useDataStore"`
				DataStoreClass               string `json:"dataStoreClass"`
				RedirectionClass             string `json:"redirectionClass"`
				DataStoreName                string `json:"dataStoreName"`
				DataStoreDriver              string `json:"dataStoreDriver"`
				DataStoreURL                 string `json:"dataStoreUrl"`
				DataStoreUser                string `json:"dataStoreUser"`
				DataStorePassword            string `json:"dataStorePassword"`
				DataStoreTestQuery           string `json:"dataStoreTestQuery"`
				DataStoreLogFile             string `json:"dataStoreLogFile"`
				DataStoreInitConns           int64  `json:"dataStoreInitConns"`
				DataStoreMaxConns            int64  `json:"dataStoreMaxConns"`
				DataStoreConnUsageLimit      int64  `json:"dataStoreConnUsageLimit"`
				DataStoreLogLevel            string `json:"dataStoreLogLevel"`
				MaxURLLength                 int64  `json:"maxUrlLength"`
			} `json:"init-param"`
		} `json:"servlet"`
		ServletMapping struct {
			CofaxCDS    string `json:"cofaxCDS"`
			CofaxEmail  string `json:"cofaxEmail"`
			CofaxAdmin  string `json:"cofaxAdmin"`
			FileServlet string `json:"fileServlet"`
			CofaxTools  string `json:"cofaxTools"`
		} `json:"servlet-mapping"`
		Taglib struct {
			TaglibURI      string `json:"taglib-uri"`
			TaglibLocation string `json:"taglib-location"`
		} `json:"taglib"`
	} `json:"web-app"`
}

const jsonStr = `{"web-app": {
  "servlet": [
    {
      "servlet-name": "cofaxCDS",
      "servlet-class": "org.cofax.cds.CDSServlet",
      "init-param": {
        "configGlossary:installationAt": "Philadelphia, PA",
        "configGlossary:adminEmail": "ksm@pobox.com",
        "configGlossary:poweredBy": "Cofax",
        "configGlossary:poweredByIcon": "/images/cofax.gif",
        "configGlossary:staticPath": "/content/static",
        "templateProcessorClass": "org.cofax.WysiwygTemplate",
        "templateLoaderClass": "org.cofax.FilesTemplateLoader",
        "templatePath": "templates",
        "templateOverridePath": "",
        "defaultListTemplate": "listTemplate.htm",
        "defaultFileTemplate": "articleTemplate.htm",
        "useJSP": false,
        "jspListTemplate": "listTemplate.jsp",
        "jspFileTemplate": "articleTemplate.jsp",
        "cachePackageTagsTrack": 200,
        "cachePackageTagsStore": 200,
        "cachePackageTagsRefresh": 60,
        "cacheTemplatesTrack": 100,
        "cacheTemplatesStore": 50,
        "cacheTemplatesRefresh": 15,
        "cachePagesTrack": 200,
        "cachePagesStore": 100,
        "cachePagesRefresh": 10,
        "cachePagesDirtyRead": 10,
        "searchEngineListTemplate": "forSearchEnginesList.htm",
        "searchEngineFileTemplate": "forSearchEngines.htm",
        "searchEngineRobotsDb": "WEB-INF/robots.db",
        "useDataStore": true,
        "dataStoreClass": "org.cofax.SqlDataStore",
        "redirectionClass": "org.cofax.SqlRedirection",
        "dataStoreName": "cofax",
        "dataStoreDriver": "com.microsoft.jdbc.sqlserver.SQLServerDriver",
        "dataStoreUrl": "jdbc:microsoft:sqlserver://LOCALHOST:1433;DatabaseName=goon",
        "dataStoreUser": "sa",
        "dataStorePassword": "dataStoreTestQuery",
        "dataStoreTestQuery": "SET NOCOUNT ON;select test='test';",
        "dataStoreLogFile": "/usr/local/tomcat/logs/datastore.log",
        "dataStoreInitConns": 10,
        "dataStoreMaxConns": 100,
        "dataStoreConnUsageLimit": 100,
        "dataStoreLogLevel": "debug",
        "maxUrlLength": 500}},
    {
      "servlet-name": "cofaxEmail",
      "servlet-class": "org.cofax.cds.EmailServlet",
      "init-param": {
      "mailHost": "mail1",
      "mailHostOverride": "mail2"}},
    {
      "servlet-name": "cofaxAdmin",
      "servlet-class": "org.cofax.cds.AdminServlet"},

    {
      "servlet-name": "fileServlet",
      "servlet-class": "org.cofax.cds.FileServlet"},
    {
      "servlet-name": "cofaxTools",
      "servlet-class": "org.cofax.cms.CofaxToolsServlet",
      "init-param": {
        "templatePath": "toolstemplates/",
        "log": 1,
        "logLocation": "/usr/local/tomcat/logs/CofaxTools.log",
        "logMaxSize": "",
        "dataLog": 1,
        "dataLogLocation": "/usr/local/tomcat/logs/dataLog.log",
        "dataLogMaxSize": "",
        "removePageCache": "/content/admin/remove?cache=pages&id=",
        "removeTemplateCache": "/content/admin/remove?cache=templates&id=",
        "fileTransferFolder": "/usr/local/tomcat/webapps/content/fileTransferFolder",
        "lookInContext": 1,
        "adminGroupID": 4,
        "betaServer": true}}],
  "servlet-mapping": {
    "cofaxCDS": "/",
    "cofaxEmail": "/cofaxutil/aemail/*",
    "cofaxAdmin": "/admin/*",
    "fileServlet": "/static/*",
    "cofaxTools": "/tools/*"},

  "taglib": {
    "taglib-uri": "cofax.tld",
    "taglib-location": "/WEB-INF/tlds/cofax.tld"}}}`

func main() {
	s := &Sample{}

	spew.Dump(s.WebApp.Servlet[0])
}
