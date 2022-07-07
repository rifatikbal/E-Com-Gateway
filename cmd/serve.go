package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"

	"github.com/rifatikbal/E-Com-Gateway/internal/config"
	"github.com/rifatikbal/E-Com-Gateway/internal/conn"
	userHttpDelivery "github.com/rifatikbal/E-Com-Gateway/user/delivery"
	userRepo "github.com/rifatikbal/E-Com-Gateway/user/repository"
	userUseCase "github.com/rifatikbal/E-Com-Gateway/user/usecase"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve will serve the gateway apis",
	Long:  `serve will serve the gateway apis`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config.LoadDBCfg()
		config.LoadAppCfg()
		config.LoadAuthCfg()

		if err := conn.ConnectDB(config.DB()); err != nil {
			log.Println(err)
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.App()
		userRepo := userRepo.New(conn.GetDB())
		userUC := userUseCase.New(userRepo)

		r := chi.NewRouter()

		apiRouter := chi.NewRouter()
		r.Mount("/api", apiRouter)

		userHttpDelivery.New(apiRouter, userUC, *config.Auth())

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)

		server := http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler: r,
		}

		go func() {
			log.Println("server started on : 8080")
			if err := server.ListenAndServe(); err != nil {
				log.Println("info shutting down server")
			}
		}()
		<-quit

		fmt.Println("serve called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
