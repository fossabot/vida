package cmd

import (
	"fmt"

	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/gangachris/vida/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	migrationsDir = "./migrations"
	migrations    = "file://./migrations" // ¯\_(ツ)_/¯
)

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(createCmd)
	migrateCmd.AddCommand(upCmd)
	migrateCmd.AddCommand(downCmd)
	migrateCmd.AddCommand(rollbackCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations on the godi application",
}

var createCmd = &cobra.Command{
	Use:   "create [migration_name]",
	Short: "Create a new migration file",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one argument")
		}

		for _, v := range args {
			if strings.ContainsRune(v, '-') {
				return errors.Errorf("character %q is not allowed in migration file name %q. Please use underscores %q", '-', v, '_')
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("creating migration files")
		if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
			if err := os.Mkdir(migrationsDir, 0755); err != nil {
				color.Red("%v", err)
				os.Exit(1)
			}
		}

		for _, v := range args {
			timestamp := time.Now().Unix()
			upFileName := fmt.Sprintf("%s/%d_%s.up.sql", migrationsDir, timestamp, v)
			downFileName := fmt.Sprintf("%s/%d_%s.down.sql", migrationsDir, timestamp, v)

			for _, file := range []string{upFileName, downFileName} {
				if _, err := os.Create(file); err != nil {
					color.Red("could not create file %q: %v", file, err)
					os.Exit(1)
				}
				color.Green("created migration file %q", file)
			}
		}
	},
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Run up for all the migrations",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		if err := runMigrations(cfg, migrations, "up"); err != nil {
			color.Red("%v", err)
			os.Exit(1)
		}
		color.Green("up migrations ran successfully")
	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Run down for all the migrations",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		if err := runMigrations(cfg, migrations, "down"); err != nil {
			color.Red("%v", err)
			os.Exit(1)
		}
		color.Green("down migrations ran successfully")
	},
}

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback the previous migration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		if err := runMigrations(cfg, migrations, "rollback"); err != nil {
			color.Red("%v", err)
			os.Exit(1)
		}
	},
}

func runMigrations(cfg config.Config, migrations, migration string) error {
	m, err := migrate.New(migrations, cfg.Database().Postgres().DSN())
	if err != nil {
		return err
	}

	switch migration {
	case "up":
		return m.Up()
	case "down":
		return m.Down()
	case "rollback":
		return m.Steps(-1)
	default:
	}
	return nil
}
