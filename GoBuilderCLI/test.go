package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	flag1   bool
	copydir bool
	rootCmd = &cobra.Command{
		Use:   "gobuildercli",
		Short: "test",
		Long:  "Go Builder CLI",
		// SilenceUsage: true,
	}
	buildexecute = &cobra.Command{
		Use:   "buildexecute [flags]",
		Short: "buildexecute runner",
		Long:  "Go Builder buildexecute",
		// SilenceUsage: true,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if args != nil {
				fmt.Println("inside buildexecute command :", args)
				if copydir != false {
					fmt.Println("directory received : ", string(args[0]), copydir)
				}
			} else {
				fmt.Println("inside buildexecute command no args")
			}
		},
	}
)

func init() {

	rootCmd.PersistentFlags().BoolVarP(&flag1, "psersistFlag", "p", false, "persistant test flag in root")
	buildexecute.Flags().BoolVarP(&copydir, "copydir", "c", false, "-c [destination]")
	rootCmd.AddCommand(buildexecute)
}
func main() {
	// cmd.AddCommand(printTimeCmd())
	// if err := cmd.Execute(); err != nil {
	// 	os.Exit(1)
	// }
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
