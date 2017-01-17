package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	port := ":" + "9999"
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK,
			"<a href=\"http://localhost"+
				port+
				"/location\">link</a>")
	})
	e.GET("/javascript", func(c echo.Context) error {
		return c.HTML(http.StatusOK,
			"<html>"+
				"<head>"+
				"<script type=\"text/javascript\" language=\"javascript\">"+
				"function redirect(){"+
				"location.replace(\"http://localhost"+
				port+
				"/destination\");"+
				"}"+
				"window.onload = function(){"+
				"redirect();"+
				"}"+
				"</script>"+
				"</head>"+
				"<body>"+
				"<a href=\"#\" onclick=\"redirect()\">link</a></body></html>")
	})
	e.GET("/location", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost"+port+"/destination")
	})
	e.GET("/destination", func(c echo.Context) error {
		refererURL := c.Request().Header.Get("Referer")
		return c.String(http.StatusOK, refererURL)
	})
	e.Logger.Fatal(e.Start(port))
}
