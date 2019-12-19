package cmd


import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println(fmt.Sprintf("Listening on port ':%v' Node: %s...", server.Port, server.Url))
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", server.Port), handlers.LoggingHandler(os.Stdout, server.Router())))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}