package swagger

import (
	// External
	"mime"
	"net/http"

	"github.com/sirupsen/logrus"

	// Internal
	"github.com/iakrevetkho/archaeopteryx/config"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

type Server struct {
	log    *logrus.Entry
	config *config.Config

	fsHandler http.Handler
	hpHandler http.Handler
}

// Function creates Open API server
func New(config *config.Config) (*Server, error) {
	var err error
	s := new(Server)
	s.log = helpers.CreateComponentLogger("swagger")
	s.config = config

	if err := mime.AddExtensionType(".svg", "image/svg+xml"); err != nil {
		return nil, err
	}

	s.fsHandler, err = s.createFsHandler()
	if err != nil {
		return nil, err
	}
	s.hpHandler, err = s.createHomePageHandler()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		s.log.Println("Serve main page")
		s.hpHandler.ServeHTTP(w, r)
	} else {
		s.log.WithField("path", r.URL.Path).Println("Serve FS")
		s.fsHandler.ServeHTTP(w, r)
	}
}
