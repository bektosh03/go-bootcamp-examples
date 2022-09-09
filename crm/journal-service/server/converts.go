package server

import (
	"journal-service/domain/journal"

	"github.com/bektosh03/crmprotos/journalpb"
)

func toProtoJournal(j journal.Journal) *journalpb.Journal {
	return &journalpb.Journal{
		Id:         j.ID().String(),
		ScheduleId: j.ScheduleID().String(),
	}
}
