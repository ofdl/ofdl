package cmd

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup the database",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := os.Open(viper.GetString("database"))
		if err != nil {
			return err
		}
		defer db.Close()

		os.Mkdir("backups", 0755)
		backupName := filepath.Join("backups", fmt.Sprintf("ofdl.sqlite.%s.gz", time.Now().Format("2006-01-02_15-04-05")))
		backupFile, err := os.OpenFile(backupName, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer backupFile.Close()

		gzipWriter := gzip.NewWriter(backupFile)
		defer gzipWriter.Close()

		_, err = io.Copy(gzipWriter, db)
		if err != nil {
			return err
		}

		fmt.Printf("Backup saved to %s\n", backupName)
		return nil
	},
}

var restoreCmd = &cobra.Command{
	Use:   "restore {backup-file}",
	Short: "Restore the database from a backup",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		backup, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer backup.Close()

		fmt.Printf("This will overwrite the existing database at %s. Continue? (y/N) ", viper.GetString("database"))
		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "y" {
			return nil
		}

		gzipReader, err := gzip.NewReader(backup)
		if err != nil {
			return err
		}

		db, err := os.OpenFile(viper.GetString("database"), os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer db.Close()

		_, err = io.Copy(db, gzipReader)
		if err != nil {
			return err
		}

		fmt.Printf("Database restored from %s\n", args[0])
		return nil
	},
}

func init() {
	CLI.AddCommand(backupCmd)
	CLI.AddCommand(restoreCmd)
}
