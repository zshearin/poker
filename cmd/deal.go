/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
	dealer "github.com/zshearin/poker/cmd/dealer"
)

var (
	hands int
	print bool
)

func shuffleAndDeal(players int) dealer.Deal {
	deck := dealer.GetDeck()
	deck.Shuffle()
	deck.Shuffle()
	//deck.PrintOrder()
	return deck.GetDeal(players)
}

func dealCmd() *cobra.Command {
	dealCmd := &cobra.Command{
		Use:   "deal",
		Short: "Deal holdem hand and get outcome",
		RunE:  runDealCmd,
	}
	dealCmd.Flags().BoolVar(&print, "print", true, "")
	dealCmd.Flags().IntVar(&hands, "hands", 6, "")
	return dealCmd
}

func runDealCmd(cmd *cobra.Command, args []string) error {
	if hands > 10 {
		return errors.New("maximum number of hands to deal in is 10")
	}
	game := shuffleAndDeal(hands)
	if print {
		game.PrintBoardAndHands()
		game.PrintRanksAndBestFive()
	}
	return nil
}

func init() {
	rootCmd.AddCommand(dealCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dealCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dealCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
