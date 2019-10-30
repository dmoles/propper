package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var seed int64
	var cmd = &cobra.Command{
		Use: "propper",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			New(seed).Go()
		},
	}
	cmd.Flags().Int64VarP(&seed, "seed", "s", 0, "random seed (0 == current Unix timestamp)")
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}