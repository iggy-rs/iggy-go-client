package binaryserialization

import (
	"encoding/binary"
	"math"

	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
)

type TcpStats struct {
	iggcon.Stats
}

// Constants for byte positions and lengths in the payload.
const (
	processIDPos           = 0
	cpuUsagePos            = 4
	totalCpuUsagePos       = 8
	memoryUsagePos         = 12
	totalMemoryPos         = 20
	availableMemoryPos     = 28
	runTimePos             = 36
	startTimePos           = 44
	readBytesPos           = 52
	writtenBytesPos        = 60
	messagesSizeBytesPos   = 68
	streamsCountPos        = 76
	topicsCountPos         = 80
	partitionsCountPos     = 84
	segmentsCountPos       = 88
	messagesCountPos       = 92
	clientsCountPos        = 100
	consumerGroupsCountPos = 104
)

func (stats *TcpStats) Deserialize(payload []byte) error {
	stats.ProcessId = int(binary.LittleEndian.Uint32(payload[processIDPos : processIDPos+4]))
	stats.CpuUsage = math.Float32frombits(binary.LittleEndian.Uint32(payload[cpuUsagePos : cpuUsagePos+4]))
	stats.TotalCpuUsage = math.Float32frombits(binary.LittleEndian.Uint32(payload[totalCpuUsagePos : totalCpuUsagePos+4]))
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
