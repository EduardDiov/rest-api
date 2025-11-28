package router

import (
	"net/http"

	"diov.local/krasivo/internal/http-server/handlers"
	positions "diov.local/krasivo/internal/http-server/handlers/positions"
	"diov.local/krasivo/internal/util"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Start() chi.Mux {
	// Router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Define handlers
	r.Get("/", handlers.IndexPage)

	staticFilePath := util.GetExecutablePath() + "/public/"
	r.Handle("/*", http.FileServer(http.Dir(staticFilePath)))

	r.Get("/search", handlers.Search)

	r.Route("/currency", func(r chi.Router) {
		r.Get("/{symbol}", handlers.Currency)
		// r.Post("/", CreateArticle)
		// r.Route("/{articleID}", func(r chi.Router) {
		//     r.Use(ArticleCtx)
		//     r.Get("/", GetArticle) // GET /articles/1234
		//     r.Put("/", UpdateArticle)    // PUT /articles/1234
		//     r.Delete("/", DeleteArticle) // DELETE /articles/1234
		//     r.Get("/edit", EditArticle) // GET /articles/1234/edit
		// })
	})

	r.Route("/positions", func(r chi.Router) {
		r.Get("/", positions.IndexPage)
		// r.Post("/", CreateArticle)
		// r.Route("/{articleID}", func(r chi.Router) {
		//     r.Use(ArticleCtx)
		//     r.Get("/", GetArticle) // GET /articles/1234
		//     r.Put("/", UpdateArticle)    // PUT /articles/1234
		//     r.Delete("/", DeleteArticle) // DELETE /articles/1234
		//     r.Get("/edit", EditArticle) // GET /articles/1234/edit
		// })
	})

	return *r
}
