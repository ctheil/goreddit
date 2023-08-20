package web

import (
	"net/http"
	"text/template"

	"github.com/ctheil/goreddit"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
)

func NewHandler(store goreddit.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)
	h.Route("/threads", func(r chi.Router) {
		r.Get("/", h.ThreadsList())
		r.Get("/new", h.ThreadsCreate())
		r.Post("/", h.ThreadsStore())
		r.Post("/{id}/delete", h.ThreadsDelete())
	})
	return h
}

type Handler struct {
	*chi.Mux

	store goreddit.Store
}

const threadsListHTML = `
<h1>Threads</h1>
<dl>
{{range .Threads}}
	<dt>
		<strong>{{.Title}}</strong>
		<form action="/threads/{{.ID}}/delete" method="POST">
			<button type="submit">delete</button>
		</form>
	</dt>	
	<dd>
		{{.Description}}
	</dd>
{{end}}
</dl>
<a href="/threads/new">Create Thread</a>
`

func (h *Handler) ThreadsList() http.HandlerFunc {
	type data struct {
		Threads []goreddit.Thread
	}

	tmpl := template.Must(template.New("").Parse(threadsListHTML))
	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Threads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, data{Threads: tt}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

const threadCreateHTML = `
<h1>New Thread</h1>
<form action="/threads" method="POST">
	<table>
		<tr>
			<td>Title</td>
			<td><input type="text" name="title" /></td>
		</tr>
		<tr>
			<td>Description</td>
			<td><input type="text" name="description" /></td>
		</tr>
	</table>
	<button type="submit">Create Thread</button>
</form>

`

func (h *Handler) ThreadsCreate() http.HandlerFunc {
	tmpl := template.Must(template.New("").Parse(threadCreateHTML))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}
func (h *Handler) ThreadsStore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := r.FormValue("title")
		d := r.FormValue("description")
		if err := h.store.CreateThread(&goreddit.Thread{
			ID:          uuid.New(),
			Title:       t,
			Description: d,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}
func (h *Handler) ThreadsDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// get threadToDelete ID here via mux...?
		if err := h.store.DeleteThread(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}
