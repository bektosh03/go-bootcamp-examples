package handler

import (
	"api-gateway/pkg/httperr"
	"api-gateway/request"
	"api-gateway/response"
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h Handler) GetStudentJournal(w http.ResponseWriter, r *http.Request) {
	studentID := chi.URLParam(r, "id")
	if studentID == "" {
		httperr.BadRequest(w, r, "student id is not provided")
		return
	}
	startValue := r.URL.Query().Get("start")
	start, err := time.Parse("2006-01-02", startValue)
	if err != nil {
		httperr.BadRequest(w, r, "start time is not valid")
		return
	}

	endValue := r.URL.Query().Get("end")
	end, err := time.Parse("2006-01-02", endValue)
	if err != nil {
		httperr.BadRequest(w, r, "end time is not valid")
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

	if err := h.service.Journal.MarkStudent(context.Background(), req); err != nil {
		httperr.Handle(w, r, err)
		return
	}

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

	journal, err := h.service.Journal.RegisterJournal(context.Background(), schedule.ID, req.Date, studentIDs)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.JSON(w, r, journal)
}

func (h Handler) GetJournal(w http.ResponseWriter, r *http.Request) {
	journalId := chi.URLParam(r, "id")
	journal, err := h.service.Journal.GetJournal(context.Background(), journalId)
	if err != nil {
		panic(err)
	}
	render.JSON(w, r, journal)
}

func (h Handler) DeleteJournal(w http.ResponseWriter, r *http.Request) {
	journalId := chi.URLParam(r, "id")
	err := h.service.Journal.DeleteJournal(context.Background(), journalId)
	if err != nil {
		panic(err)
	}
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
		panic(err)
	}
	render.JSON(w, r, res)
}
