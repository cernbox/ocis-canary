package svc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/cernbox/ocis-canary/pkg/config"
	v0proto "github.com/cernbox/ocis-canary/pkg/proto/v0"
	revauser "github.com/cs3org/reva/pkg/user"
	"github.com/owncloud/ocis/ocis-pkg/log"

	//mysql package requires blank import
	_ "github.com/go-sql-driver/mysql"
)

var (
	// ErrMissingVersion defines the error if version is missing.
	ErrMissingVersion = errors.New("Version not specified")
	// ErrNoUser defines the error if there is no logged in user
	ErrNoUser = errors.New("No user supplied")
)

// NewService returns a service implementation for CanaryHandler.
func NewService(cfg *config.Config, logger log.Logger) v0proto.CanaryHandler {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Service.DB.Username, cfg.Service.DB.Password, cfg.Service.DB.Host, cfg.Service.DB.Port, cfg.Service.DB.Name))
	if err != nil {
		panic(err)
	}

	return Canary{
		config: cfg,
		logger: logger,
		db:     db,
	}
}

// Canary defines implements the business logic for CanaryHandler.
type Canary struct {
	config *config.Config
	logger log.Logger
	db     *sql.DB
}

// SetCanary implements the CanaryHandler interface.
func (s Canary) SetCanary(ctx context.Context, req *v0proto.CanaryRequest, rsp *v0proto.CanaryResponse) error {

	if req.Version == "" {
		return ErrMissingVersion
	}

	u, ok := revauser.ContextGetUser(ctx)
	if !ok {
		return ErrNoUser
	}

	s.logger.Info().
		Str("version", req.Version).
		Str("username", u.Username).
		Msg("Setting Canary version")

	query := fmt.Sprintf("INSERT INTO %s (username, adopter) VALUES (?, ?) ON DUPLICATE KEY UPDATE adopter = ?", s.config.Service.DB.Table)
	stmt, err := s.db.Prepare(query)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to prepare sql insert/update")
		return err
	}

	_, err = stmt.Exec(u.Username, req.Version, req.Version)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to execute sql insert/update")
		return err
	}

	rsp.Ttl = int32(s.config.Service.Cookie.TTL)
	rsp.Cookie = s.config.Service.Cookie.Name

	return nil
}
