package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/michaelkleyn/tictactoe-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Establish gRPC connection to Game Server
	conn, error := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if error != nil {
		log.Fatalf("Could not connect to server: %v", error)
	}
	defer conn.Close()

	client := pb.NewGameServiceClient(conn)

	// Prompt for username
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter username:")
	playerName, _ := reader.ReadString('\n')
	playerName = strings.TrimSpace(playerName)

	// Prompt for New/Join Game
	fmt.Println("1. Create Game")
	fmt.Println("2. Join Game")
	fmt.Print("Choice:")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var gameID, playerID string

	switch choice {
	case "1":
		// Create Game
		//  Send request to server to create new session
		resp, err := client.CreateGame(context.Background(), &pb.CreateGameRequest{
			PlayerName: playerName,
		})
		if err != nil {
			log.Fatalf("Error creating game: %v", err)
		}
		gameID = resp.GameId
		playerID = resp.PlayerId
		fmt.Printf("Game created. Game ID: %s\n", gameID)

	case "2":
		// Join Game
		//  Send request to server to join existing session
		fmt.Print("Enter Game ID to Join: ")
		gameID, _ := reader.ReadString('\n')
		gameID = strings.TrimSpace(gameID)
		resp, err := client.JoinGame(context.Background(), &pb.JoinGameRequest{
			GameId:     gameID,
			PlayerName: playerName,
		})
		if err != nil {
			log.Fatalf("Error joining game: %v", err)
		}
		playerID = resp.PlayerId
		fmt.Printf("Join game. Game ID: %s\n", gameID)
	default:
		fmt.Println("Invalid choice")
		return
	}
	// Real Time Update
	go streamGameUpdates(client, gameID, playerID)

	// Gameplay
	//  Allow user to make moves by selecting positions on board
	for {
		fmt.Print("Enter position (0-8) to make a move, or 'exit' to quit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		position, err := parsePosition(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 0 and 8.")
			continue
		}

		resp, err := client.MakeMove(context.Background(), &pb.MakeMoveRequest{
			GameId:   gameID,
			PlayerId: playerID,
			Position: int32(position),
		})
		if err != nil {
			fmt.Printf("Error making move: %v\n", err)
			continue
		}

		fmt.Println(resp.Message)
	}
}

// Parse Position
func parsePosition(input string) (int, error) {
	var position int
	_, err := fmt.Sscanf(input, "%d", &position)
	if err != nil || position < 0 || position > 8 {
		return 0, fmt.Errorf("Invalid Position")
	}
	return position, nil
}

// Receive game updates of Board State, Player Disconnection, Game Status
func streamGameUpdates(client pb.GameServiceClient, gameID, playerID string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.StreamGameUpdates(ctx, &pb.StreamGameUpdatesRequest{
		GameId:   gameID,
		PlayerId: playerID,
	})
	if err != nil {
		log.Fatalf("Error starting update stream: %v", err)
	}

	for {
		update, err := stream.Recv()
		if err != nil {
			log.Printf("Stream closed: %v", err)
			return
		}

		displayGameUpdate(update)
	}
}

// Display the game board to the user after each update
func displayGameUpdate(update *pb.GameUpdate) {
	fmt.Println("\nGame update:")
	board := update.BoardState.Cells
	for i, cell := range board {
		mark := cell.Mark.String()
		if mark == "EMPTY" {
			mark = "_"
		}
		fmt.Print(mark)
		if (i+1)%3 == 0 {
			fmt.Println()
		} else {
			fmt.Print("|")
		}
	}
	fmt.Printf("Next Player ID: %s\n", update.GetNextPlayerId())
	fmt.Printf("Game Status: %s\n", update.Status.String())
}

//  Send moves to server and handle server response
//
// Handle End Game
//  Detect when game has ended
