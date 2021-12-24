package server

import (
	"net/http"

	"github.com/truescotian/golang-rest-projstructure/internal/httptransport"
	"github.com/truescotian/golang-rest-projstructure/internal/service/user"
	"github.com/truescotian/golang-rest-projstructure/pkg/logger"
)

func NewHTTPServer(addr string) error {
	// db, err := sql.Open("pgx", "...")
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return err
	// }

	// s := user.New(db)
	s := user.New()
	httptransport.UserHandler(s)

	logger.Infof("Running server on port ", addr)
	return http.ListenAndServe(addr, httptransport.Routes())
}
