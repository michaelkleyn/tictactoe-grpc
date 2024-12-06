package main

import (
	"context"
	"fmt"
	"sync"
	"testing"

	pb "github.com/michaelkleyn/tictactoe-grpc/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Helper functions for testing
type testGame struct {
	gameID    string
	player1ID string
	player2ID string
	server    *server
}

func setupTestGame(t *testing.T) *testGame {
	srv := NewServer()
	game := &testGame{server: srv}

	// Create game with first player
	resp, err := srv.CreateGame(context.Background(), &pb.CreateGameRequest{
		PlayerName: "Player1",
	})
	require.NoError(t, err)
	game.gameID = resp.GameId
	game.player1ID = resp.PlayerId

	return game
}

func (g *testGame) addSecondPlayer(t *testing.T) {
	resp, err := g.server.JoinGame(context.Background(), &pb.JoinGameRequest{
		GameId:     g.gameID,
		PlayerName: "Player2",
	})
	require.NoError(t, err)
	g.player2ID = resp.PlayerId
}

// Unit Tests

func TestCreateGame(t *testing.T) {
	srv := NewServer()

	tests := []struct {
		name       string
		playerName string
		wantErr    bool
	}{
		{
			name:       "Valid game creation",
			playerName: "TestPlayer",
			wantErr:    false,
		},
		{
			name:       "Empty player name",
			playerName: "",
			wantErr:    false, // Current implementation allows empty names
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := srv.CreateGame(context.Background(), &pb.CreateGameRequest{
				PlayerName: tt.playerName,
			})

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, resp.GameId)
			assert.NotEmpty(t, resp.PlayerId)

			// Verify game state
			game, exists := srv.games[resp.GameId]
			assert.True(t, exists)
			assert.Equal(t, pb.GameStatus_WAITING_FOR_PLAYER, game.Status)
			assert.Equal(t, tt.playerName, game.Players[resp.PlayerId].PlayerName)
		})
	}
}

func TestJoinGame(t *testing.T) {
	game := setupTestGame(t)

	tests := []struct {
		name       string
		gameID     string
		playerName string
		wantErr    bool
		errCode    codes.Code
	}{
		{
			name:       "Valid join",
			gameID:     game.gameID,
			playerName: "Player2",
			wantErr:    false,
		},
		{
			name:       "Invalid game ID",
			gameID:     "invalid",
			playerName: "Player2",
			wantErr:    true,
		},
		{
			name:       "Join full game",
			gameID:     game.gameID,
			playerName: "Player3",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := game.server.JoinGame(context.Background(), &pb.JoinGameRequest{
				GameId:     tt.gameID,
				PlayerName: tt.playerName,
			})

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, resp.PlayerId)
			assert.Equal(t, tt.gameID, resp.GameId)

			// Verify game state
			g := game.server.games[resp.GameId]
			assert.Equal(t, pb.GameStatus_IN_PROGRESS, g.Status)
			assert.Len(t, g.Players, 2)
		})
	}
}

func TestMakeMove(t *testing.T) {
	game := setupTestGame(t)
	game.addSecondPlayer(t)

	tests := []struct {
		name      string
		gameID    string
		playerID  string
		position  int32
		wantErr   bool
		wantValid bool
	}{
		{
			name:      "Valid move by first player",
			gameID:    game.gameID,
			playerID:  game.player1ID,
			position:  0,
			wantErr:   false,
			wantValid: true,
		},
		{
			name:      "Move out of turn",
			gameID:    game.gameID,
			playerID:  game.player1ID,
			position:  1,
			wantErr:   false,
			wantValid: false,
		},
		{
			name:      "Invalid position",
			gameID:    game.gameID,
			playerID:  game.player2ID,
			position:  9,
			wantErr:   false,
			wantValid: false,
		},
		{
			name:      "Move on occupied position",
			gameID:    game.gameID,
			playerID:  game.player2ID,
			position:  0,
			wantErr:   false,
			wantValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := game.server.MakeMove(context.Background(), &pb.MakeMoveRequest{
				GameId:   tt.gameID,
				PlayerId: tt.playerID,
				Position: tt.position,
			})

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantValid, resp.Success)
		})
	}
}

func TestWinConditions(t *testing.T) {
	winningPatterns := [][]int32{
		{0, 1, 2}, // Top row
		{3, 4, 5}, // Middle row
		{6, 7, 8}, // Bottom row
		{0, 3, 6}, // Left column
		{1, 4, 7}, // Middle column
		{2, 5, 8}, // Right column
		{0, 4, 8}, // Diagonal
		{2, 4, 6}, // Diagonal
	}

	for _, pattern := range winningPatterns {
		t.Run(fmt.Sprintf("Win pattern %v", pattern), func(t *testing.T) {
			game := setupTestGame(t)
			game.addSecondPlayer(t)

			// Play winning pattern for player 1
			for i, pos := range pattern {
				// Player 1's moves
				resp, err := game.server.MakeMove(context.Background(), &pb.MakeMoveRequest{
					GameId:   game.gameID,
					PlayerId: game.player1ID,
					Position: pos,
				})
				require.NoError(t, err)
				assert.True(t, resp.Success)

				if i < len(pattern)-1 {
					// Player 2's moves (in non-winning positions)
					altPos := int32(8 - i) // Use positions from the bottom up
					if contains(pattern, altPos) {
						altPos--
					}
					resp, err = game.server.MakeMove(context.Background(), &pb.MakeMoveRequest{
						GameId:   game.gameID,
						PlayerId: game.player2ID,
						Position: altPos,
					})
					require.NoError(t, err)
					assert.True(t, resp.Success)
				}
			}

			// Verify game is finished and player 1 won
			g := game.server.games[game.gameID]
			assert.Equal(t, pb.GameStatus_FINISHED, g.Status)
		})
	}
}

// Integration Tests

func TestFullGameFlow(t *testing.T) {
	ctx := context.Background()
	game := setupTestGame(t)

	// Player 2 joins
	game.addSecondPlayer(t)

	// Setup update streams for both players
	updates1 := make(chan *pb.GameUpdate, 10)
	updates2 := make(chan *pb.GameUpdate, 10)

	var wg sync.WaitGroup
	wg.Add(2)

	// Start streaming for both players
	go func() {
		defer wg.Done()
		stream := &mockUpdateStream{updates: updates1}
		err := game.server.StreamGameUpdates(&pb.StreamGameUpdatesRequest{
			GameId:   game.gameID,
			PlayerId: game.player1ID,
		}, stream)
		require.NoError(t, err)
	}()

	go func() {
		defer wg.Done()
		stream := &mockUpdateStream{updates: updates2}
		err := game.server.StreamGameUpdates(&pb.StreamGameUpdatesRequest{
			GameId:   game.gameID,
			PlayerId: game.player2ID,
		}, stream)
		require.NoError(t, err)
	}()

	// Play a full game
	moves := []struct {
		playerID string
		position int32
	}{
		{game.player1ID, 0}, // X: top-left
		{game.player2ID, 4}, // O: center
		{game.player1ID, 1}, // X: top-middle
		{game.player2ID, 3}, // O: middle-left
		{game.player1ID, 2}, // X: top-right (winning move)
	}

	for _, move := range moves {
		resp, err := game.server.MakeMove(ctx, &pb.MakeMoveRequest{
			GameId:   game.gameID,
			PlayerId: move.playerID,
			Position: move.position,
		})
		require.NoError(t, err)
		assert.True(t, resp.Success)

		// Verify both players received the update
		update1 := <-updates1
		update2 := <-updates2
		assert.Equal(t, update1.BoardState, update2.BoardState)
	}

	// Verify game finished
	g := game.server.games[game.gameID]
	assert.Equal(t, pb.GameStatus_FINISHED, g.Status)
}

// Helper types and functions

type mockUpdateStream struct {
	grpc.ServerStream
	updates chan<- *pb.GameUpdate
}

func (s *mockUpdateStream) Send(update *pb.GameUpdate) error {
	s.updates <- update
	return nil
}

func contains(slice []int32, val int32) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
