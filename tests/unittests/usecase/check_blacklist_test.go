package usecase

import (
	"testing"

	new_black_list_usecase "github.com/GeovanniGomes/blacklist/internal/application/usecase"
	"github.com/GeovanniGomes/blacklist/tests/fixtures"
	check_mock "github.com/GeovanniGomes/blacklist/tests/unittests/mocks"
	"github.com/golang/mock/gomock"
)

func TestCheckBlacklist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCheckBlacklist := check_mock.NewMockIBlackListRepository(ctrl)
	entityMock := fixtures.CreateBlacklist(10,nil,"Fraude detectada","123456","global","permanent",nil)
	mockCheckBlacklist.EXPECT().Check(gomock.Any(), gomock.Any()).Return(entityMock, nil).AnyTimes()
	usecase := new_black_list_usecase.NewCheckBlacklist(mockCheckBlacklist)
	eventId := "event_id"
	message, err := usecase.Execute(10, &eventId)
	if err != nil {
		t.Errorf("Expected nil, got %v", err.Error())
	}

	if message != "Fraude detectada" {
		t.Errorf("Expected 'Fraude detectada', got %v",message)
	}
}
