package apiserver

import (
	"EmployeeAPI/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func Start(server *APIServer) error {
	if err := configLogger(server); err != nil {
		return err
	}
	configRouter(server)
	if err := server.configStore(); err != nil {
		return err
	}
	server.logger.Info("Starting API server")
	return http.ListenAndServe(server.config.BindAddr, server.router)
}

func (s *APIServer) configStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func configLogger(server *APIServer) error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}
	server.logger.SetLevel(level)
	return nil
}

func configRouter(server *APIServer) {
	server.router.HandleFunc("/hello", server.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := io.WriteString(writer, "Hello")
		if err != nil {
			return
		}
	}
}
