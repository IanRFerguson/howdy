/*
Copyright © 2024 IAN RICHARD FERGUSON IANFERGUSONRVA@gmail.com
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "howdy user [--shout]",
	Args: cobra.ExactArgs(1),
	Short: "Send a greeting to the user",
	Long: `Supply a name at the command line and receive a greeting back`,
	Run: func(cmd *cobra.Command, args []string) { 
		// This is the value provided at the command line
		_USER := args[0]

		// We'll give the user the option to be shouted at
		_SHOUT, _ := cmd.Flags().GetBool("shout")

		// This is where we'll build the message to print
		message := buildMessage(_USER)

		// If the user wants to be shouted at, we'll capitalize the message
		if (_SHOUT) {
			message =  strings.ToUpper(message)
		}

		// We'll print the message to the command line here
		fmt.Println(message)
	},
}

func buildMessage(userName string) string {
	// Randomly select a phrase to attach to the message
	salutation := getSalutation()

	// Build out the string we want to print
	msg := "Howdy " + userName + "! " + salutation

	return msg
}

func getSalutation() string {
	// Build an array of salutation options
	options := [4]string{
		"Have a great day!",
		"You're the best!",
		"Everyone needs a friend like you!",
		"Keep up the great work!",
	}

	// Randomly pick an array index to return
	index := rand.Int() % len(options)

	return options[index]
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Give the user the option to shout the message
	rootCmd.PersistentFlags().Bool("shout", false, "Shout the greeting to the end user")
}


