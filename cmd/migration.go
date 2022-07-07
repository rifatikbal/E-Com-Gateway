package cmd

import (
	"fmt"
	"github.com/rifatikbal/E-Com-Gateway/domain"
	"github.com/rifatikbal/E-Com-Gateway/internal/config"
	"github.com/rifatikbal/E-Com-Gateway/internal/conn"
	"github.com/spf13/cobra"
	"log"
)

var migrationCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migration command apply the db migration",
	Long:  `migration command apply the db migration`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config.LoadDBCfg()
		config.LoadAppCfg()

		if err := conn.ConnectDB(config.DB()); err != nil {
			log.Println(err)
			return err
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		conn.GetDB().GormDB.AutoMigrate(domain.User{})
		conn.GetDB().GormDB.AutoMigrate(domain.Token{})
	},
}

func init() {
	fmt.Println("initialized")
	rootCmd.AddCommand(migrationCmd)
}
