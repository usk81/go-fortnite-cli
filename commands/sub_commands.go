package commands

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	fortnite "github.com/usk81/go-fortnite"
)

var (
	statsCmd = &cobra.Command{
		Use:   "stats",
		Short: "Get Fornite BR Player Stats",
		Long:  "Get Fornite BR Player Stats",
		Run:   statsCommand,
	}
	historyCmd = &cobra.Command{
		Use:   "history",
		Short: "Get Fornite Match History",
		Long:  "Get Fornite Match History",
		Run:   historyCommand,
	}
	storeCmd = &cobra.Command{
		Use:   "store",
		Short: "Get Current Fortnite Store",
		Long:  "Get Current Fortnite Store",
		Run:   storeCommand,
	}
	challengeCmd = &cobra.Command{
		Use:   "challenges",
		Short: "Get Current Active Challenges",
		Long:  "Get Current Active Challenges",
		Run:   challengeCommand,
	}
	// ErrNotSetAccessToken is returned when access token is not set
	ErrNotSetAccessToken = errors.New("access token is not set. get from TRN: https://fortnitetracker.com/site-api")
	token                string
)

func statsCommand(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		Exit(fmt.Errorf("How to run:\n\tstats [platform] [nickname]"), 1)
	}
	if token == "" {
		Exit(ErrNotSetAccessToken, 1)
	}
	s, err := fortnite.New(nil, token)
	if err != nil {
		Exit(err, 1)
	}
	result, err := s.BRPlayerStats(args[0], args[1])
	if err != nil {
		Exit(err, 1)
	}
	fmt.Printf("%+v\n", result)
}

func historyCommand(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		Exit(fmt.Errorf("How to run:\n\thistory [account id]"), 1)
	}
	if token == "" {
		Exit(ErrNotSetAccessToken, 1)
	}
	s, err := fortnite.New(nil, token)
	if err != nil {
		Exit(err, 1)
	}
	result, err := s.MatchHistory(args[0])
	if err != nil {
		Exit(err, 1)
	}
	fmt.Printf("%+v\n", result)
}

func storeCommand(cmd *cobra.Command, args []string) {
	if token == "" {
		Exit(ErrNotSetAccessToken, 1)
	}
	s, err := fortnite.New(nil, token)
	if err != nil {
		Exit(err, 1)
	}
	result, err := s.Store()
	if err != nil {
		Exit(err, 1)
	}
	fmt.Printf("%+v\n", result)
}

func challengeCommand(cmd *cobra.Command, args []string) {
	if token == "" {
		Exit(ErrNotSetAccessToken, 1)
	}
	s, err := fortnite.New(nil, token)
	if err != nil {
		Exit(err, 1)
	}
	result, err := s.Challenges()
	if err != nil {
		Exit(err, 1)
	}
	fmt.Printf("%+v\n", result)
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "access token")
	RootCmd.AddCommand(statsCmd)
	RootCmd.AddCommand(historyCmd)
	RootCmd.AddCommand(storeCmd)
	RootCmd.AddCommand(challengeCmd)
}
