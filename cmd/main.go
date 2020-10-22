package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"playing-with-golang-on-k8s/auth"
	"playing-with-golang-on-k8s/routes"
	"playing-with-golang-on-k8s/server"
	"playing-with-golang-on-k8s/service"
	"playing-with-golang-on-k8s/store"
	"playing-with-golang-on-k8s/store/postgres"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ory/viper"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main() {
	cmd := &cobra.Command{
		Use:   "playing-with-golang-on-k8s",
		Short: "playing-with-golang-on-k8s starts the API server",
		Long:  "playing-with-golang-on-k8s is the launcher for the server handling all API endpoints and connecting to our data storage",
		RunE:  start,
	}

	f := cmd.Flags()
	f.String("dbHost", "localhost", "database host")
	f.Int("dbPort", 5432, "database port")
	f.String("dbName", "", "database name")
	f.String("dbUser", "", "database username")
	f.String("dbPassword", "", "database password")

	if err := viper.BindPFlags(f); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Error handling flags in viper.BindPFlags:", err)
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

func start(cmd *cobra.Command, args []string) error {
	isProd := viper.Get("ENV") == "production"
	//Context
	ctx, cancel := context.WithCancel(context.Background())
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		// block until receiving interrupt signal, then close context
		<-signals
		cancel()
	}()

	// Logger
	var logger *zap.Logger
	var err error
	if isProd {
		logger, err = zap.NewProduction()
		gin.SetMode("release")
	} else {
		logger, err = zap.NewDevelopment()
		gin.SetMode("debug")
	}
	if err != nil {
		cancel()
		return errors.Wrap(err, "failed to instantiate logger")
	}
	defer logger.Sync()

	ms := []interface{}{
		&store.Product{},
		&store.User{},
	}
	db, err := postgres.NewClient(postgres.NewConfig(), ms)
	if err != nil {
		cancel()
		return errors.Wrap(err, "failed to initialize DB connection")
	}
	defer db.GracefulShutdown()

	permsService := auth.NewPermissionService(ctx, db)
	userService := service.NewUserService(db)
	proService := service.NewProService(db)
	/*es, err := es.NewClient(ctx)
	if err != nil {
		cancel()
		return errors.Wrap(err, "Oops err when creating elastic client")
	}*/
	//indexService := service.NewIndex(es, db)

	userActions := routes.NewUserActions(userService)
	proActions := routes.NewProductActions(db, proService, permsService)

	authCfg := auth.NewConfig()
	authMiddleware, err := auth.NewAuthMiddleware(authCfg, db)
	if err != nil {
		log.Fatalf("Oops err when creating auth middleware %s", err)
	}
	serverConfig := server.NewConfig(authMiddleware)

	server := server.Server{
		Config:            serverConfig,
		UserActions:       userActions,
		ProdsActions:      proActions,
		PermissionService: permsService,
	}

	errChan := make(chan error)

	go func() {
		logger.Debug("running the server")
		if err := server.Run(); err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		db.GracefulShutdown()
		return handleError(logger, cancel, err, "runtime")
	case <-ctx.Done():
		db.GracefulShutdown()
		cancel()
		return nil
	}
}

func handleError(log *zap.Logger, cancel context.CancelFunc, err error, detail string) error {
	log.Error(fmt.Sprintf("%s error", detail),
		zap.Error(err),
	)
	cancel()
	os.Exit(1)
	return nil
}
