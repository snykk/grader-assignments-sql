package api

import (
	repo "a21hc3NpZ25tZW50/repository"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

type API struct {
	usersRepo    repo.UserRepository
	sessionsRepo repo.SessionsRepository
	todosRepo    repo.TodoRepository
	mux          *http.ServeMux
}

func (api *API) BaseViewPath() (*template.Template, error) {
	basePath, errPath := filepath.Abs("./template/html/*")
	if errPath != nil {
		return nil, errPath
	}
	var tmpl, err = template.ParseGlob(basePath)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func NewAPI(usersRepo repo.UserRepository, sessionsRepo repo.SessionsRepository, todosRepo repo.TodoRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo,
		sessionsRepo,
		todosRepo,
		mux,
	}

	mux.HandleFunc("/", api.homePage)
	mux.HandleFunc("/page/login", api.loginPage)
	mux.HandleFunc("/page/register", api.registerPage)
	mux.HandleFunc("/page/todo", api.todoPage)

	fileServer := http.FileServer(http.Dir("./template"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.Handle("/user/register", api.Post(http.HandlerFunc(api.Register)))
	mux.Handle("/user/login", api.Post(http.HandlerFunc(api.Login)))
	mux.Handle("/user/session/valid", api.Get(api.Auth(http.HandlerFunc(api.SessionValid))))
	mux.Handle("/user/logout", api.Get(api.Auth(http.HandlerFunc(api.Logout))))

	mux.Handle("/todo/add", api.Post(api.Auth(http.HandlerFunc(api.addTodo))))
	mux.Handle("/todo/remove", api.Delete(api.Auth(http.HandlerFunc(api.deleteTodo))))
	mux.Handle("/todo/change-status", api.Put(api.Auth(http.HandlerFunc(api.changeStatusTodo))))
	mux.Handle("/todo/list", api.Get(api.Auth(http.HandlerFunc(api.listTodo))))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
