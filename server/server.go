package main

import (
	"context"
	"errors"
	"log"
	"sync"

	pb "github.com/michaelkleyn/tictactoe-grpc/proto"
)

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

	playerID := geneateID()
	mark := pb.Mark_O

	for _, player := range game.Players {
		if player.Mark == pb.Mark_0 {
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
		GameID:   game.GameID,
		PlayerID: playerID,
	}, nil
}

func main() {
	// Connect to the PostgreSQL database
}
