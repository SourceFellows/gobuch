package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"golang.source-fellows.com/samples/applicationx/mocks"
)

func TestHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := mocks.NewMockUserService(ctrl)

	handlerFunc := Handler(userService)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/egal", nil)

	//wird der Service überhaupt aufgerufen?
	userService.EXPECT().CreateUser(gomock.Any()).MinTimes(1)

	handlerFunc(w, r)

}
