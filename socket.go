package netlink

import "net"

// Socket States
const (
	SS_FREE = iota
	SS_UNCONNECTED
	SS_CONNECTING
	SS_CONNECTED
	SS_DISCONNECTING
)

const (
	SS_UNKNOWN = iota
	SS_ESTABLISHED
	SS_SYN_SENT
	SS_SYN_RECV
	SS_FIN_WAIT1
	SS_FIN_WAIT2
	SS_TIME_WAIT
	SS_CLOSE
	SS_CLOSE_WAIT
	SS_LAST_ACK
	SS_LISTEN
	SS_CLOSING
	SS_MAX
)

// SocketID identifies a single socket.
type SocketID struct {
	SourcePort      uint16
	DestinationPort uint16
	Source          net.IP
	Destination     net.IP
	Interface       uint32
	Cookie          [2]uint32
}

// Socket represents a netlink socket.
type Socket struct {
	Family  uint8
	State   uint8
	Timer   uint8
	Retrans uint8
	ID      SocketID
	Expires uint32
	RQueue  uint32
	WQueue  uint32
	UID     uint32
	INode   uint32
}
