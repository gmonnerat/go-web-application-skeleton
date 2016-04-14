package main

import (
	"bytes"
	"github.com/labstack/echo"
	"io"
	"mime"
	"net/http"
	"path/filepath"
)

func home() echo.HandlerFunc {
	return func(c echo.Context) error {
		params := map[string]interface{}{"title": "Home"}
		str, _ := executeTemplate("templates/home.tmpl", params)
		c.HTML(http.StatusOK, str)
		return nil
	}
}

func detectContentType(name string) (t string) {
	if t = mime.TypeByExtension(filepath.Ext(name)); t == "" {
		t = echo.MIMEOctetStream
	}
	return
}

func serveFavicon() echo.HandlerFunc {
	return func(c echo.Context) error {
		path := "static/favicon.ico"
		data, err := Asset(path)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		fi, _ := AssetInfo(path)
		buffer := bytes.NewReader(data)
		rs := c.Response()
		rs.Header().Set(echo.HeaderContentType, detectContentType(fi.Name()))
		rs.Header().Set(echo.HeaderLastModified, fi.ModTime().UTC().Format(http.TimeFormat))
		rs.WriteHeader(http.StatusOK)
		_, err = io.Copy(rs, buffer)
		return err
	}
}

func serveFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		folder := c.Param("folder")
		file := c.Param("file")
		path := "static/" + folder + "/" + file
		data, err := Asset(path)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		fi, _ := AssetInfo(path)
		buffer := bytes.NewReader(data)
		rs := c.Response()
		rs.Header().Set(echo.HeaderContentType, detectContentType(fi.Name()))
		rs.Header().Set(echo.HeaderLastModified, fi.ModTime().UTC().Format(http.TimeFormat))
		rs.WriteHeader(http.StatusOK)
		_, err = io.Copy(rs, buffer)
		return err
	}
}
