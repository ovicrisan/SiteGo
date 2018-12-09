// (c) 2018 Ovi Crisan
// http://ovi.crisan.ca
// https://github.com/ovicrisan/SiteGo

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"SiteGo/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	e     *echo.Echo
	home  []byte
	about []byte
)

func page(c echo.Context) error {
	buffer := new(bytes.Buffer)
	switch id := c.Param("id"); strings.ToLower(id) {
	case "":
		template.Header("Home", buffer)
		if len(home) == 0 {
			template.Home(template.Site, buffer)
		} else {
			buffer.Write(home)
		}
	case "about":
		template.Header("About", buffer)
		if len(about) == 0 {
			template.About(buffer)
		} else {
			buffer.Write(about)
		}
	case "contact":
		template.Header("Contact", buffer)
		template.Contact(c, buffer)
	default:
		template.Header("Error", buffer)
		template.Error(buffer)
	}
	template.Footer(buffer)
	return c.HTMLBlob(http.StatusOK, buffer.Bytes())
}

func main() {
	var (
		port string = "8000"
		log  bool   = false
	)
	// Read envinronment variables
	if tmp, ok := os.LookupEnv("SITE_NAME"); ok {
		template.Site = tmp
	}
	if tmp, ok := os.LookupEnv("SITE_PORT"); ok {
		port = tmp
	}

	if tmp, ok := os.LookupEnv("SITE_LOGO"); ok {
		template.Logo = tmp != "0"
	}

	if tmp, ok := os.LookupEnv("SITE_LOG"); ok {
		log = tmp != "0"
	}

	// Setup command line flags
	flag.StringVar(&template.Site, "s", template.Site, "site name")
	flag.StringVar(&port, "p", port, "port")
	flag.BoolVar(&template.Logo, "i", template.Logo, "logo image")
	flag.BoolVar(&log, "l", log, "log to console")
	flag.Bool("b", false, "hide Echo banner")
	flag.Bool("version", false, "show version")

	// Read parameter flags
	flag.Parse()

	// Show version and exit
	if f := flag.CommandLine.Lookup("version"); f != nil && f.Value.String() == "true" {
		fmt.Println("sitego version 1.0")
		os.Exit(0)
	}

	// Get hostname and local path
	template.Hostname, _ = os.Hostname()
	exe, _ := os.Executable()
	path := filepath.Dir(exe)

	// Read custom pages, if any (fail gracefully)
	home, _ = ioutil.ReadFile(path + "/home.html")
	about, _ = ioutil.ReadFile(path + "/about.html")

	e = echo.New()
	e.Static("/static", path+"/static")

	if f := flag.CommandLine.Lookup("l"); log || (f != nil && f.Value.String() == "true") {
		e.Use(middleware.Logger())
	}

	if f := flag.CommandLine.Lookup("b"); f != nil && f.Value.String() == "true" {
		e.HideBanner = true
		fmt.Println("Starting on port ", port)
	}

	e.GET("/", page)
	e.Any("/:id", page)

	e.Logger.Fatal(e.Start(":" + port))
}
