package main

import (
	"flag"

	"github.com/go-course/06-error_handling/internal/storage/ontario"

	beerscli "github.com/go-course/06-error_handling/internal"
	"github.com/go-course/06-error_handling/internal/cli"
	"github.com/spf13/cobra"
	"github.com/winternal/storage/csv"
)

func main() {

	csvData := flag.Bool("csv", false, "load data from csv")
	flag.Parse()
	var repo beerscli.BeerRepo
	repo = csv.NewRepository()

	if !*csvData {
		repo = ontario.NewOntarioRepository()
	}

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(repo))
	rootCmd.Execute()
}
