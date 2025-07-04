/*
Copyright © 2025 Atharv Singh

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt" // For printing output to the terminal
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra" // Core CLI framework
	"github.com/spf13/viper" // For config management (if using --viper flag or config files)
	"log"
	"os" // Used for handling OS-level stuff like exit codes
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "JustDO",
	Short: "JustDO is a todo application1",
	Long: `JustDO will help you get more done in less time.
It's designed to be as simple as possible to help
you accomplish your goals.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var datafile string
var cfgFile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// To read in config files and set env if there
func initConfig() {
	home, _ := homedir.Dir()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(home)
		viper.SetConfigName(".JustDO")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}



func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.GO-CLI.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	cobra.OnInitialize(initConfig)
	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect Home Directory. Please set data file path using  --datafile.")

	}
	// adding flag
	rootCmd.PersistentFlags().StringVar(&datafile, "datafile", home+string(os.PathSeparator)+"tasks.json", "data file to store todos")
	viper.BindPFlag("datafile", rootCmd.PersistentFlags().Lookup("datafile"))
	viper.BindEnv("datafile") 
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file default is ($HOME/.JustDO.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
