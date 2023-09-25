package binaryserialization

import (
	"encoding/binary"
	"unsafe"

	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
)

type TcpStats struct {
	iggcon.Stats
}

// Constants for byte positions and lengths in the payload.
const (
	processIDPos           = 0
	cpuUsagePos            = 4
	memoryUsagePos         = 8
	totalMemoryPos         = 16
	availableMemoryPos     = 24
	runTimePos             = 32
	startTimePos           = 40
	readBytesPos           = 48
	writtenBytesPos        = 56
	messagesSizeBytesPos   = 64
	streamsCountPos        = 72
	topicsCountPos         = 76
	partitionsCountPos     = 80
	segmentsCountPos       = 84
	messagesCountPos       = 88
	clientsCountPos        = 96
	consumerGroupsCountPos = 100
)

// Deserialize deserializes a byte slice into a TcpStats structure.
// It populates the TcpStats fields by interpreting data from the payload byte slice.
func (stats *TcpStats) Deserialize(payload []byte) error {
	stats.ProcessId = int(binary.LittleEndian.Uint32(payload[processIDPos : processIDPos+4]))
	stats.CpuUsage = *(*float32)(unsafe.Pointer(&payload[cpuUsagePos]))
	stats.MemoryUsage = binary.LittleEndian.Uint64(payload[memoryUsagePos : memoryUsagePos+8])
	stats.TotalMemory = binary.LittleEndian.Uint64(payload[totalMemoryPos : totalMemoryPos+8])
	stats.AvailableMemory = binary.LittleEndian.Uint64(payload[availableMemoryPos : availableMemoryPos+8])
	stats.RunTime = binary.LittleEndian.Uint64(payload[runTimePos : runTimePos+8])
	stats.StartTime = binary.LittleEndian.Uint64(payload[startTimePos : startTimePos+8])
	stats.ReadBytes = binary.LittleEndian.Uint64(payload[readBytesPos : readBytesPos+8])
	stats.WrittenBytes = binary.LittleEndian.Uint64(payload[writtenBytesPos : writtenBytesPos+8])
	stats.MessagesSizeBytes = binary.LittleEndian.Uint64(payload[messagesSizeBytesPos : messagesSizeBytesPos+8])
	stats.StreamsCount = int(binary.LittleEndian.Uint32(payload[streamsCountPos : streamsCountPos+4]))
	stats.TopicsCount = int(binary.LittleEndian.Uint32(payload[topicsCountPos : topicsCountPos+4]))
	stats.PartitionsCount = int(binary.LittleEndian.Uint32(payload[partitionsCountPos : partitionsCountPos+4]))
	stats.SegmentsCount = int(binary.LittleEndian.Uint32(payload[segmentsCountPos : segmentsCountPos+4]))
	stats.MessagesCount = binary.LittleEndian.Uint64(payload[messagesCountPos : messagesCountPos+8])
	stats.ClientsCount = int(binary.LittleEndian.Uint32(payload[clientsCountPos : clientsCountPos+4]))
	stats.ConsumerGroupsCount = int(binary.LittleEndian.Uint32(payload[consumerGroupsCountPos : consumerGroupsCountPos+4]))

	position := consumerGroupsCountPos + 4
	hostnameLength := int(binary.LittleEndian.Uint32(payload[position : position+4]))
	stats.Hostname = string(payload[position+4 : position+4+hostnameLength])
	position += 4 + hostnameLength

	osNameLength := int(binary.LittleEndian.Uint32(payload[position : position+4]))
	stats.OsName = string(payload[position+4 : position+4+osNameLength])
	position += 4 + osNameLength

	osVersionLength := int(binary.LittleEndian.Uint32(payload[position : position+4]))
	stats.OsVersion = string(payload[position+4 : position+4+osVersionLength])
	position += 4 + osVersionLength

	kernelVersionLength := int(binary.LittleEndian.Uint32(payload[position : position+4]))
	stats.KernelVersion = string(payload[position+4 : position+4+kernelVersionLength])

	return nil
}
