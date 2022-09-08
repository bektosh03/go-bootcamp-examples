package journal

import (
	"github.com/bektosh03/crmcommon/id"
	"github.com/google/uuid"
)

type Factory struct {
	idGenerator id.Generator
}

func NewFactory(idGenerator id.Generator) Factory {
	return Factory{
		idGenerator: idGenerator,
	}
}
func (f Factory) NewJournal(scheduleId, studentId uuid.UUID, attended bool, mark int32) (Journal, error) {
	j := Journal{
		id:         f.idGenerator.GenerateUUID(),
		scheduleId: scheduleId,
		studentId:  studentId,
		attended:   attended,
		mark:       mark,
	}
	if err := j.validate(); err != nil {
		return Journal{}, err
	}
	return j, nil
}
