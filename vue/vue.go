package vue

import (
	"cubizy/util"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
)

const (
	compiledFolder = "compiled"
	codeFolder     = "code"
)

var myminify *minify.M

func init() {
	util.Log("vue module initiating")

	if util.Settings.Minify {
		myminify = minify.New()
		myminify.AddFunc("text/javascript", js.Minify)
		myminify.AddFunc("text/html", html.Minify)
	}

	WalkFileChanges(filepath.Join("vue", codeFolder))
	defer util.Log("vue module initiated")
}

// Handler will handle component request and conforms providing letest build of componnts
func Handler(w http.ResponseWriter, r *http.Request) {
	var err error

	SubDomain := strings.ReplaceAll(r.Host, util.Settings.Domain, "")
	SubDomain = strings.Trim(SubDomain, ".")
	SubDomain = strings.TrimSpace(SubDomain)

	componentPath := r.URL.Path[len("/vue/"):]
	if util.Settings.Domain == "localhost" {
		componentCodePath := strings.Replace(componentPath, ".js", ".vue", 1)
		if SubDomain == "admin" || SubDomain == "seller" {
			subdominCodePath := filepath.Join("vue", codeFolder, SubDomain, componentCodePath)
			if util.FileExists(subdominCodePath) {
				componentCodePath = subdominCodePath
			} else {
				componentCodePath = filepath.Join("vue", codeFolder, "default", componentCodePath)
			}
		} else {
			componentCodePath = filepath.Join("vue", codeFolder, "default", componentCodePath)
		}
		err = checkVueVersion(componentCodePath)
		if err != nil {
			util.Log("Vue Handler Code:", componentCodePath, err)
		}
	}

	if err == nil {
		if SubDomain == "admin" || SubDomain == "seller" {
			subdominPath := filepath.Join("vue", compiledFolder, SubDomain, componentPath)
			if util.FileExists(subdominPath) {
				componentPath = subdominPath
			} else {
				componentPath = filepath.Join("vue", compiledFolder, "default", componentPath)
			}
		} else {
			componentPath = filepath.Join("vue", compiledFolder, "default", componentPath)
		}
		if util.FileExists(componentPath) {
			http.ServeFile(w, r, componentPath)
		} else {
			util.Log("Vue Handler:", componentPath, "Not present")
		}
	}
}

// WalkFileChanges will Walk folder to check changes and update vue js files
func WalkFileChanges(folder string) {
	err := filepath.Walk(folder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				destinationFolder := strings.Replace(path, codeFolder, compiledFolder, 1)
				if _, err := os.Stat(destinationFolder); os.IsNotExist(err) {
					os.MkdirAll(destinationFolder, os.ModeDir)
				}
			} else if strings.Contains(info.Name(), ".vue") {
				err = checkVueVersion(path)
			}
			if err != nil {
				util.Log("WalkFileChanges:", path, err)
			}
			return nil
		})
	if err != nil {
		util.Log(err)
	}
}

func checkVueVersion(path string) error {
	var err error
	sourceFilePath := path
	destinationFilePath := path

	destinationFilePath = strings.Replace(destinationFilePath, codeFolder, compiledFolder, 1)
	destinationFilePath = strings.Replace(destinationFilePath, ".vue", ".js", 1)

	vuefileStat, err := os.Stat(sourceFilePath)
	if err == nil {
		jsfileStat, _err := os.Stat(destinationFilePath)
		err = _err
		if err != nil || vuefileStat.ModTime().Unix() > jsfileStat.ModTime().Unix() {
			err = generateJs(sourceFilePath, destinationFilePath)
		}
	}
	return err
}

func generateJs(sourceFilePath, destinationFilePath string) error {
	var err error
	var body []byte
	body, err = ioutil.ReadFile(sourceFilePath)
	if err == nil {
		jscode := string(body)
		if strings.Contains(jscode, "<script>") {
			finalCode := GetStringInBetween(jscode, "<script>", "</script>")
			if util.Settings.Minify {
				finalCode, err = myminify.String("text/javascript", finalCode)
			}
			if err == nil {
				if strings.Contains(jscode, "<template>") {
					template := GetStringInBetween(jscode, "<template>", "</template>")
					if util.Settings.Minify {
						template, err = myminify.String("text/html", template)
					}
					if err == nil {
						if strings.Contains("{{{template}}}", "{{{template}}}") {
							finalCode = strings.Replace(finalCode, "{{{template}}}", template, 1)
						} else {
							util.Log("Do you Forgot to add '{{{template}}}' in ", sourceFilePath)
						}
					}
				}
				jscode = finalCode
			}
		}
		if err == nil {

			destinationFolder := filepath.Dir(destinationFilePath)
			if _, err := os.Stat(destinationFolder); os.IsNotExist(err) {
				os.MkdirAll(destinationFolder, os.ModeDir)
			}

			err = ioutil.WriteFile(destinationFilePath, []byte(jscode), 0644)
			//if err == nil {
			//util.Log("Writen component at : ", destinationFilePath)
			//}
		} else {
			util.Log(sourceFilePath, " component have errors : ", err)
		}
	}
	return err
}

// GetStringInBetween Returns empty string if no start string found
func GetStringInBetween(str string, start string, end string) (result string) {
	s := 0
	if len(start) > 0 {
		s = strings.Index(str, start)
		if s == -1 {
			return
		}
		s += len(start)
	}
	e := len(str) - len(end)
	if len(end) > 0 {
		e = strings.LastIndex(str[s:], end) + s //.Index(str[s:], end) + s
		if e == -1 {
			return
		}
	}
	return str[s:e]
}
