package server

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"journal-service/domain/journal"

	"github.com/bektosh03/crmprotos/journalpb"
)

func toProtoJournal(j journal.Journal) *journalpb.Journal {
	return &journalpb.Journal{
		Id:         j.ID().String(),
		ScheduleId: j.ScheduleID().String(),
		Date: &timestamppb.Timestamp{
			Seconds: j.Date().Unix(),
			Nanos:   int32(j.Date().UnixNano()),
		},
	}
}

func toProtoEntry(entry journal.Entry) *journalpb.Entry {
	return &journalpb.Entry{
		JournalId:  entry.JournalID().String(),
		ScheduleId: entry.ScheduleID().String(),
		StudentId:  entry.StudentID().String(),
		Attended:   entry.Attended(),
		Mark:       entry.Mark(),
		Date: &timestamppb.Timestamp{
			Seconds: entry.Date().Unix(),
			Nanos:   int32(entry.Date().Nanosecond()),
		},
	}
}

func toProtoEntries(entries []journal.Entry) *journalpb.Entries {
	protoEntries := make([]*journalpb.Entry, 0, len(entries))
	for _, e := range entries {
		protoEntries = append(protoEntries, toProtoEntry(e))
	}

	return &journalpb.Entries{
		Entries: protoEntries,
	}
}
