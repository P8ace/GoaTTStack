package main

import (
	"log/slog"
	"net/http"

	"github.com/P8ace/GoaTTStack/microservices/common/components"
	"github.com/P8ace/GoaTTStack/microservices/frontend/controllers"
	"github.com/a-h/templ"
)

type app struct {
	addr string
}

func NewApp(addr string) *app {
	return &app{
		addr: addr,
	}
}

func (self *app) registerControllers() http.Handler {
	mux := http.NewServeMux()

	controller := controllers.NewController()
	mux.HandleFunc("/health", controller.HealthController)

	navigation := components.NavigationElement()
	mainParagraph := components.ParagraphElement("In the Main page")
	coursesParagraph := components.ParagraphElement("In the Courses page")
	pricesParagraph := components.ParagraphElement("In the Prices page")
	aboutParagraph := components.ParagraphElement("In the About page")

	mainPage := components.SimpleHtmlSkeleton(navigation, mainParagraph)
	mux.Handle("/", templ.Handler(mainPage))

	coursesPage := components.SimpleHtmlSkeleton(navigation, coursesParagraph)
	mux.Handle("/courses", templ.Handler(coursesPage))

	pricesPage := components.SimpleHtmlSkeleton(navigation, pricesParagraph)
	mux.Handle("/pricing", templ.Handler(pricesPage))

	AboutPage := components.SimpleHtmlSkeleton(navigation, aboutParagraph)
	mux.Handle("/about", templ.Handler(AboutPage))

	return mux
}

func (self *app) getServer(mux http.Handler) http.Server {
	return http.Server{
		Addr:    self.addr,
		Handler: mux,
	}
}

func main() {
	app := app{
		addr: ":8080",
	}

	mux := app.registerControllers()
	server := app.getServer(mux)

	slog.Info("Starting http server", "Addr", app.addr)
	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Error while starting the server", "Error", err.Error(), "Addr", app.addr)
	}
}
