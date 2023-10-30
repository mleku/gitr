package cmd

import (
	"os"

	"github.com/mleku/signr/pkg/signr"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type config struct {
	*signr.Signr
	c              *cobra.Command
	DataDir        string
	verbose, color bool
}

var s config
var Pass, Custom string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitr",
	Short: "A Git CLI with nostr support",
	Long: `gitr
	
A replacement for git built using github.com/go-git/go-git that enables the use of nostr keys for authentication of commits, and integration into decentralised git hosting protocols such as HORNETS and IPFS.
`,
}

// Execute adds all child commands to the root command and sets flags
// appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var err error
	// because this is a CLI app we know the user can enter passwords this way.
	// other types of apps using this can load the environment variables.
	s.Signr, err = signr.Init(signr.PasswordEntryViaTTY)
	if err != nil {
		s.Fatal("fatal error: %s\n", err)
	}
	rootCmd.PersistentFlags().BoolVarP(&s.Verbose,
		"verbose", "v", false, "prints more things")
	rootCmd.PersistentFlags().BoolVarP(&s.Color,
		"color", "c", false, "prints things with colour")
	s.c = rootCmd
	cobra.OnInitialize(initConfig(s.Signr))
}

// initConfig reads in config file and ENV variables if set.
func initConfig(cfg *signr.Signr) func() {
	return func() {
		viper.SetConfigName(signr.ConfigName)
		viper.SetConfigType(signr.ConfigExt)
		viper.AddConfigPath(cfg.DataDir)
		// read in environment variables that match
		viper.SetEnvPrefix(signr.AppName)
		viper.AutomaticEnv()
		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil && cfg.Verbose {
			cfg.Log("Using config file: %s\n", viper.ConfigFileUsed())
		}

		// if pass is given on CLI it overrides environment, but if it is empty and environment has a value, load it
		if Pass == "" {
			if p := viper.GetString("pass"); p != "" {
				Pass = p
			}
		}

		cfg.DefaultKey = viper.GetString("default")
	}
}
