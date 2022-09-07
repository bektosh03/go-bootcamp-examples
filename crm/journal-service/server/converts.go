package server

import (
	"journal-service/domain/journal"
	journalpb "journal-service/protos"
)

func toProtoJournal(j journal.Journal) *journalpb.Journal {
	return &journalpb.Journal{
		Id:         j.Id().String(),
		ScheduleId: j.ScheduleId().String(),
		StudentId:  j.StudentId().String(),
		Attended:   j.Attended(),
		Mark:       j.Mark(),
	}
}
