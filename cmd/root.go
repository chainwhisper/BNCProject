package cmd

import (
	"fmt"
	"github.com/huangsuyu/BNCProject/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var (
	// Path to config
	cfgFile string

	// The actual app config
	server *api.Server

	// Version for the application. Set via ldflags
	Version = "1.0"

	//// Commit (git) for the application. Set via ldflags
	//Commit = "undefined"
	//
	//// Branch (git) for the application. Set via ldflags
	//Branch = "undefined"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bnc-micro",
	Short: "A microservice for querying data from ",
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
		//Commit:  Commit,
		//Branch:  Branch,
	}
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		//home, err := homedir.Dir()
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Println(exPath)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(exPath)
		viper.SetConfigName("decode-micro")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		viper.Unmarshal(&server)
	}
}