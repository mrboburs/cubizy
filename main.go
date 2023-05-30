package main

import (
	"context"
	"crypto/tls"
	"cubizy/api"
	"cubizy/model"
	"cubizy/myws"
	"cubizy/templates"
	"cubizy/util"
	"cubizy/vue"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	util.Log("Welcome")
	http.HandleFunc("/", serveAppPage)
	http.HandleFunc("/vue/", vue.Handler)
	http.HandleFunc("/api/", api.RestApihandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/default"))))
	http.HandleFunc("/ws/", myws.Handler)

	if util.Settings.Domain != "" {
		if strings.Contains(util.Settings.Protocol, "https") {

			util.Log("runing https server :", util.Settings.Domain)
			certManager := autocert.Manager{
				Prompt:     autocert.AcceptTOS,
				Cache:      autocert.DirCache("certs"),
				HostPolicy: myHostPolicy(),
			}

			server := &http.Server{
				Addr: ":443",
				TLSConfig: &tls.Config{
					GetCertificate: certManager.GetCertificate,
				},
			}
			go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
			server.ListenAndServeTLS("", "")
		} else {
			util.Log("runing http server :", util.Settings.Domain)
			util.Log(http.ListenAndServe(":80", nil)) // , http.HandlerFunc(handle)
		}
	} else {
		util.Log("No Domin to start")
	}
}

func myHostPolicy() autocert.HostPolicy {
	return func(_ context.Context, host string) error {
		if host == util.Settings.Domain {
			return nil
		}

		SubDomain := strings.ReplaceAll(host, util.Settings.Domain, "")
		SubDomain = strings.Trim(SubDomain, ".")
		SubDomain = strings.TrimSpace(SubDomain)
		if strings.EqualFold(SubDomain, "admin") || strings.EqualFold(SubDomain, "seller") {
			return nil
		}
		account := model.GetAccountByDomin(SubDomain)
		if account != nil {
			return nil
		}
		return fmt.Errorf("invalid sub domain %q ", SubDomain)
	}
}

func serveAppPage(w http.ResponseWriter, r *http.Request) {

	SubDomain := strings.ReplaceAll(r.Host, util.Settings.Domain, "")
	SubDomain = strings.Trim(SubDomain, ".")
	SubDomain = strings.TrimSpace(SubDomain)
	//util.Log("Host : ", r.Host)
	//util.Log("SubDomain :", SubDomain)
	//util.Log("SubDomain check :", SubDomain == "admin")

	path := r.URL.Path

	if SubDomain == "admin" || SubDomain == "seller" {
		if path == "/" {
			path = "/userindex.html"
		}
		filePath := "static/" + SubDomain + path
		//util.Log(filePath)
		if !util.FileExists(filePath) {
			//util.Log(filePath, " not exist")
			filePath = "static/default" + path
		}
		if util.FileExists(filePath) {
			http.ServeFile(w, r, filePath)
		} else {
			if strings.Contains(path, ".css.map") || strings.Contains(path, ".js.map") {
				fmt.Fprint(w, "")
			} else {
				templates.RunTemplateFile(w, r, SubDomain, path)
			}
		}
	} else {
		if path == "/" {
			path = "/index.html"
		}
		filePath := "static/default" + path
		//util.Log(filePath)
		if util.FileExists(filePath) {
			http.ServeFile(w, r, filePath)
		} else {
			templates.RunTemplateFile(w, r, SubDomain, path)
		}
	}
}
