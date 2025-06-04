package api

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

const (
	ProtocolMagic      uint16 = 0xF6D5
	ProtocolVersion    uint8  = 1
	TerminatorByte     uint8  = 0x1
	MessagePollState   uint8  = 0x0
	MessageServerState uint8  = 0x1
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

type ServerStateResponse struct {
	Cookie       uint64
	ServerState  uint8
	ServerNetCL  uint32
	ServerFlags  uint64
	NumSubStates uint8
}

func BuildPollServerStateEnvelope(cookie uint64) ([]byte, error) {
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
			// Got response
			return buffer[:respLen], nil
		}

		// Log and wait before retrying
		fmt.Println("Retry", i+1, "due to error:", err)
		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("no response after %d retries", maxRetries)
}

func ParseServerStateResponse(data []byte) (*ServerStateResponse, error) {
	if len(data) < 22 { // min size of fixed fields
		return nil, fmt.Errorf("response too short")
	}

	r := bytes.NewReader(data)

	var magic uint16
	err := binary.Read(r, binary.LittleEndian, &magic)
	if err != nil {
		return nil, err
	}
	if magic != ProtocolMagic {
		return nil, fmt.Errorf("invalid magic: got %x", magic)
	}

	var msgType uint8
	err = binary.Read(r, binary.LittleEndian, &msgType)
	if err != nil {
		return nil, err
	}
	if msgType != MessageServerState {
		return nil, fmt.Errorf("unexpected message type: %d", msgType)
	}

	var version uint8
	err = binary.Read(r, binary.LittleEndian, &version)
	if err != nil {
		return nil, err
	}

	var resp ServerStateResponse
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

	return &resp, nil
}
