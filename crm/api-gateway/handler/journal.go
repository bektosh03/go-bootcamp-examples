package handler

import (
	"api-gateway/pkg/httperr"
	"api-gateway/pkg/producer"
	"api-gateway/request"
	"api-gateway/response"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h Handler) GetTeacherJournal(w http.ResponseWriter, r *http.Request) {
	teacherID := chi.URLParam(r, "id")
	if teacherID == "" {
		httperr.BadRequest(w, r, "teacher id is not provided")
		return
	}

	start, end, err := extractStartEndTimeQueries(r)
	if err != nil {
		httperr.BadRequest(w, r, err.Error())
		return
	}

	journals, err := h.service.Journal.GetTeacherJournal(context.Background(), teacherID, start, end)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	entries := make([]response.FullJournalEntry, 0, len(journals))
	for _, j := range journals {
		schedule, err := h.service.Schedule.GetSchedule(context.Background(), j.ScheduleID)
		if err != nil {
			httperr.Handle(w, r, err)
			return
		}

		subject, err := h.service.Teacher.GetSubject(context.Background(), schedule.SubjectID)
		if err != nil {
			httperr.Handle(w, r, err)
			return
		}

		teacher, err := h.service.Teacher.GetTeacher(
			context.Background(),
			request.GetTeacherRequest{TeacherID: schedule.TeacherID},
		)

		if err != nil {
			httperr.Handle(w, r, err)
			return
		}

		entries = append(entries, response.FullJournalEntry{
			JournalID: j.JournalID,
			Date:      j.Date,
			Subject:   subject,
			Teacher:   teacher,
			Mark:      j.Mark,
			Attended:  j.Attended,
		})
	}

	render.JSON(w, r, entries)
}

func (h Handler) GetStudentJournal(w http.ResponseWriter, r *http.Request) {
	studentID := chi.URLParam(r, "id")
	if studentID == "" {
		httperr.BadRequest(w, r, "student id is not provided")
		return
	}

	start, end, err := extractStartEndTimeQueries(r)
	if err != nil {
		httperr.BadRequest(w, r, err.Error())
		return
	}

	journals, err := h.service.Journal.GetStudentJournal(context.Background(), studentID, start, end)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	entries := make([]response.FullJournalEntry, 0, len(journals))
	for _, j := range journals {
		schedule, err := h.service.Schedule.GetSchedule(context.Background(), j.ScheduleID)
		if err != nil {
			httperr.Handle(w, r, err)
			return
		}

		subject, err := h.service.Teacher.GetSubject(context.Background(), schedule.SubjectID)
		if err != nil {
			httperr.Handle(w, r, err)
			return
		}

		teacher, err := h.service.Teacher.GetTeacher(
			context.Background(),
			request.GetTeacherRequest{TeacherID: schedule.TeacherID},
		)

		if err != nil {
			httperr.Handle(w, r, err)
			return
		}

		entries = append(entries, response.FullJournalEntry{
			JournalID: j.JournalID,
			Date:      j.Date,
			Subject:   subject,
			Teacher:   teacher,
			Mark:      j.Mark,
			Attended:  j.Attended,
		})
	}

	render.JSON(w, r, entries)
}

func (h Handler) MarkStudent(w http.ResponseWriter, r *http.Request) {
	var req request.MarkStudentRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}

	if err := h.producer.Produce(producer.StudentMarkedEvent{
		Mark:      req.Mark,
		StudentID: req.StudentID,
		JournalID: req.JournalID,
	}); err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, render.M{
		"ok": true,
	})
}

func (h Handler) SetStudentAttendance(w http.ResponseWriter, r *http.Request) {
	var req request.SetStudentAttendanceRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}

	if err := h.service.Journal.SetStudentAttendance(context.Background(), req); err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, render.M{
		"ok": true,
	})
}

func (h Handler) RegisterJournal(w http.ResponseWriter, r *http.Request) {
	var req request.CreateJournalRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}

	schedule, err := h.service.Schedule.GetSchedule(context.Background(), req.ScheduleID)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	studentIDs, err := h.service.Student.GetGroupStudentIDs(context.Background(), schedule.GroupId)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	journal, err := h.service.Journal.RegisterJournal(context.Background(), schedule.ID, schedule.TeacherID, req.Date, studentIDs)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, journal)
}

func (h Handler) GetJournal(w http.ResponseWriter, r *http.Request) {
	journalId := chi.URLParam(r, "id")
	journal, err := h.service.Journal.GetJournal(context.Background(), journalId)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, journal)
}

func (h Handler) DeleteJournal(w http.ResponseWriter, r *http.Request) {
	journalId := chi.URLParam(r, "id")
	err := h.service.Journal.DeleteJournal(context.Background(), journalId)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, render.M{
		"ok": "deleted",
	})
}

func (h Handler) UpdateJournal(w http.ResponseWriter, r *http.Request) {
	var req request.Journal
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}
	res, err := h.service.Journal.UpdateJournal(context.Background(), req)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

func extractStartEndTimeQueries(r *http.Request) (start, end time.Time, err error) {
	startValue := r.URL.Query().Get("start")
	start, err = time.Parse("2006-01-02", startValue)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("start time is invalid")
	}

	endValue := r.URL.Query().Get("end")
	end, err = time.Parse("2006-01-02", endValue)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("end time is invalid")
	}

	return start, end, nil
}
