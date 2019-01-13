package cobraCmd

import (
	"fmt"
	"os"

	"Product-Management-API/pkg/cmd"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start grpc server",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
 
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(c *cobra.Command, args []string) {
		if err := cmd.RunServer(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
