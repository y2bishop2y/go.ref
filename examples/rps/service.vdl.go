// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: service.vdl

// Package rps defines interfaces for playing the game Rock-Paper-Scissors.  It
// is an example of a simple Vanadium service.
//
// http://en.wikipedia.org/wiki/Rock-paper-scissors
//
// There are three different roles in the game:
//
// 1. Judge: A judge enforces the rules of the game and decides who the winner
// is. At the end of the game, the judge reports the final score to all the
// score keepers.
//
// 2. Player: A player can ask a judge to start a new game, it can challenge
// another player, and it can play a game.
//
// 3. ScoreKeeper: A score keeper receives the final score for a game after it
// ended.
package rps

import (
	// VDL system imports
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/vdl"

	// VDL user imports
	"time"
	"v.io/v23/security/access"
	_ "v.io/v23/vdlroot/time"
)

// A GameId is used to uniquely identify a game within one Judge.
type GameId struct {
	Id string
}

func (GameId) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/examples/rps.GameId"`
}) {
}

// GameOptions specifies the parameters of a game.
type GameOptions struct {
	NumRounds int32       // The number of rounds that a player must win to win the game.
	GameType  GameTypeTag // The type of game to play: Classic or LizardSpock.
}

func (GameOptions) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/examples/rps.GameOptions"`
}) {
}

type GameTypeTag byte

func (GameTypeTag) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/examples/rps.GameTypeTag"`
}) {
}

type (
	// PlayerAction represents any single field of the PlayerAction union type.
	PlayerAction interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the PlayerAction union type.
		__VDLReflect(__PlayerActionReflect)
	}
	// PlayerActionMove represents field Move of the PlayerAction union type.
	PlayerActionMove struct{ Value string } // The move that the player wants to make.
	// PlayerActionQuit represents field Quit of the PlayerAction union type.
	PlayerActionQuit struct{ Value unused } // Indicates that the player is quitting the game.
	// __PlayerActionReflect describes the PlayerAction union type.
	__PlayerActionReflect struct {
		Name  string `vdl:"v.io/x/ref/examples/rps.PlayerAction"`
		Type  PlayerAction
		Union struct {
			Move PlayerActionMove
			Quit PlayerActionQuit
		}
	}
)

func (x PlayerActionMove) Index() int                         { return 0 }
func (x PlayerActionMove) Interface() interface{}             { return x.Value }
func (x PlayerActionMove) Name() string                       { return "Move" }
func (x PlayerActionMove) __VDLReflect(__PlayerActionReflect) {}

func (x PlayerActionQuit) Index() int                         { return 1 }
func (x PlayerActionQuit) Interface() interface{}             { return x.Value }
func (x PlayerActionQuit) Name() string                       { return "Quit" }
func (x PlayerActionQuit) __VDLReflect(__PlayerActionReflect) {}

type unused struct {
}

func (unused) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/examples/rps.unused"`
}) {
}

type (
	// JudgeAction represents any single field of the JudgeAction union type.
	JudgeAction interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the JudgeAction union type.
		__VDLReflect(__JudgeActionReflect)
	}
	// JudgeActionPlayerNum represents field PlayerNum of the JudgeAction union type.
	JudgeActionPlayerNum struct{ Value int32 } // The player's number.
	// JudgeActionOpponentName represents field OpponentName of the JudgeAction union type.
	JudgeActionOpponentName struct{ Value string } // The name of the opponent.
	// JudgeActionMoveOptions represents field MoveOptions of the JudgeAction union type.
	JudgeActionMoveOptions struct{ Value []string } // A list of allowed moves that the player must choose from.
	// JudgeActionRoundResult represents field RoundResult of the JudgeAction union type.
	JudgeActionRoundResult struct{ Value Round } // The result of the previous round.
	// JudgeActionScore represents field Score of the JudgeAction union type.
	JudgeActionScore struct{ Value ScoreCard } // The result of the game.
	// __JudgeActionReflect describes the JudgeAction union type.
	__JudgeActionReflect struct {
		Name  string `vdl:"v.io/x/ref/examples/rps.JudgeAction"`
		Type  JudgeAction
		Union struct {
			PlayerNum    JudgeActionPlayerNum
			OpponentName JudgeActionOpponentName
			MoveOptions  JudgeActionMoveOptions
			RoundResult  JudgeActionRoundResult
			Score        JudgeActionScore
		}
	}
)

func (x JudgeActionPlayerNum) Index() int                        { return 0 }
func (x JudgeActionPlayerNum) Interface() interface{}            { return x.Value }
func (x JudgeActionPlayerNum) Name() string                      { return "PlayerNum" }
func (x JudgeActionPlayerNum) __VDLReflect(__JudgeActionReflect) {}

func (x JudgeActionOpponentName) Index() int                        { return 1 }
func (x JudgeActionOpponentName) Interface() interface{}            { return x.Value }
func (x JudgeActionOpponentName) Name() string                      { return "OpponentName" }
func (x JudgeActionOpponentName) __VDLReflect(__JudgeActionReflect) {}

func (x JudgeActionMoveOptions) Index() int                        { return 2 }
func (x JudgeActionMoveOptions) Interface() interface{}            { return x.Value }
func (x JudgeActionMoveOptions) Name() string                      { return "MoveOptions" }
func (x JudgeActionMoveOptions) __VDLReflect(__JudgeActionReflect) {}

func (x JudgeActionRoundResult) Index() int                        { return 3 }
func (x JudgeActionRoundResult) Interface() interface{}            { return x.Value }
func (x JudgeActionRoundResult) Name() string                      { return "RoundResult" }
func (x JudgeActionRoundResult) __VDLReflect(__JudgeActionReflect) {}

func (x JudgeActionScore) Index() int                        { return 4 }
func (x JudgeActionScore) Interface() interface{}            { return x.Value }
func (x JudgeActionScore) Name() string                      { return "Score" }
func (x JudgeActionScore) __VDLReflect(__JudgeActionReflect) {}

type PlayersMoves [2]string

func (PlayersMoves) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/examples/rps.PlayersMoves"`
}) {
}

// Round represents the state of a round.
type Round struct {
	Moves     PlayersMoves // Each player's move.
	Comment   string       // A text comment from judge about the round.
	Winner    WinnerTag    // Who won the round.
	StartTime time.Time    // The time at which the round started.
	EndTime   time.Time    // The time at which the round ended.
}

func (Round) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/examples/rps.Round"`
}) {
}

// WinnerTag is a type used to indicate whether a round or a game was a draw,
// was won by player 1 or was won by player 2.
type WinnerTag byte

func (WinnerTag) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/examples/rps.WinnerTag"`
}) {
}

// PlayResult is the value returned by the Play method. It indicates the outcome of the game.
type PlayResult struct {
	YouWon bool // True if the player receiving the result won the game.
}

func (PlayResult) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/examples/rps.PlayResult"`
}) {
}

type ScoreCard struct {
	Opts      GameOptions // The game options.
	Judge     string      // The name of the judge.
	Players   []string    // The name of the players.
	Rounds    []Round     // The outcome of each round.
	StartTime time.Time   // The time at which the game started.
	EndTime   time.Time   // The time at which the game ended.
	Winner    WinnerTag   // Who won the game.
}

func (ScoreCard) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/examples/rps.ScoreCard"`
}) {
}

func init() {
	vdl.Register((*GameId)(nil))
	vdl.Register((*GameOptions)(nil))
	vdl.Register((*GameTypeTag)(nil))
	vdl.Register((*PlayerAction)(nil))
	vdl.Register((*unused)(nil))
	vdl.Register((*JudgeAction)(nil))
	vdl.Register((*PlayersMoves)(nil))
	vdl.Register((*Round)(nil))
	vdl.Register((*WinnerTag)(nil))
	vdl.Register((*PlayResult)(nil))
	vdl.Register((*ScoreCard)(nil))
}

const Classic = GameTypeTag(0) // Rock-Paper-Scissors

const LizardSpock = GameTypeTag(1) // Rock-Paper-Scissors-Lizard-Spock

const Draw = WinnerTag(0)

const Player1 = WinnerTag(1)

const Player2 = WinnerTag(2)

// JudgeClientMethods is the client interface
// containing Judge methods.
type JudgeClientMethods interface {
	// CreateGame creates a new game with the given game options and returns a game
	// identifier that can be used by the players to join the game.
	CreateGame(ctx *context.T, Opts GameOptions, opts ...rpc.CallOpt) (GameId, error)
	// Play lets a player join an existing game and play.
	Play(ctx *context.T, Id GameId, opts ...rpc.CallOpt) (JudgePlayClientCall, error)
}

// JudgeClientStub adds universal methods to JudgeClientMethods.
type JudgeClientStub interface {
	JudgeClientMethods
	rpc.UniversalServiceMethods
}

// JudgeClient returns a client stub for Judge.
func JudgeClient(name string) JudgeClientStub {
	return implJudgeClientStub{name}
}

type implJudgeClientStub struct {
	name string
}

func (c implJudgeClientStub) CreateGame(ctx *context.T, i0 GameOptions, opts ...rpc.CallOpt) (o0 GameId, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "CreateGame", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implJudgeClientStub) Play(ctx *context.T, i0 GameId, opts ...rpc.CallOpt) (ocall JudgePlayClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Play", []interface{}{i0}, opts...); err != nil {
		return
	}
	ocall = &implJudgePlayClientCall{ClientCall: call}
	return
}

// JudgePlayClientStream is the client stream for Judge.Play.
type JudgePlayClientStream interface {
	// RecvStream returns the receiver side of the Judge.Play client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() JudgeAction
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Judge.Play client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors
		// encountered while sending, or if Send is called after Close or
		// the stream has been canceled.  Blocks if there is no buffer
		// space; will unblock when buffer space is available or after
		// the stream has been canceled.
		Send(item PlayerAction) error
		// Close indicates to the server that no more items will be sent;
		// server Recv calls will receive io.EOF after all sent items.
		// This is an optional call - e.g. a client might call Close if it
		// needs to continue receiving items from the server after it's
		// done sending.  Returns errors encountered while closing, or if
		// Close is called after the stream has been canceled.  Like Send,
		// blocks if there is no buffer space available.
		Close() error
	}
}

// JudgePlayClientCall represents the call returned from Judge.Play.
type JudgePlayClientCall interface {
	JudgePlayClientStream
	// Finish performs the equivalent of SendStream().Close, then blocks until
	// the server is done, and returns the positional return values for the call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() (PlayResult, error)
}

type implJudgePlayClientCall struct {
	rpc.ClientCall
	valRecv JudgeAction
	errRecv error
}

func (c *implJudgePlayClientCall) RecvStream() interface {
	Advance() bool
	Value() JudgeAction
	Err() error
} {
	return implJudgePlayClientCallRecv{c}
}

type implJudgePlayClientCallRecv struct {
	c *implJudgePlayClientCall
}

func (c implJudgePlayClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implJudgePlayClientCallRecv) Value() JudgeAction {
	return c.c.valRecv
}
func (c implJudgePlayClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implJudgePlayClientCall) SendStream() interface {
	Send(item PlayerAction) error
	Close() error
} {
	return implJudgePlayClientCallSend{c}
}

type implJudgePlayClientCallSend struct {
	c *implJudgePlayClientCall
}

func (c implJudgePlayClientCallSend) Send(item PlayerAction) error {
	return c.c.Send(item)
}
func (c implJudgePlayClientCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implJudgePlayClientCall) Finish() (o0 PlayResult, err error) {
	err = c.ClientCall.Finish(&o0)
	return
}

// JudgeServerMethods is the interface a server writer
// implements for Judge.
type JudgeServerMethods interface {
	// CreateGame creates a new game with the given game options and returns a game
	// identifier that can be used by the players to join the game.
	CreateGame(ctx *context.T, call rpc.ServerCall, Opts GameOptions) (GameId, error)
	// Play lets a player join an existing game and play.
	Play(ctx *context.T, call JudgePlayServerCall, Id GameId) (PlayResult, error)
}

// JudgeServerStubMethods is the server interface containing
// Judge methods, as expected by rpc.Server.
// The only difference between this interface and JudgeServerMethods
// is the streaming methods.
type JudgeServerStubMethods interface {
	// CreateGame creates a new game with the given game options and returns a game
	// identifier that can be used by the players to join the game.
	CreateGame(ctx *context.T, call rpc.ServerCall, Opts GameOptions) (GameId, error)
	// Play lets a player join an existing game and play.
	Play(ctx *context.T, call *JudgePlayServerCallStub, Id GameId) (PlayResult, error)
}

// JudgeServerStub adds universal methods to JudgeServerStubMethods.
type JudgeServerStub interface {
	JudgeServerStubMethods
	// Describe the Judge interfaces.
	Describe__() []rpc.InterfaceDesc
}

// JudgeServer returns a server stub for Judge.
// It converts an implementation of JudgeServerMethods into
// an object that may be used by rpc.Server.
func JudgeServer(impl JudgeServerMethods) JudgeServerStub {
	stub := implJudgeServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implJudgeServerStub struct {
	impl JudgeServerMethods
	gs   *rpc.GlobState
}

func (s implJudgeServerStub) CreateGame(ctx *context.T, call rpc.ServerCall, i0 GameOptions) (GameId, error) {
	return s.impl.CreateGame(ctx, call, i0)
}

func (s implJudgeServerStub) Play(ctx *context.T, call *JudgePlayServerCallStub, i0 GameId) (PlayResult, error) {
	return s.impl.Play(ctx, call, i0)
}

func (s implJudgeServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implJudgeServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{JudgeDesc}
}

// JudgeDesc describes the Judge interface.
var JudgeDesc rpc.InterfaceDesc = descJudge

// descJudge hides the desc to keep godoc clean.
var descJudge = rpc.InterfaceDesc{
	Name:    "Judge",
	PkgPath: "v.io/x/ref/examples/rps",
	Methods: []rpc.MethodDesc{
		{
			Name: "CreateGame",
			Doc:  "// CreateGame creates a new game with the given game options and returns a game\n// identifier that can be used by the players to join the game.",
			InArgs: []rpc.ArgDesc{
				{"Opts", ``}, // GameOptions
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // GameId
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "Play",
			Doc:  "// Play lets a player join an existing game and play.",
			InArgs: []rpc.ArgDesc{
				{"Id", ``}, // GameId
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // PlayResult
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
	},
}

// JudgePlayServerStream is the server stream for Judge.Play.
type JudgePlayServerStream interface {
	// RecvStream returns the receiver side of the Judge.Play server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() PlayerAction
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Judge.Play server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item JudgeAction) error
	}
}

// JudgePlayServerCall represents the context passed to Judge.Play.
type JudgePlayServerCall interface {
	rpc.ServerCall
	JudgePlayServerStream
}

// JudgePlayServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements JudgePlayServerCall.
type JudgePlayServerCallStub struct {
	rpc.StreamServerCall
	valRecv PlayerAction
	errRecv error
}

// Init initializes JudgePlayServerCallStub from rpc.StreamServerCall.
func (s *JudgePlayServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// RecvStream returns the receiver side of the Judge.Play server stream.
func (s *JudgePlayServerCallStub) RecvStream() interface {
	Advance() bool
	Value() PlayerAction
	Err() error
} {
	return implJudgePlayServerCallRecv{s}
}

type implJudgePlayServerCallRecv struct {
	s *JudgePlayServerCallStub
}

func (s implJudgePlayServerCallRecv) Advance() bool {
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implJudgePlayServerCallRecv) Value() PlayerAction {
	return s.s.valRecv
}
func (s implJudgePlayServerCallRecv) Err() error {
	if s.s.errRecv == io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Judge.Play server stream.
func (s *JudgePlayServerCallStub) SendStream() interface {
	Send(item JudgeAction) error
} {
	return implJudgePlayServerCallSend{s}
}

type implJudgePlayServerCallSend struct {
	s *JudgePlayServerCallStub
}

func (s implJudgePlayServerCallSend) Send(item JudgeAction) error {
	return s.s.Send(item)
}

// PlayerClientMethods is the client interface
// containing Player methods.
//
// Player can receive challenges from other players.
type PlayerClientMethods interface {
	// Challenge is used by other players to challenge this player to a game. If
	// the challenge is accepted, the method returns nil.
	Challenge(ctx *context.T, Address string, Id GameId, Opts GameOptions, opts ...rpc.CallOpt) error
}

// PlayerClientStub adds universal methods to PlayerClientMethods.
type PlayerClientStub interface {
	PlayerClientMethods
	rpc.UniversalServiceMethods
}

// PlayerClient returns a client stub for Player.
func PlayerClient(name string) PlayerClientStub {
	return implPlayerClientStub{name}
}

type implPlayerClientStub struct {
	name string
}

func (c implPlayerClientStub) Challenge(ctx *context.T, i0 string, i1 GameId, i2 GameOptions, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Challenge", []interface{}{i0, i1, i2}, nil, opts...)
	return
}

// PlayerServerMethods is the interface a server writer
// implements for Player.
//
// Player can receive challenges from other players.
type PlayerServerMethods interface {
	// Challenge is used by other players to challenge this player to a game. If
	// the challenge is accepted, the method returns nil.
	Challenge(ctx *context.T, call rpc.ServerCall, Address string, Id GameId, Opts GameOptions) error
}

// PlayerServerStubMethods is the server interface containing
// Player methods, as expected by rpc.Server.
// There is no difference between this interface and PlayerServerMethods
// since there are no streaming methods.
type PlayerServerStubMethods PlayerServerMethods

// PlayerServerStub adds universal methods to PlayerServerStubMethods.
type PlayerServerStub interface {
	PlayerServerStubMethods
	// Describe the Player interfaces.
	Describe__() []rpc.InterfaceDesc
}

// PlayerServer returns a server stub for Player.
// It converts an implementation of PlayerServerMethods into
// an object that may be used by rpc.Server.
func PlayerServer(impl PlayerServerMethods) PlayerServerStub {
	stub := implPlayerServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implPlayerServerStub struct {
	impl PlayerServerMethods
	gs   *rpc.GlobState
}

func (s implPlayerServerStub) Challenge(ctx *context.T, call rpc.ServerCall, i0 string, i1 GameId, i2 GameOptions) error {
	return s.impl.Challenge(ctx, call, i0, i1, i2)
}

func (s implPlayerServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implPlayerServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{PlayerDesc}
}

// PlayerDesc describes the Player interface.
var PlayerDesc rpc.InterfaceDesc = descPlayer

// descPlayer hides the desc to keep godoc clean.
var descPlayer = rpc.InterfaceDesc{
	Name:    "Player",
	PkgPath: "v.io/x/ref/examples/rps",
	Doc:     "// Player can receive challenges from other players.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Challenge",
			Doc:  "// Challenge is used by other players to challenge this player to a game. If\n// the challenge is accepted, the method returns nil.",
			InArgs: []rpc.ArgDesc{
				{"Address", ``}, // string
				{"Id", ``},      // GameId
				{"Opts", ``},    // GameOptions
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
	},
}

// ScoreKeeperClientMethods is the client interface
// containing ScoreKeeper methods.
//
// ScoreKeeper receives the outcome of games from Judges.
type ScoreKeeperClientMethods interface {
	Record(ctx *context.T, Score ScoreCard, opts ...rpc.CallOpt) error
}

// ScoreKeeperClientStub adds universal methods to ScoreKeeperClientMethods.
type ScoreKeeperClientStub interface {
	ScoreKeeperClientMethods
	rpc.UniversalServiceMethods
}

// ScoreKeeperClient returns a client stub for ScoreKeeper.
func ScoreKeeperClient(name string) ScoreKeeperClientStub {
	return implScoreKeeperClientStub{name}
}

type implScoreKeeperClientStub struct {
	name string
}

func (c implScoreKeeperClientStub) Record(ctx *context.T, i0 ScoreCard, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Record", []interface{}{i0}, nil, opts...)
	return
}

// ScoreKeeperServerMethods is the interface a server writer
// implements for ScoreKeeper.
//
// ScoreKeeper receives the outcome of games from Judges.
type ScoreKeeperServerMethods interface {
	Record(ctx *context.T, call rpc.ServerCall, Score ScoreCard) error
}

// ScoreKeeperServerStubMethods is the server interface containing
// ScoreKeeper methods, as expected by rpc.Server.
// There is no difference between this interface and ScoreKeeperServerMethods
// since there are no streaming methods.
type ScoreKeeperServerStubMethods ScoreKeeperServerMethods

// ScoreKeeperServerStub adds universal methods to ScoreKeeperServerStubMethods.
type ScoreKeeperServerStub interface {
	ScoreKeeperServerStubMethods
	// Describe the ScoreKeeper interfaces.
	Describe__() []rpc.InterfaceDesc
}

// ScoreKeeperServer returns a server stub for ScoreKeeper.
// It converts an implementation of ScoreKeeperServerMethods into
// an object that may be used by rpc.Server.
func ScoreKeeperServer(impl ScoreKeeperServerMethods) ScoreKeeperServerStub {
	stub := implScoreKeeperServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implScoreKeeperServerStub struct {
	impl ScoreKeeperServerMethods
	gs   *rpc.GlobState
}

func (s implScoreKeeperServerStub) Record(ctx *context.T, call rpc.ServerCall, i0 ScoreCard) error {
	return s.impl.Record(ctx, call, i0)
}

func (s implScoreKeeperServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implScoreKeeperServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{ScoreKeeperDesc}
}

// ScoreKeeperDesc describes the ScoreKeeper interface.
var ScoreKeeperDesc rpc.InterfaceDesc = descScoreKeeper

// descScoreKeeper hides the desc to keep godoc clean.
var descScoreKeeper = rpc.InterfaceDesc{
	Name:    "ScoreKeeper",
	PkgPath: "v.io/x/ref/examples/rps",
	Doc:     "// ScoreKeeper receives the outcome of games from Judges.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Record",
			InArgs: []rpc.ArgDesc{
				{"Score", ``}, // ScoreCard
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
	},
}

// RockPaperScissorsClientMethods is the client interface
// containing RockPaperScissors methods.
type RockPaperScissorsClientMethods interface {
	JudgeClientMethods
	// Player can receive challenges from other players.
	PlayerClientMethods
	// ScoreKeeper receives the outcome of games from Judges.
	ScoreKeeperClientMethods
}

// RockPaperScissorsClientStub adds universal methods to RockPaperScissorsClientMethods.
type RockPaperScissorsClientStub interface {
	RockPaperScissorsClientMethods
	rpc.UniversalServiceMethods
}

// RockPaperScissorsClient returns a client stub for RockPaperScissors.
func RockPaperScissorsClient(name string) RockPaperScissorsClientStub {
	return implRockPaperScissorsClientStub{name, JudgeClient(name), PlayerClient(name), ScoreKeeperClient(name)}
}

type implRockPaperScissorsClientStub struct {
	name string

	JudgeClientStub
	PlayerClientStub
	ScoreKeeperClientStub
}

// RockPaperScissorsServerMethods is the interface a server writer
// implements for RockPaperScissors.
type RockPaperScissorsServerMethods interface {
	JudgeServerMethods
	// Player can receive challenges from other players.
	PlayerServerMethods
	// ScoreKeeper receives the outcome of games from Judges.
	ScoreKeeperServerMethods
}

// RockPaperScissorsServerStubMethods is the server interface containing
// RockPaperScissors methods, as expected by rpc.Server.
// The only difference between this interface and RockPaperScissorsServerMethods
// is the streaming methods.
type RockPaperScissorsServerStubMethods interface {
	JudgeServerStubMethods
	// Player can receive challenges from other players.
	PlayerServerStubMethods
	// ScoreKeeper receives the outcome of games from Judges.
	ScoreKeeperServerStubMethods
}

// RockPaperScissorsServerStub adds universal methods to RockPaperScissorsServerStubMethods.
type RockPaperScissorsServerStub interface {
	RockPaperScissorsServerStubMethods
	// Describe the RockPaperScissors interfaces.
	Describe__() []rpc.InterfaceDesc
}

// RockPaperScissorsServer returns a server stub for RockPaperScissors.
// It converts an implementation of RockPaperScissorsServerMethods into
// an object that may be used by rpc.Server.
func RockPaperScissorsServer(impl RockPaperScissorsServerMethods) RockPaperScissorsServerStub {
	stub := implRockPaperScissorsServerStub{
		impl:                  impl,
		JudgeServerStub:       JudgeServer(impl),
		PlayerServerStub:      PlayerServer(impl),
		ScoreKeeperServerStub: ScoreKeeperServer(impl),
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implRockPaperScissorsServerStub struct {
	impl RockPaperScissorsServerMethods
	JudgeServerStub
	PlayerServerStub
	ScoreKeeperServerStub
	gs *rpc.GlobState
}

func (s implRockPaperScissorsServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implRockPaperScissorsServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{RockPaperScissorsDesc, JudgeDesc, PlayerDesc, ScoreKeeperDesc}
}

// RockPaperScissorsDesc describes the RockPaperScissors interface.
var RockPaperScissorsDesc rpc.InterfaceDesc = descRockPaperScissors

// descRockPaperScissors hides the desc to keep godoc clean.
var descRockPaperScissors = rpc.InterfaceDesc{
	Name:    "RockPaperScissors",
	PkgPath: "v.io/x/ref/examples/rps",
	Embeds: []rpc.EmbedDesc{
		{"Judge", "v.io/x/ref/examples/rps", ``},
		{"Player", "v.io/x/ref/examples/rps", "// Player can receive challenges from other players."},
		{"ScoreKeeper", "v.io/x/ref/examples/rps", "// ScoreKeeper receives the outcome of games from Judges."},
	},
}
