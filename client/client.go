package main

import (
	"log"
  "fmt"
  "bufio"
  "strings"
  "os"
  "context"

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
    resp, err := client.CreateGame(context.Background(), &pb.CreateGameRequest{
                PlayerName: playerName,
            })
    if err != nil {
      log.Fatalf("Error creating game: %v", err)
    }
    // Create Game
    gameID = resp.GameId
    playerID = resp.PlayerId
    fmt.Printf("Game created. Game ID: %s\n", gameID)
  
  case "2":
    // Join Game
    fmt.Print("Enter Game ID to Join: ")
    gameID, _ := reader.ReadString('\n')
        gameID = strings.TrimSpace(gameID)
        resp, err := client.JoinGame(context.Background(), &pb.JoinGameRequest{
            GameId:     gameID,
            PlayerName: playerName,
        })
  default:
    fmt.Println("Invalid choice")
    return
}

//  Send request to server to create new session
//  Send request to server to join existing session
//
// Real Time Update
//  Receive game updates of Board State, Player Disconnection, Game Status
//  Display the game board to the user after each update
//
// Gameplay
//  Allow user to make moves by selecting positions on board
//  Send moves to server and handle server response
//
// Handle End Game
//  Detect when game has ended
