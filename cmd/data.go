package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dataCmd)
}

var dataCmd = &cobra.Command{
	Use:                        "data",
	Short:                      "generates a data directory with movie names based on the './movies.txt' file",
	Example:                    "vida data",
	Run: func(cmd *cobra.Command, args []string) {
		_ : os.Mkdir("./data", 0777);

		file, err := os.Open("./movies.txt")
		if err != nil {
			exit(errors.Wrap(err, "could not open './movies.txt' file"))
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			name := "./data/" + strings.Join(strings.Split(strings.ToLower(strings.Split(scanner.Text(), "	")[1]), " "), "-") + ".mp4"
			f, err := os.Create(name)
			if err != nil {
				exit(errors.Wrap(err, "could not create file: " + name))
			}
			_ = f.Close()
		}
	},
}