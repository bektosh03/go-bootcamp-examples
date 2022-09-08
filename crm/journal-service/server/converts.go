package server

import (
	"github.com/bektosh03/crmprotos/journalpb"
	"journal-service/domain/journal"
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
