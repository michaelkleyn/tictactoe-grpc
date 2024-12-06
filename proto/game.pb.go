// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: proto/game.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GameStatus int32

const (
	GameStatus_WAITING_FOR_PLAYER GameStatus = 0
	GameStatus_IN_PROGRESS        GameStatus = 1
	GameStatus_FINISHED           GameStatus = 2
)

// Enum value maps for GameStatus.
var (
	GameStatus_name = map[int32]string{
		0: "WAITING_FOR_PLAYER",
		1: "IN_PROGRESS",
		2: "FINISHED",
	}
	GameStatus_value = map[string]int32{
		"WAITING_FOR_PLAYER": 0,
		"IN_PROGRESS":        1,
		"FINISHED":           2,
	}
)

func (x GameStatus) Enum() *GameStatus {
	p := new(GameStatus)
	*p = x
	return p
}

func (x GameStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_game_proto_enumTypes[0].Descriptor()
}

func (GameStatus) Type() protoreflect.EnumType {
	return &file_proto_game_proto_enumTypes[0]
}

func (x GameStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameStatus.Descriptor instead.
func (GameStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_game_proto_rawDescGZIP(), []int{0}
}

type Mark int32

const (
	Mark_EMPTY Mark = 0
	Mark_X     Mark = 1
	Mark_O     Mark = 2
)

// Enum value maps for Mark.
var (
	Mark_name = map[int32]string{
		0: "EMPTY",
		1: "X",
		2: "O",
	}
	Mark_value = map[string]int32{
		"EMPTY": 0,
		"X":     1,
		"O":     2,
	}
)

func (x Mark) Enum() *Mark {
	p := new(Mark)
	*p = x
	return p
}

func (x Mark) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Mark) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_game_proto_enumTypes[1].Descriptor()
}

func (Mark) Type() protoreflect.EnumType {
	return &file_proto_game_proto_enumTypes[1]
}

func (x Mark) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Mark.Descriptor instead.
func (Mark) EnumDescriptor() ([]byte, []int) {
	return file_proto_game_proto_rawDescGZIP(), []int{1}
}

type CreateGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerName string `protobuf:"bytes,1,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
}

func (x *CreateGameRequest) Reset() {
	*x = CreateGameRequest{}
	mi := &file_proto_game_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameRequest) ProtoMessage() {}

func (x *CreateGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameRequest.ProtoReflect.Descriptor instead.
func (*CreateGameRequest) Descriptor() ([]byte, []int) {
	return file_proto_game_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGameRequest) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

type CreateGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId   string `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	PlayerId string `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *CreateGameResponse) Reset() {
	*x = CreateGameResponse{}
	mi := &file_proto_game_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameResponse) ProtoMessage() {}

func (x *CreateGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameResponse.ProtoReflect.Descriptor instead.
func (*CreateGameResponse) Descriptor() ([]byte, []int) {
	return file_proto_game_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGameResponse) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *CreateGameResponse) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type JoinGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId     string `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	PlayerName string `protobuf:"bytes,2,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
}

func (x *JoinGameRequest) Reset() {
	*x = JoinGameRequest{}
	mi := &file_proto_game_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGameRequest) ProtoMessage() {}

func (x *JoinGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGameRequest.ProtoReflect.Descriptor instead.
func (*JoinGameRequest) Descriptor() ([]byte, []int) {
	return file_proto_game_proto_rawDescGZIP(), []int{2}
}

func (x *JoinGameRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *JoinGameRequest) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

type JoinGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId   string `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	PlayerId string `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *JoinGameResponse) Reset() {
	*x = JoinGameResponse{}
	mi := &file_proto_game_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGameResponse) ProtoMessage() {}

func (x *JoinGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGameResponse.ProtoReflect.Descriptor instead.
func (*JoinGameResponse) Descriptor() ([]byte, []int) {
	return file_proto_game_proto_rawDescGZIP(), []int{3}
}

func (x *JoinGameResponse) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *JoinGameResponse) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type MakeMoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId   string `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	PlayerId string `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Position int32  `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"` // Positions from 0 to 8 for the 9 cells
}

func (x *MakeMoveRequest) Reset() {
	*x = MakeMoveRequest{}
	mi := &file_proto_game_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MakeMoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeMoveRequest) ProtoMessage() {}

func (x *MakeMoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeMoveRequest.ProtoReflect.Descriptor instead.
func (*MakeMoveRequest) Descriptor() ([]byte, []int) {
	return file_proto_game_proto_rawDescGZIP(), []int{4}
}

func (x *MakeMoveRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *MakeMoveRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *MakeMoveRequest) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

type MakeMoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MakeMoveResponse) Reset() {
	*x = MakeMoveResponse{}
	mi := &file_proto_game_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MakeMoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeMoveResponse) ProtoMessage() {}

func (x *MakeMoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeMoveResponse.ProtoReflect.Descriptor instead.
func (*MakeMoveResponse) Descriptor() ([]byte, []int) {
	return file_proto_game_proto_rawDescGZIP(), []int{5}
}

func (x *MakeMoveResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *MakeMoveResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StreamGameUpdatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId   string `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	PlayerId string `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *StreamGameUpdatesRequest) Reset() {
	*x = StreamGameUpdatesRequest{}
	mi := &file_proto_game_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamGameUpdatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamGameUpdatesRequest) ProtoMessage() {}

func (x *StreamGameUpdatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamGameUpdatesRequest.ProtoReflect.Descriptor instead.
func (*StreamGameUpdatesRequest) Descriptor() ([]byte, []int) {
	return file_proto_game_proto_rawDescGZIP(), []int{6}
}

func (x *StreamGameUpdatesRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *StreamGameUpdatesRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type GameUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId       string      `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	BoardState   *BoardState `protobuf:"bytes,2,opt,name=board_state,json=boardState,proto3" json:"board_state,omitempty"`
	NextPlayerId string      `protobuf:"bytes,3,opt,name=next_player_id,json=nextPlayerId,proto3" json:"next_player_id,omitempty"`
	Status       GameStatus  `protobuf:"varint,4,opt,name=status,proto3,enum=proto.GameStatus" json:"status,omitempty"`
}

func (x *GameUpdate) Reset() {
	*x = GameUpdate{}
	mi := &file_proto_game_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameUpdate) ProtoMessage() {}

func (x *GameUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameUpdate.ProtoReflect.Descriptor instead.
func (*GameUpdate) Descriptor() ([]byte, []int) {
	return file_proto_game_proto_rawDescGZIP(), []int{7}
}

func (x *GameUpdate) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *GameUpdate) GetBoardState() *BoardState {
	if x != nil {
		return x.BoardState
	}
	return nil
}

func (x *GameUpdate) GetNextPlayerId() string {
	if x != nil {
		return x.NextPlayerId
	}
	return ""
}

func (x *GameUpdate) GetStatus() GameStatus {
	if x != nil {
		return x.Status
	}
	return GameStatus_WAITING_FOR_PLAYER
}

type BoardState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cells []*Cell `protobuf:"bytes,1,rep,name=cells,proto3" json:"cells,omitempty"` // 9 cells for the board
}

func (x *BoardState) Reset() {
	*x = BoardState{}
	mi := &file_proto_game_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoardState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardState) ProtoMessage() {}

func (x *BoardState) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardState.ProtoReflect.Descriptor instead.
func (*BoardState) Descriptor() ([]byte, []int) {
	return file_proto_game_proto_rawDescGZIP(), []int{8}
}

func (x *BoardState) GetCells() []*Cell {
	if x != nil {
		return x.Cells
	}
	return nil
}

type Cell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position int32 `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	Mark     Mark  `protobuf:"varint,2,opt,name=mark,proto3,enum=proto.Mark" json:"mark,omitempty"`
}

func (x *Cell) Reset() {
	*x = Cell{}
	mi := &file_proto_game_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cell) ProtoMessage() {}

func (x *Cell) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cell.ProtoReflect.Descriptor instead.
func (*Cell) Descriptor() ([]byte, []int) {
	return file_proto_game_proto_rawDescGZIP(), []int{9}
}

func (x *Cell) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Cell) GetMark() Mark {
	if x != nil {
		return x.Mark
	}
	return Mark_EMPTY
}

var File_proto_game_proto protoreflect.FileDescriptor

var file_proto_game_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x4a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x0f, 0x4a,
	0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67,
	0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x63, 0x0a, 0x0f, 0x4d, 0x61, 0x6b, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x46, 0x0a, 0x10, 0x4d, 0x61, 0x6b, 0x65, 0x4d,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x50, 0x0a, 0x18, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x22, 0xaa, 0x01, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0b, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x0a, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a,
	0x0e, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2f,
	0x0a, 0x0a, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x05,
	0x63, 0x65, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x22,
	0x43, 0x0a, 0x04, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x04,
	0x6d, 0x61, 0x72, 0x6b, 0x2a, 0x43, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f,
	0x52, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e,
	0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46,
	0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x1f, 0x0a, 0x04, 0x4d, 0x61, 0x72,
	0x6b, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01,
	0x58, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x4f, 0x10, 0x02, 0x32, 0x95, 0x02, 0x0a, 0x0b, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a,
	0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x4d, 0x61,
	0x6b, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x61, 0x6b, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x47, 0x61, 0x6d, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x30, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x6b, 0x6c, 0x65, 0x79, 0x6e, 0x2f, 0x74, 0x69,
	0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_game_proto_rawDescOnce sync.Once
	file_proto_game_proto_rawDescData = file_proto_game_proto_rawDesc
)

func file_proto_game_proto_rawDescGZIP() []byte {
	file_proto_game_proto_rawDescOnce.Do(func() {
		file_proto_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_game_proto_rawDescData)
	})
	return file_proto_game_proto_rawDescData
}

var file_proto_game_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_game_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_game_proto_goTypes = []any{
	(GameStatus)(0),                  // 0: proto.GameStatus
	(Mark)(0),                        // 1: proto.Mark
	(*CreateGameRequest)(nil),        // 2: proto.CreateGameRequest
	(*CreateGameResponse)(nil),       // 3: proto.CreateGameResponse
	(*JoinGameRequest)(nil),          // 4: proto.JoinGameRequest
	(*JoinGameResponse)(nil),         // 5: proto.JoinGameResponse
	(*MakeMoveRequest)(nil),          // 6: proto.MakeMoveRequest
	(*MakeMoveResponse)(nil),         // 7: proto.MakeMoveResponse
	(*StreamGameUpdatesRequest)(nil), // 8: proto.StreamGameUpdatesRequest
	(*GameUpdate)(nil),               // 9: proto.GameUpdate
	(*BoardState)(nil),               // 10: proto.BoardState
	(*Cell)(nil),                     // 11: proto.Cell
}
var file_proto_game_proto_depIdxs = []int32{
	10, // 0: proto.GameUpdate.board_state:type_name -> proto.BoardState
	0,  // 1: proto.GameUpdate.status:type_name -> proto.GameStatus
	11, // 2: proto.BoardState.cells:type_name -> proto.Cell
	1,  // 3: proto.Cell.mark:type_name -> proto.Mark
	2,  // 4: proto.GameService.CreateGame:input_type -> proto.CreateGameRequest
	4,  // 5: proto.GameService.JoinGame:input_type -> proto.JoinGameRequest
	6,  // 6: proto.GameService.MakeMove:input_type -> proto.MakeMoveRequest
	8,  // 7: proto.GameService.StreamGameUpdates:input_type -> proto.StreamGameUpdatesRequest
	3,  // 8: proto.GameService.CreateGame:output_type -> proto.CreateGameResponse
	5,  // 9: proto.GameService.JoinGame:output_type -> proto.JoinGameResponse
	7,  // 10: proto.GameService.MakeMove:output_type -> proto.MakeMoveResponse
	9,  // 11: proto.GameService.StreamGameUpdates:output_type -> proto.GameUpdate
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_game_proto_init() }
func file_proto_game_proto_init() {
	if File_proto_game_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_game_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_game_proto_goTypes,
		DependencyIndexes: file_proto_game_proto_depIdxs,
		EnumInfos:         file_proto_game_proto_enumTypes,
		MessageInfos:      file_proto_game_proto_msgTypes,
	}.Build()
	File_proto_game_proto = out.File
	file_proto_game_proto_rawDesc = nil
	file_proto_game_proto_goTypes = nil
	file_proto_game_proto_depIdxs = nil
}
