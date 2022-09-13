package handler

import (
	"api-gateway/pkg/httperr"
	"api-gateway/request"
	"api-gateway/response"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h Handler) RegisterSchedule(w http.ResponseWriter, r *http.Request) {
	var req request.CreateScheduleRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}

	schedule, err := h.service.Schedule.RegisterSchedule(context.Background(), req)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, schedule)
}

func (h Handler) UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	var req request.Schedule
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}

	res, err := h.service.Schedule.UpdateSchedule(context.Background(), req)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

func (h Handler) DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	scheduleID := chi.URLParam(r, "id")

	if err := h.service.Schedule.DeleteSchedule(context.Background(), scheduleID); err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, render.M{
		"ok": true,
	})
}

func (h Handler) GetFullScheduleForTeacher(w http.ResponseWriter, r *http.Request) {
	teacherID := chi.URLParam(r, "id")

	schedules, err := h.service.Schedule.GetFullScheduleForTeacher(context.Background(), teacherID)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	populatedSchedules, err := h.populateSchedules(schedules)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, populatedSchedules)
}

func (h Handler) GetFullScheduleForGroup(w http.ResponseWriter, r *http.Request) {
	groupID := chi.URLParam(r, "id")

	schedules, err := h.service.Schedule.GetFullScheduleForGroup(context.Background(), groupID)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	populatedSchedules, err := h.populateSchedules(schedules)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, populatedSchedules)
}

func (h Handler) GetSpecificDateScheduleForTeacher(w http.ResponseWriter, r *http.Request) {
	var req request.GetSpecificDateScheduleForTeacherRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}

	schedules, err := h.service.Schedule.GetSpecificDateScheduleForTeacher(context.Background(), req.TeacherID, req.Date)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	populatedSchedules, err := h.populateSchedules(schedules)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, populatedSchedules)
}

func (h Handler) GetSpecificDateScheduleForGroup(w http.ResponseWriter, r *http.Request) {
	var req request.GetSpecificDateScheduleForGroupRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}

	schedules, err := h.service.Schedule.GetSpecificDateScheduleForGroup(context.Background(), req.GroupID, req.Date)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	populatedSchedules, err := h.populateSchedules(schedules)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, populatedSchedules)
}

func (h Handler) GetScheduleById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	schedule, err := h.service.Schedule.GetSchedule(context.Background(), id)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, schedule)
}

func (h Handler) populateSchedules(schedules []response.Schedule) ([]response.PopulatedSchedule, error) {
	populatedSchedules := make([]response.PopulatedSchedule, 0, len(schedules))
	for _, sch := range schedules {
		group, err := h.service.Student.GetGroup(context.Background(), sch.GroupId)
		if err != nil {
			return nil, err
		}

		subject, err := h.service.Teacher.GetSubject(context.Background(), sch.SubjectID)
		if err != nil {
			return nil, err
		}

		teacher, err := h.service.Teacher.GetTeacher(context.Background(), request.GetTeacherRequest{
			TeacherID:   sch.TeacherID,
			Email:       "",
			PhoneNumber: "",
		})
		if err != nil {
			return nil, err
		}

		populatedSchedules = append(populatedSchedules, response.PopulatedSchedule{
			ID:           sch.ID,
			Group:        group,
			Subject:      subject,
			Teacher:      teacher,
			Weekday:      sch.WeekDay,
			LessonNumber: sch.LessonNumber,
		})
	}

	return populatedSchedules, nil
}
