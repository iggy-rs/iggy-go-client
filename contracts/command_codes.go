package iggcon

type CommandCode int

const (
	KillCode         CommandCode = 0
	PingCode         CommandCode = 1
	GetStatsCode     CommandCode = 10
	SendMessagesCode CommandCode = 101
	PollMessagesCode CommandCode = 100
	StoreOffsetCode  CommandCode = 121
	GetOffsetCode    CommandCode = 120
	GetStreamCode    CommandCode = 200
	GetStreamsCode   CommandCode = 201
	CreateStreamCode CommandCode = 202
	DeleteStreamCode CommandCode = 203
	GetTopicCode     CommandCode = 300
	GetTopicsCode    CommandCode = 301
	CreateTopicCode  CommandCode = 302
	DeleteTopicCode  CommandCode = 303
	GetGroupCode     CommandCode = 600
	GetGroupsCode    CommandCode = 601
	CreateGroupCode  CommandCode = 602
	DeleteGroupCode  CommandCode = 603
	JoinGroupCode    CommandCode = 604
	LeaveGroupCode   CommandCode = 605
)
