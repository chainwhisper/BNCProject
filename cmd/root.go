package cmd
import (
	"fmt"
	"os"

	"github.com/huangsuyu/BNCProject/api"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Path to config
	cfgFile string

	// The actual app config
	server *api.Server

	// Version for the application. Set via ldflags
	Version = "undefined"

	// Commit (git) for the application. Set via ldflags
	Commit = "undefined"

	// Branch (git) for the application. Set via ldflags
	Branch = "undefined"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "amino-micro",
	Short: "A microservice for encoding JSON to amino",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.amino-micro.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	server = &api.Server{
		Version: Version,
		Commit:  Commit,
		Branch:  Branch,
	}
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".amino-micro")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		viper.Unmarshal(&server)
	}
}