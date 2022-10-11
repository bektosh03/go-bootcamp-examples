package service

import (
	"context"
	"journal-service/consumer/attend"
	"journal-service/consumer/mark"
	"journal-service/domain/journal"
	"log"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo           JournalRepository
	journalFactory journal.Factory
	attconsumer    attend.Consumer
	markconsumer   mark.Consumer
}

func New(jouRepo JournalRepository, journalFactory journal.Factory, attconsumer attend.Consumer, markconsumer mark.Consumer) Service {
	return Service{
		repo:           jouRepo,
		journalFactory: journalFactory,
		attconsumer:    attconsumer,
		markconsumer:   markconsumer,
	}
}

func (s Service) RunAtt()  {
	go func() {
		att := s.attconsumer.Events()
		s.handleAttendance(att)
	}()

}

func (s Service) RunMark() {
	go func() {
		mark := s.markconsumer.Events()
		s.handleMarks(mark)
	}()
}

func (s Service) handleMarks(events <-chan mark.StudentMarkedEvent) {
	for e := range events {
		journalID, err := uuid.Parse(e.JournalID)
		if err != nil {
			log.Println("handleMarks:error with parsing journalID", err)
			return
		}

		studentID, err := uuid.Parse(e.StudentID)
		if err != nil {
			log.Println("handleMarks:error with parsing studentID", err)
			return
		}

		stat, err := journal.NewStatus(journalID, studentID, true, e.Mark)
		if err != nil {
			log.Println("handleMarks:error with creating new journal tools", err)
			return
		}

		err = s.MarkStudent(context.Background(), stat)
		if err != nil {
			log.Println("handleMarks:error with completing marking student", err)
			return
		}
	}
}

func (s Service) handleAttendance(events <-chan attend.SetStudentAttendanceEvent) {
	for e := range events {
		journalID, err := uuid.Parse(e.JournalID)
		if err != nil {
			log.Println("handleAttendance:error with parsing journalID", err)
			return
		}

		studentID, err := uuid.Parse(e.StudentID)
		if err != nil {
			log.Println("handleAttendance:error with parsing studentID", err)
			return
		}

		stat, err := journal.NewStatus(journalID, studentID, e.Attended, 5)
		if err != nil {
			log.Println("handleAttendance:error with creating new journal tools", err)
			return
		}

		err = s.SetStudentAttendance(context.Background(), stat)
		if err != nil {
			log.Println("handleAttendance:error with completing setting Student Attendance", err)
			return
		}
	}
}

func (s Service) RegisterJournal(ctx context.Context, journal journal.Journal, studentIDs []uuid.UUID) error {
	if err := s.repo.CreateJournal(ctx, journal); err != nil {
		return err
	}
	if err := s.repo.CreateJournalStats(ctx, journal.ID(), studentIDs); err != nil {
		if err = s.repo.DeleteJournal(ctx, journal.ID()); err != nil {
			log.Println("[ERROR] could not delete journal after failed insertion of journal statuses")
		}
		return err
	}

	return nil
}

func (s Service) GetJournal(ctx context.Context, id uuid.UUID) (journal.Journal, error) {
	jour, err := s.repo.GetJournal(ctx, id)
	if err != nil {
		return journal.Journal{}, nil
	}
	return jour, nil
}

func (s Service) UpdateJournal(ctx context.Context, jour journal.Journal) error {
	return s.repo.UpdateJournal(ctx, jour)
}

func (s Service) DeleteJournal(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteJournal(ctx, id)
}

func (s Service) MarkStudent(ctx context.Context, st journal.Stat) error {
	return s.repo.MarkStudent(ctx, st)
}

func (s Service) SetStudentAttendance(ctx context.Context, st journal.Stat) error {
	return s.repo.SetStudentAttendance(ctx, st)
}

func (s Service) GetStudentJournalEntries(ctx context.Context, studentID uuid.UUID, start, end time.Time) ([]journal.Entry, error) {
	return s.repo.GetStudentJournalEntries(ctx, studentID, start, end)
}

func (s Service) GetTeacherJournalEntries(ctx context.Context, teacherID uuid.UUID, start, end time.Time) ([]journal.Entry, error) {
	return s.repo.GetTeacherJournalEntries(ctx, teacherID, start, end)
}
