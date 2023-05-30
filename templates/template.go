package templates

import (
	"cubizy/model"
	"cubizy/plugins/awstorage"
	"cubizy/util"
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type PageData struct {
	AppName      string
	BaseDomin    string
	BaseUrl      string
	CanonicalUrl string
	Path         string
	ThemePath    string
	Account      *model.Account
	Address      *model.Address
	ThemeSetting map[string]ThemeSetting
}

type ThemeSetting struct {
	Title string      `json:"title"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

// IndexTemplate holds all public templates
var IndexTemplate map[uint]*template.Template

// UserTemplate holds only user templates
var UserTemplate *template.Template

func init() {
	IndexTemplate = make(map[uint]*template.Template)
}

// RunTemplateFile run template of website
func RunTemplateFile(w http.ResponseWriter, r *http.Request, SubDomain, path string) error {
	var err error
	var pageData PageData
	var templateToRun *template.Template

	pageData.AppName = util.Settings.AppName
	pageData.BaseDomin = util.Settings.Domain
	pageData.CanonicalUrl = util.Settings.Protocol + "://" + SubDomain + "." + util.Settings.Domain
	pageData.Path = path
	pageData.Account = model.GetAccountByDomin(SubDomain)

	if pageData.Account == nil {
		err = errors.New("no valid account")
		return err
	}

	pageData.ThemeSetting = make(map[string]ThemeSetting)

	if pageData.Account != nil && pageData.Account.ID > 0 {
		json.Unmarshal([]byte(pageData.Account.ThemeSettings), &pageData.ThemeSetting)
		if pageData.Account.WideLogo == "" {
			pageData.Account.WideLogo = pageData.Account.Logo
		}
		if pageData.Account.Banner == "" {
			pageData.Account.Banner = pageData.Account.WideLogo
		}
	}

	if SubDomain == "admin" || SubDomain == "seller" {
		if UserTemplate == nil || util.Settings.Domain == "localhost" {
			UserTemplate, err = template.ParseFiles("templates/userindex.html")
			if err != nil {
				util.Log("error in template.ParseFiles for application ", SubDomain, err.Error())
			}
		}
		if strings.Contains(r.RequestURI, ".") {
			util.Log("User index got run from template", r.RequestURI)
		}
		templateToRun = UserTemplate
	} else {
		if _, ok := IndexTemplate[0]; !ok {
			var ThemePath = "templates/index.html"
			IndexTemplate[0], err = template.ParseFiles(ThemePath)
		}
		var ThemeID uint = 0
		var preprefix = ""
		if pageData.Account.AccountType != "admin" {
			preprefix = "account_" + strconv.Itoa(int(pageData.Account.ID))
		}
		if pageData.Account.ThemeID > 0 {
			ThemeID = pageData.Account.ThemeID
			if preprefix != "" {
				pageData.ThemePath = awstorage.GetAccessURL() + preprefix + "/themes/theme_" + strconv.FormatUint(uint64(pageData.Account.ThemeID), 10)
			} else {
				pageData.ThemePath = awstorage.GetAccessURL() + "themes/theme_" + strconv.FormatUint(uint64(pageData.Account.ThemeID), 10)
			}

			if _, ok := IndexTemplate[ThemeID]; !ok {
				var index_file string
				util.Log("reading theme file ", pageData.Account.ThemeID)
				var path = "themes/theme_" + strconv.FormatUint(uint64(pageData.Account.ThemeID), 10) + "/index.html"

				if preprefix != "" {
					path = preprefix + "/" + path
				}

				//util.Log("Reading file at ", path)
				index_file, err = awstorage.ReadFile(path)
				if err == nil && index_file != "" {
					_template := template.New(strconv.FormatUint(uint64(pageData.Account.ThemeID), 10))
					_template.Parse(index_file)
					IndexTemplate[ThemeID] = _template
					//util.Log("got theme :", index_file)
				} else if err != nil {
					util.Log(err)
				}
			} else {
				//util.Log("theme is allready present")
			}
		}
		if err != nil {
			util.Log("error in template.ParseFiles for website", SubDomain, err.Error())
		}
		if _, ok := IndexTemplate[ThemeID]; ok {
			templateToRun = IndexTemplate[ThemeID]
			//util.Log(templateToRun.DefinedTemplates())
		} else {
			templateToRun = IndexTemplate[0]
			util.Log("template not found, runing default")
		}
	}

	if err == nil {
		err = templateToRun.Execute(w, &pageData)
	} else {
		util.Log("RunTemplateFile", err.Error())
	}
	return err
}
