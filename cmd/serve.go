package cmd

import (
	"fmt"
	"github.com/rifatikbal/E-Com-Gateway/internal/config"
	"github.com/rifatikbal/E-Com-Gateway/internal/conn"
	"log"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve will serve the gateway apis",
	Long:  `serve will serve the gateway apis`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config.LoadDBCfg()
		config.LoadAppCfg()

		if err := conn.ConnectDB(config.DB()); err != nil {
			log.Println(err)
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("serve called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
