package iggerr

import (
	"testing"
)

func TestIggyError_Error(t *testing.T) {
	iggyErr := &IggyError{
		Code:    42,
		Message: "test_error",
	}

	expectedErrorString := "42: 'test_error'"
	actualErrorString := iggyErr.Error()

	if expectedErrorString != actualErrorString {
		t.Errorf("Error() method mismatch, expected: %s, got: %s", expectedErrorString, actualErrorString)
	}
}

func TestTranslateErrorCode(t *testing.T) {
	testCases := map[int]string{
		1:    "error",
		2:    "invalid_configuration",
		3:    "invalid_command",
		4:    "invalid_format",
		5:    "feature_unavailable",
		10:   "cannot_create_base_directory",
		51:   "not_connected",
		52:   "request_error",
		100:  "client_not_found",
		101:  "invalid_client_id",
		200:  "io_error",
		201:  "write_error",
		202:  "cannot_parse_utf8",
		203:  "cannot_parse_int",
		204:  "cannot_parse_slice",
		300:  "http_response_error",
		301:  "request_middleware_error",
		302:  "cannot_create_endpoint",
		303:  "cannot_parse_url",
		304:  "invalid_response",
		305:  "empty_response",
		306:  "cannot_parse_address",
		307:  "read_error",
		308:  "connection_error",
		309:  "read_to_end_error",
		1000: "cannot_create_streams_directory",
		1001: "cannot_create_stream_directory",
		1002: "cannot_create_stream_info",
		1003: "cannot_update_stream_info",
		1004: "cannot_open_stream_info",
		1005: "cannot_read_stream_info",
		1006: "cannot_create_stream",
		1007: "cannot_delete_stream",
		1008: "cannot_delete_stream_directory",
		1009: "stream_not_found",
		1010: "stream_already_exists",
		1011: "invalid_stream_name",
		1012: "invalid_stream_id",
		1013: "cannot_read_streams",
		2000: "cannot_create_topics_directory",
		2001: "cannot_create_topic_directory",
		2002: "cannot_create_topic_info",
		2003: "cannot_update_topic_info",
		2004: "cannot_open_topic_info",
		2005: "cannot_read_topic_info",
		2006: "cannot_create_topic",
		2007: "cannot_delete_topic",
		2008: "cannot_delete_topic_directory",
		2009: "cannot_poll_topic",
		2010: "topic_not_found",
		2011: "topic_already_exists",
		2012: "invalid_topic_name",
		2013: "invalid_topic_partitions",
		2014: "invalid_topic_id",
		2015: "cannot_read_topics",
		3000: "cannot_create_partition",
		3001: "cannot_create_partitions_directory",
		3002: "cannot_create_partition_directory",
		3003: "cannot_open_partition_log_file",
		3004: "cannot_read_partitions",
		3005: "cannot_delete_partition",
		3006: "cannot_delete_partition_directory",
		3007: "partition_not_found",
		4000: "segment_not_found",
		4001: "segment_closed",
		4002: "invalid_segment_size",
		4003: "cannot_create_segment_log_file",
		4004: "cannot_create_segment_index_file",
		4005: "cannot_create_segment_time_index_file",
		4006: "cannot_save_messages_to_segment",
		4007: "cannot_save_index_to_segment",
		4008: "cannot_save_time_index_to_segment",
		4009: "invalid_messages_count",
		4010: "cannot_append_message",
		4011: "cannot_read_message",
		4012: "cannot_read_message_timestamp",
		4013: "cannot_read_message_id",
		4014: "cannot_read_message_length",
		4015: "cannot_read_message_payload",
		4016: "too_big_message_payload",
		4017: "too_many_messages",
		4018: "empty_message_payload",
		4019: "invalid_message_payload_length",
		4020: "cannot_read_message_checksum",
		4021: "invalid_message_checksum",
		4022: "invalid_key_value_length",
		4100: "invalid_offset",
		4101: "cannot_read_consumer_offsets",
		5000: "consumer_group_not_found",
		5001: "consumer_group_already_exists",
		5002: "consumer_group_member_not_found",
		5003: "invalid_consumer_group_id",
		5004: "cannot_create_consumer_groups_directory",
		5005: "cannot_read_consumer_groups",
		5006: "cannot_create_consumer_group_info",
		5007: "cannot_delete_consumer_group_info",
	}

	for code, expectedMessage := range testCases {
		actualMessage := TranslateErrorCode(code)
		if actualMessage != expectedMessage {
			t.Errorf("TranslateErrorCode(%d) returned '%s', expected '%s'", code, actualMessage, expectedMessage)
		}
	}
}

func TestMapFromCode(t *testing.T) {
	testCases := map[int]struct {
		expectedCode    int
		expectedMessage string
	}{
		1:    {1, "error"},
		2:    {2, "invalid_configuration"},
		3:    {3, "invalid_command"},
		4:    {4, "invalid_format"},
		5:    {5, "feature_unavailable"},
		10:   {10, "cannot_create_base_directory"},
		51:   {51, "not_connected"},
		52:   {52, "request_error"},
		100:  {100, "client_not_found"},
		101:  {101, "invalid_client_id"},
		200:  {200, "io_error"},
		201:  {201, "write_error"},
		202:  {202, "cannot_parse_utf8"},
		203:  {203, "cannot_parse_int"},
		204:  {204, "cannot_parse_slice"},
		300:  {300, "http_response_error"},
		301:  {301, "request_middleware_error"},
		302:  {302, "cannot_create_endpoint"},
		303:  {303, "cannot_parse_url"},
		304:  {304, "invalid_response"},
		305:  {305, "empty_response"},
		306:  {306, "cannot_parse_address"},
		307:  {307, "read_error"},
		308:  {308, "connection_error"},
		309:  {309, "read_to_end_error"},
		1000: {1000, "cannot_create_streams_directory"},
		1001: {1001, "cannot_create_stream_directory"},
		1002: {1002, "cannot_create_stream_info"},
		1003: {1003, "cannot_update_stream_info"},
		1004: {1004, "cannot_open_stream_info"},
		1005: {1005, "cannot_read_stream_info"},
		1006: {1006, "cannot_create_stream"},
		1007: {1007, "cannot_delete_stream"},
		1008: {1008, "cannot_delete_stream_directory"},
		1009: {1009, "stream_not_found"},
		1010: {1010, "stream_already_exists"},
		1011: {1011, "invalid_stream_name"},
		1012: {1012, "invalid_stream_id"},
		1013: {1013, "cannot_read_streams"},
		2000: {2000, "cannot_create_topics_directory"},
		2001: {2001, "cannot_create_topic_directory"},
		2002: {2002, "cannot_create_topic_info"},
		2003: {2003, "cannot_update_topic_info"},
		2004: {2004, "cannot_open_topic_info"},
		2005: {2005, "cannot_read_topic_info"},
		2006: {2006, "cannot_create_topic"},
		2007: {2007, "cannot_delete_topic"},
		2008: {2008, "cannot_delete_topic_directory"},
		2009: {2009, "cannot_poll_topic"},
		2010: {2010, "topic_not_found"},
		2011: {2011, "topic_already_exists"},
		2012: {2012, "invalid_topic_name"},
		2013: {2013, "invalid_topic_partitions"},
		2014: {2014, "invalid_topic_id"},
		2015: {2015, "cannot_read_topics"},
		3000: {3000, "cannot_create_partition"},
		3001: {3001, "cannot_create_partitions_directory"},
		3002: {3002, "cannot_create_partition_directory"},
		3003: {3003, "cannot_open_partition_log_file"},
		3004: {3004, "cannot_read_partitions"},
		3005: {3005, "cannot_delete_partition"},
		3006: {3006, "cannot_delete_partition_directory"},
		3007: {3007, "partition_not_found"},
		4000: {4000, "segment_not_found"},
		4001: {4001, "segment_closed"},
		4002: {4002, "invalid_segment_size"},
		4003: {4003, "cannot_create_segment_log_file"},
		4004: {4004, "cannot_create_segment_index_file"},
		4005: {4005, "cannot_create_segment_time_index_file"},
		4006: {4006, "cannot_save_messages_to_segment"},
		4007: {4007, "cannot_save_index_to_segment"},
		4008: {4008, "cannot_save_time_index_to_segment"},
		4009: {4009, "invalid_messages_count"},
		4010: {4010, "cannot_append_message"},
		4011: {4011, "cannot_read_message"},
		4012: {4012, "cannot_read_message_timestamp"},
		4013: {4013, "cannot_read_message_id"},
		4014: {4014, "cannot_read_message_length"},
		4015: {4015, "cannot_read_message_payload"},
		4016: {4016, "too_big_message_payload"},
		4017: {4017, "too_many_messages"},
		4018: {4018, "empty_message_payload"},
		4019: {4019, "invalid_message_payload_length"},
		4020: {4020, "cannot_read_message_checksum"},
		4021: {4021, "invalid_message_checksum"},
		4022: {4022, "invalid_key_value_length"},
		4100: {4100, "invalid_offset"},
		4101: {4101, "cannot_read_consumer_offsets"},
		5000: {5000, "consumer_group_not_found"},
		5001: {5001, "consumer_group_already_exists"},
		5002: {5002, "consumer_group_member_not_found"},
		5003: {5003, "invalid_consumer_group_id"},
		5004: {5004, "cannot_create_consumer_groups_directory"},
		5005: {5005, "cannot_read_consumer_groups"},
		5006: {5006, "cannot_create_consumer_group_info"},
		5007: {5007, "cannot_delete_consumer_group_info"},
	}

	for code, expected := range testCases {
		err := MapFromCode(code)
		iggyErr, ok := err.(*IggyError)
		if !ok {
			t.Errorf("MapFromCode(%d) did not return an *IggyError", code)
			continue
		}

		if iggyErr.Code != expected.expectedCode {
			t.Errorf("MapFromCode(%d) returned an *IggyError with code %d, expected %d", code, iggyErr.Code, expected.expectedCode)
		}

		if iggyErr.Message != expected.expectedMessage {
			t.Errorf("MapFromCode(%d) returned an *IggyError with message '%s', expected '%s'", code, iggyErr.Message, expected.expectedMessage)
		}
	}
}
