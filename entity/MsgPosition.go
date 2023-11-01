package entity

type MsgPosition struct {
	ChannelName string
	MsgID       []byte
	MsgGroup    string
	Timestamp   uint64
}
