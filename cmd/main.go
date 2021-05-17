package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"

	//	"net"

	//	"google.golang.org/grpc"

	"github.com/zshearin/poker/cmd/poker"
	proto "github.com/zshearin/poker/cmd/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	handleServer()

	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterPokerServiceServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(lis); e != nil {
		panic(e)
	}
	/*
		lis, err := net.Listen("tcp", ":9000")
		if err != nil {
			log.Fatalf("failed to listen on port 9000: %v", err)
		}

		s := poker.Server{}

		grpcServer := grpc.NewServer()

		//poker.RegisterChatServiceServer(grpcServer, &s)

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
		}
	*/

}

func handleServer() {
	http.HandleFunc("/", pokerHandler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func pokerHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		deal1 := shuffleAndDeal(4)

		//deal1.PrintBoardAndHands()
		//deal1.Players.Print()

		type result struct {
			Board       poker.Cards
			Hands       poker.Hands
			HandResults []poker.HandResult
		}
		res := result{
			Board:       deal1.GetBoard(),
			Hands:       deal1.Hands,
			HandResults: deal1.HandResults,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
		//Now format the body with all the objects we want
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}

func shuffleAndDeal(players int) poker.Deal {

	deck := poker.GetDeck()
	deck.Shuffle()
	deck.Shuffle()
	//deck.PrintOrder()
	return deck.GetDeal(players)
}

/////////////////////////////////////////////////////////////////////////////////////////
//
/////////////////////////////////////////////////////////////////////////////////////////
type server struct {
}

func (s *server) GetDeal(ctx context.Context, request *proto.Request) (*proto.Deal, error) {
	numHandsInDeal := request.GetHands()

	deal := shuffleAndDeal(int(numHandsInDeal))

	board := deal.GetBoard()
	hands := deal.Hands
	handResults := deal.HandResults

	var newBoard []*proto.Card

	for _, card := range board {
		suit := card.Suit
		value := card.Value
		var newCard *proto.Card
		newCard.Suit = suit
		newCard.Value = value
		newBoard = append(newBoard, newCard)
	}

	var newHands []*proto.Hand

	for _, hand := range hands {

		var curNewHand *proto.Hand

		for _, card := range hand {
			var curNewCard *proto.Card

			curNewCard.Suit = card.Suit
			curNewCard.Value = card.Value
			curNewHand.Hand = append(curNewHand.Hand, curNewCard)
		}

		newHands = append(newHands, curNewHand)
	}

	var newResults []*proto.HandResult

	for _, result := range handResults {
		var newResult *proto.HandResult

		newResult.HandName = result.HandName
		newResult.Player = int32(result.PlayerNumber)

		newResults = append(newResults, newResult)

	}

	return &proto.Deal{
		Board:   newBoard,
		Hands:   newHands,
		Results: newResults,
	}, nil

}
