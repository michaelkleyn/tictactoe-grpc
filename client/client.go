package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	pb "github.com/michaelkleyn/tictactoe-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, error := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if error != nil {
		log.Fatalf("Could not connect to server: %v", error)
	}
	defer conn.Close()

	client := pb.NewGameServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter username: ")
	playerName, _ := reader.ReadString('\n')
	playerName = strings.TrimSpace(playerName)

	fmt.Println("1. Create Game")
	fmt.Println("2. Join Game")
	fmt.Print("Choice: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var gameID, playerID string
	var waitGroup sync.WaitGroup

	switch choice {
	case "1":
		resp, err := client.CreateGame(context.Background(), &pb.CreateGameRequest{
			PlayerName: playerName,
		})
		if err != nil {
			log.Fatalf("Error creating game: %v", err)
		}
		gameID = resp.GameId
		playerID = resp.PlayerId
		fmt.Printf("Game created. Game ID: %s\n", gameID)
		fmt.Println("Waiting for another player to join...")

	case "2":
		fmt.Print("Enter Game ID to Join: ")
		inputGameID, _ := reader.ReadString('\n')
		gameID = strings.TrimSpace(inputGameID)
		resp, err := client.JoinGame(context.Background(), &pb.JoinGameRequest{
			GameId:     gameID,
			PlayerName: playerName,
		})
		if err != nil {
			log.Fatalf("Error joining game: %v", err)
		}
		playerID = resp.PlayerId
		fmt.Printf("Joined game. Game ID: %s\n", gameID)

	default:
		fmt.Println("Invalid choice")
		return
	}

	// Channel to coordinate between update stream and move input
	gameStateChan := make(chan *pb.GameUpdate, 1)
	waitGroup.Add(1)

	// Start streaming updates in a goroutine
	go func() {
		defer waitGroup.Done()
		streamGameUpdates(client, gameID, playerID, gameStateChan)
	}()

	// Handle game moves
	handleGameplay(client, gameID, playerID, reader, gameStateChan)

	// Clean up
	close(gameStateChan)
	waitGroup.Wait()
}

func streamGameUpdates(client pb.GameServiceClient, gameID, playerID string, gameStateChan chan<- *pb.GameUpdate) {
	ctx := context.Background()
	stream, err := client.StreamGameUpdates(ctx, &pb.StreamGameUpdatesRequest{
		GameId:   gameID,
		PlayerId: playerID,
	})
	if err != nil {
		log.Printf("Error starting update stream: %v", err)
		return
	}

	for {
		update, err := stream.Recv()
		if err != nil {
			log.Printf("Stream closed: %v", err)
			return
		}

		displayGameUpdate(update)
		gameStateChan <- update

		// Show prompt for move if it's this player's turn
		if update.NextPlayerId == playerID {
			fmt.Print("\nYour turn! Enter position (0-8), or 'exit' to quit: ")
		} else if update.Status == pb.GameStatus_WAITING_FOR_PLAYER {
			fmt.Println("\nWaiting for another player to join...")
		} else if update.Status == pb.GameStatus_IN_PROGRESS {
			fmt.Println("\nWaiting for other player's move...")
		} else if update.Status == pb.GameStatus_FINISHED {
			fmt.Println("\nGame Over!")
			return
		}
	}
}

func handleGameplay(client pb.GameServiceClient, gameID, playerID string, reader *bufio.Reader, gameStateChan <-chan *pb.GameUpdate) {
	for update := range gameStateChan {
		if update.Status == pb.GameStatus_FINISHED {
			return
		}

		if update.NextPlayerId == playerID {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "exit" {
				return
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

			if !resp.Success {
				fmt.Printf("Invalid move: %s\n", resp.Message)
			}
		}
	}
}

func parsePosition(input string) (int, error) {
	var position int
	_, err := fmt.Sscanf(input, "%d", &position)
	if err != nil || position < 0 || position > 8 {
		return 0, fmt.Errorf("invalid position")
	}
	return position, nil
}

func displayGameUpdate(update *pb.GameUpdate) {
	fmt.Print("\033[H\033[2J") // Clear screen
	fmt.Println("\nCurrent Board:")
	board := update.BoardState.Cells
	for i := 0; i < 9; i += 3 {
		fmt.Printf(" %s | %s | %s\n",
			getMarkSymbol(board[i].Mark),
			getMarkSymbol(board[i+1].Mark),
			getMarkSymbol(board[i+2].Mark))
		if i < 6 {
			fmt.Println("-----------")
		}
	}
}

func getMarkSymbol(mark pb.Mark) string {
	switch mark {
	case pb.Mark_X:
		return "X"
	case pb.Mark_O:
		return "O"
	default:
		return " "
	}
}
