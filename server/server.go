/*
Package main implements a gRPC server for a Tic-Tac-Toe game, allowing players to create games, join existing games, make moves, and receive game updates.

This package uses Protocol Buffers for defining the game service and data structures.
*/

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/michaelkleyn/tictactoe-grpc/proto"
	"google.golang.org/grpc"

	// "google.golang.org/grpc/credentials/insecure"
	"github.com/google/uuid"
)

// server represents the gRPC server for managing the Tic-Tac-Toe games.
type server struct {
	pb.UnimplementedGameServiceServer
	games map[string]*GameSession
	mu    sync.Mutex
}

func NewServer() *server {
	return &server{
		games: make(map[string]*GameSession),
	}
}

type GameSession struct {
	GameID       string
	Players      map[string]*Player
	Board        [9]pb.Mark
	NextPlayerID string
	Status       pb.GameStatus
	UpdateChans  map[string]chan *pb.GameUpdate
	mu           sync.Mutex
}

type Player struct {
	PlayerID   string
	PlayerName string
	Mark       pb.Mark
}

func (s *server) CreateGame(ctx context.Context, req *pb.CreateGameRequest) (*pb.CreateGameResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	gameID := generateID()
	playerID := generateID()

	player := &Player{
		PlayerID:   playerID,
		PlayerName: req.PlayerName,
		Mark:       pb.Mark_X,
	}

	game := &GameSession{
		GameID:       gameID,
		Players:      map[string]*Player{playerID: player},
		Board:        [9]pb.Mark{},
		NextPlayerID: playerID,
		Status:       pb.GameStatus_WAITING_FOR_PLAYER,
		UpdateChans:  make(map[string]chan *pb.GameUpdate),
	}

	s.games[gameID] = game

	log.Printf("Game %s created by player %s", gameID, playerID)

	return &pb.CreateGameResponse{
		GameId:   gameID,
		PlayerId: playerID,
	}, nil
}

func (s *server) JoinGame(ctx context.Context, req *pb.JoinGameRequest) (*pb.JoinGameResponse, error) {
	s.mu.Lock()
	game, exists := s.games[req.GameId]
	s.mu.Unlock()

	if !exists {
		return nil, errors.New("Game Not Found")
	}

	game.mu.Lock()
	defer game.mu.Unlock()

	if len(game.Players) >= 2 {
		return nil, errors.New("Game Already Has 2 Players")
	}

	playerID := generateID()
	mark := pb.Mark_O

	for _, player := range game.Players {
		if player.Mark == pb.Mark_O {
			mark = pb.Mark_X
			break
		}
	}

	player := &Player{
		PlayerID:   playerID,
		PlayerName: req.PlayerName,
		Mark:       mark,
	}

	game.Players[playerID] = player
	game.Status = pb.GameStatus_IN_PROGRESS

	log.Printf("Player %s joined game %s", playerID, game.GameID)

	// Notify players about the game start
	game.broadcastUpdate()

	return &pb.JoinGameResponse{
		GameId:   game.GameID,
		PlayerId: playerID,
	}, nil
}

func (s *server) MakeMove(ctx context.Context, req *pb.MakeMoveRequest) (*pb.MakeMoveResponse, error) {
	s.mu.Lock()
	game, exists := s.games[req.GameId]
	s.mu.Unlock()

	if !exists {
		return nil, errors.New("Game Not Found")
	}

	game.mu.Lock()
	defer game.mu.Unlock()

	if game.NextPlayerID != req.PlayerId {
		return &pb.MakeMoveResponse{
			Success: false,
			Message: "Not Your Turn",
		}, nil
	}

	if req.Position < 0 || req.Position > 8 {
		return &pb.MakeMoveResponse{
			Success: false,
			Message: "Invalid Move",
		}, nil
	}

	if game.Board[req.Position] != pb.Mark_EMPTY {
		return &pb.MakeMoveResponse{
			Success: false,
			Message: "Space Already Taken",
		}, nil
	}

	player := game.Players[req.PlayerId]
	game.Board[req.Position] = player.Mark

	// Check for a win or draw
	winner := checkWinner(game.Board)
	if winner != pb.Mark_EMPTY {
		game.Status = pb.GameStatus_FINISHED
		game.broadcastUpdate()
		return &pb.MakeMoveResponse{
			Success: true,
			Message: fmt.Sprintf("Player %s wins!", req.PlayerId),
		}, nil
	}

	if isBoardFull(game.Board) {
		game.Status = pb.GameStatus_FINISHED
		game.broadcastUpdate()
		return &pb.MakeMoveResponse{
			Success: true,
			Message: "Game is a draw!",
		}, nil
	}

	// Switch turns
	for pid := range game.Players {
		if pid != req.PlayerId {
			game.NextPlayerID = pid
			break
		}
	}

	game.broadcastUpdate()

	return &pb.MakeMoveResponse{
		Success: true,
		Message: "move accepted",
	}, nil
}

func (s *server) StreamGameUpdates(req *pb.StreamGameUpdatesRequest, stream pb.GameService_StreamGameUpdatesServer) error {
	s.mu.Lock()
	game, exists := s.games[req.GameId]
	s.mu.Unlock()

	if !exists {
		return errors.New("Game Not Found")
	}

	game.mu.Lock()
	updateChan := make(chan *pb.GameUpdate, 10)
	game.UpdateChans[req.PlayerId] = updateChan
	game.mu.Unlock()

	defer func() {
		game.mu.Lock()
		delete(game.UpdateChans, req.PlayerId)
		game.mu.Unlock()
	}()

	// Send initial game state
	err := stream.Send(game.toGameUpdate())
	if err != nil {
		return err
	}

	// Stream updates
	for update := range updateChan {
		if err := stream.Send(update); err != nil {
			return err
		}
	}

	return nil
}

// Utility functions

func (game *GameSession) broadcastUpdate() {
	gameUpdate := game.toGameUpdate()
	for _, ch := range game.UpdateChans {
		ch := ch
		select {
		case ch <- gameUpdate:
		default:
			// Skip if the channel is full
		}
	}
}

func (game *GameSession) toGameUpdate() *pb.GameUpdate {
	cells := make([]*pb.Cell, 9)
	for i, mark := range game.Board {
		cells[i] = &pb.Cell{
			Position: int32(i),
			Mark:     mark,
		}
	}

	boardState := &pb.BoardState{
		Cells: cells,
	}

	return &pb.GameUpdate{
		GameId:       game.GameID,
		BoardState:   boardState,
		NextPlayerId: game.NextPlayerID,
		Status:       game.Status,
	}
}

func checkWinner(board [9]pb.Mark) pb.Mark {
	winningCombinations := [][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, combo := range winningCombinations {
		if board[combo[0]] != pb.Mark_EMPTY &&
			board[combo[0]] == board[combo[1]] &&
			board[combo[1]] == board[combo[2]] {
			return board[combo[0]]
		}
	}

	return pb.Mark_EMPTY
}

func isBoardFull(board [9]pb.Mark) bool {
	for _, mark := range board {
		if mark == pb.Mark_EMPTY {
			return false
		}
	}
	return true
}

func generateID() string {
	return uuid.NewString()
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGameServiceServer(grpcServer, NewServer())

	log.Println("Game server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
