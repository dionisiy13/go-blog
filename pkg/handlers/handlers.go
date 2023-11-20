package handlers

import (
	"github.com/dionisiy13/go-web/pkg/config"
	"github.com/dionisiy13/go-web/pkg/models"
	"github.com/dionisiy13/go-web/pkg/render"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {
	a := config.GetAppConfig()

	str := render.RenderTemplate("home.page.tmpl", &models.TemplateData{
		Data: map[string]interface{}{
			"var1": a.Session.Get(c.Request().Context(), "23"),
			"var2": "KEK AZAZA",
		},
	})

	return c.HTML(http.StatusOK, str)
}

func About(c echo.Context) error {
	a := config.GetAppConfig()
	a.Session.Put(c.Request().Context(), "23", "131331")
	str := render.RenderTemplate("about.page.tmpl", &models.TemplateData{})

	return c.HTML(http.StatusOK, str)
}

func DoNothing(c echo.Context) error {
	return c.HTML(http.StatusOK, "")
}
