/*
Copyright © 2023 Hamilton Geraldo Fantin hfantin@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

type Config struct {
	Port              string   `mapstructure:"port"`
	TargetServer      string   `mapstructure:"target-server"`
	RecResponse       bool     `mapstructure:"rec-response"`
	Endpoints         []string `mapstructure:"endpoints"`
	ResponseFilesPath string   `mapstructure:"response-files-path"`
}

var configuration *Config

var rootCmd = &cobra.Command{
	Use:   "srm",
	Short: "Simple Rest Mock",
	Long:  `Simple Rest Mock is a request/response interceptor that can replace the response of the target server returning the mock file content`,
	Run: func(cmd *cobra.Command, args []string) {
		validate()
		createResponseFilesDir()
		startServer()
	},
	Version: fmt.Sprintf("v%s %s/%s\n", versionNumber, runtime.GOOS, runtime.GOARCH),
}

func validate() {
	errors := []string{}
	if len(configuration.TargetServer) == 0 {
		errors = append(errors, "target server is invalid")
	}
	if len(configuration.Endpoints) == 0 {
		errors = append(errors, "endpoints list is empty")
	}
	if len(configuration.ResponseFilesPath) == 0 {
		errors = append(errors, "response files path is invalid")
	}
	if len(errors) > 0 {
		for _, e := range errors {
			fmt.Println("-", e)
		}
		os.Exit(1)
	}
}

func init() {
	showBanner()
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.srm/config.yaml)")
	// define flags
	rootCmd.Flags().StringP("port", "p", "9000", "server port")
	rootCmd.Flags().StringP("target-server", "t", "", "target server to intercept request/response")
	rootCmd.Flags().BoolP("rec-mode", "r", false, "recorde response")
	rootCmd.Flags().StringSliceP("endpoints", "e", []string{}, "endpoints filtered by regex")
	rootCmd.Flags().StringP("response-files-path", "f", "jsons", "path to write response files")
	// bind flags from yaml file
	viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
	viper.BindPFlag("target-server", rootCmd.Flags().Lookup("target-server"))
	viper.BindPFlag("rec-mode", rootCmd.Flags().Lookup("rec-mode"))
	viper.BindPFlag("endpoints", rootCmd.Flags().Lookup("endpoints"))
	viper.BindPFlag("response-files-path", rootCmd.Flags().Lookup("response-files-path"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		fmt.Println("using custom config file", cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		createConfigFile(home)

		cobra.CheckErr(err)

		// Search config in home directory with name ".simple-rest-mock" (without extension).
		viper.AddConfigPath(home + "/.srm")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "reading configuration from", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Println("configuration error", err)
		os.Exit(1)
	}

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func createResponseFilesDir() {
	home, _ := os.UserHomeDir()
	responseFilesDir := home + "/.srm/" + configuration.ResponseFilesPath
	_, err := os.Stat(responseFilesDir)
	if os.IsNotExist(err) {
		fmt.Printf("creating response files folder: %s\n", responseFilesDir)
		err := os.Mkdir(responseFilesDir, 0755)
		if err != nil {
			fmt.Printf("could not create response files folder: %s\n", err)
			os.Exit(1)
		}
	}
}

func createConfigFile(homeDir string) {
	configDir := homeDir + "/.srm"
	configFile := configDir + "/config.yaml"
	_, err := os.Stat(configDir)
	if os.IsNotExist(err) {
		fmt.Printf("creating config folder: %s\n", configDir)
		err := os.Mkdir(configDir, 0755)
		if err != nil {
			fmt.Printf("could not create config folder: %s\n", err)
			os.Exit(1)
		}
	}

	fileExists := fileExists(configFile)
	if !fileExists {
		fmt.Printf("creating config file: %s\n", configFile)
		f, err := os.Create(configFile)
		if err != nil {
			fmt.Printf("could not create config file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
	}

}
