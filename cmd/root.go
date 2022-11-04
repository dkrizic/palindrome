/*
Copyright © 2022 Darko Krizic

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"github.com/dkrizic/palindrome/logic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var lowercase = false
var silent = false
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "palindrome",
	Short: "Checks if a given string is a palindrome",
	Long: `Checks if a given string is a palindrome. 
For example:

palindrome racecar
palindrome kayak
palindrome aibohphobia
palindrome radar
palindrome madam
palindrome level
palindrome rotor
palindrome civic
palindrome kayak
palindrome reviver
palindrome racecar
palindrome redder
palindrome madam
palindrome civic
`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) == 0 {
			cmd.Help()
			return nil
		}
		palindrome := args[0]
		isPalindrome, inputString := logic.IsPalindrome(palindrome, lowercase)
		if isPalindrome {
			if !silent {
				cmd.Printf("Yes, %s is a palindrome\n", inputString)
			}
			return nil
		} else {
			if !silent {
				cmd.Println("No, %s is not a palindrome\n", inputString)
			}
			return errors.New("No palindrome")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&lowercase, "lowercase", "l", false, "Converts the string to lowercase")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "Silent mode")
	viper.BindEnv("lowercase", "PALINDROME_LOWERCASE")
	viper.BindEnv("silent", "PALINDROME_SILENT")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
