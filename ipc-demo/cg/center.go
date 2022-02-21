package cg

import (
	"goyard/ipc-demo/ipc"
	"sync"
)

var _ ipc.Server = &CenterServer{} //确认实现了Server接口

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

type Room struct {
	RoomID int `json:"roomID"`
}

type CenterServer struct {
	servers map[string]ipc.Server
	players []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}

func NewCenterServer() *CenterServer {
	servers := make(map[string]ipc.Server)
	players := make([]*Player, 0)

	return &CenterServer{
		servers: servers,
		players: players,
	}
}

func (server *CenterServer) addPlayer(params string) error {

}
