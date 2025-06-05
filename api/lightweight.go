package api

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

const (
	ProtocolMagic        uint16 = 0xF6D5
	ProtocolVersion      uint8  = 1
	TerminatorByte       uint8  = 0x1
	MessagePollState     uint8  = 0x0
	MessageStateResponse uint8  = 0x1
)

type Envelope struct {
	Magic           uint16
	MessageType     uint8
	ProtocolVersion uint8
	Payload         []byte
	Terminator      uint8
}

type PollServerState struct {
	Cookie uint64
}

func CheckMagicPacket(r *bytes.Reader) error {
	var magic uint16
	err := binary.Read(r, binary.LittleEndian, &magic)
	if err != nil {
		return nil
	}
	if magic != ProtocolMagic {
		return fmt.Errorf("invalid magic packet, expected %v, got %v", ProtocolMagic, magic)
	}
	return nil
}

func CheckMessageType(r *bytes.Reader, t uint8) error {
	var messageType uint8
	err := binary.Read(r, binary.LittleEndian, &messageType)
	if err != nil {
		return nil
	}
	if messageType != t {
		return fmt.Errorf("invalid message type, expected %v, got %v", t, messageType)
	}
	return nil
}

func CheckVersion(r *bytes.Reader) error {
	var version uint8
	err := binary.Read(r, binary.LittleEndian, &version)
	if err != nil {
		return err
	}
	if version != ProtocolVersion {
		return fmt.Errorf("invalid protocol version, expected %v, got %v", ProtocolVersion, version)
	}
	return nil
}

func BuildEnvelope(cookie uint64) ([]byte, error) {
	var buf bytes.Buffer

	err := binary.Write(&buf, binary.LittleEndian, ProtocolMagic)
	if err != nil {
		return nil, err
	}

	buf.WriteByte(MessagePollState)
	buf.WriteByte(ProtocolVersion)

	err = binary.Write(&buf, binary.LittleEndian, cookie)
	if err != nil {
		return nil, err
	}

	buf.WriteByte(TerminatorByte)

	return buf.Bytes(), nil
}

func SendUDPQuery(server string, request []byte, maxRetries int, retryDelay time.Duration) ([]byte, error) {
	addr, err := net.ResolveUDPAddr("udp", server)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	var respLen int

	for i := 0; i < maxRetries; i++ {
		_, err = conn.Write(request)
		if err != nil {
			return nil, err
		}

		err := conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		if err != nil {
			return nil, err
		}

		respLen, _, err = conn.ReadFromUDP(buffer)
		if err == nil {
			return buffer[:respLen], nil
		}

		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("no response after %d retries", maxRetries)
}

type ServerState int
type ServerSubStateId int

const (
	ServerStateOffline ServerState = iota
	ServerStateIdle
	ServerStateLoading
	ServerStatePlaying
)

var ServerStateMap = map[ServerState]string{
	ServerStateOffline: "Offline",
	ServerStateIdle:    "Idle",
	ServerStateLoading: "Loading",
	ServerStatePlaying: "Playing",
}

type ServerStateResponse struct {
	Cookie           uint64
	ServerState      uint8
	ServerNetCL      uint32
	ServerFlags      uint64
	NumSubStates     uint8
	SubStates        []ServerSubState
	ServerNameLength uint16
	ServerName       []byte
}

type ServerSubState struct {
	SubStateId      uint8
	SubStateVersion uint16
}

func (state ServerState) String() string {
	return ServerStateMap[state]

}

func ParseServerStateResponse(data []byte) (*ServerStateResponse, error) {
	if len(data) < 22 { // min size of fixed fields
		return nil, fmt.Errorf("response too short")
	}

	r := bytes.NewReader(data)

	err := CheckMagicPacket(r)
	if err != nil {
		return nil, err
	}

	err = CheckMessageType(r, MessageStateResponse)
	if err != nil {
		return nil, err
	}

	err = CheckVersion(r)
	if err != nil {
		return nil, err
	}

	var resp ServerStateResponse

	/*
		I really don't like the line by line reading of the offset to each struct member.
		Will need to figure out a way to seek to offset 21 where substate number is and then
		populate the struct. Binary requires fixed sizes if you want to do it all at once,
		hence the annoying line by line.
	*/

	err = binary.Read(r, binary.LittleEndian, &resp.Cookie)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &resp.ServerState)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &resp.ServerNetCL)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &resp.ServerFlags)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &resp.NumSubStates)
	if err != nil {
		return nil, err
	}

	subStates := make([]ServerSubState, resp.NumSubStates)

	for i := 0; i < int(resp.NumSubStates); i++ {
		var subStateId uint8
		var subStateVersion uint16

		err = binary.Read(r, binary.LittleEndian, &subStateId)
		if err != nil {
			return nil, err
		}

		err = binary.Read(r, binary.LittleEndian, &subStateVersion)
		if err != nil {
			return nil, err
		}

		subStates[i] = ServerSubState{
			SubStateId:      subStateId,
			SubStateVersion: subStateVersion,
		}
	}

	resp.SubStates = subStates

	err = binary.Read(r, binary.LittleEndian, &resp.ServerNameLength)
	if err != nil {
		return nil, err
	}

	resp.ServerName = make([]byte, resp.ServerNameLength+1)

	for {
		_, err := r.Read(resp.ServerName)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
	}

	return &resp, nil
}
