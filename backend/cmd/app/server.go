package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/rs/cors"
)

var (
	grpcAddr = ":7123"
	httpAddr = ":8088"
)

func newServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Poker service server command",
		Run:   runServerCmd,
	}

	//TODO - ADD FLAGS

	return cmd
}

func runServerCmd(cmd *cobra.Command, args []string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/game", getGame)
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(httpAddr, handler)
}

func getHandsValueOrDefault(queryParam string) (int, error) {
	// if not specified, use 6 as default
	if queryParam == "" {
		return 6, nil
	}

	hands, err := strconv.Atoi(queryParam)
	if err != nil {
		return 0, err
	} else if hands > 10 || hands < 2 {
		return hands, err
	}

	return hands, nil
}

func getPrintParameter(queryParam string) (bool, error) {
	if queryParam == "" {
		return false, nil
	}

	return strconv.ParseBool(queryParam)
}

func getGame(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("got /game request\n")

	hands, err := getHandsValueOrDefault(r.FormValue("hands"))
	if err != nil {
		io.WriteString(w, "Bad hands parameter\n")
		return
	}

	game := shuffleAndDeal(hands)

	shouldPrintStr := r.FormValue("print")
	shouldPrint, err := getPrintParameter(shouldPrintStr)
	if err != nil {
		io.WriteString(w, "Bad print parameter\n")
		return
	}
	if shouldPrint {
		game.PrintBoardAndHands()
		game.PrintRanksAndBestFive()
	}

	err = json.NewEncoder(w).Encode(game)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
