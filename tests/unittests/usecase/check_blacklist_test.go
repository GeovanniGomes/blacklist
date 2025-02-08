package usecase

import (
	"testing"

	new_black_list_usecase "github.com/GeovanniGomes/blacklist/internal/application/usecase"
	check_mock "github.com/GeovanniGomes/blacklist/tests/unittests/mocks"
	"github.com/golang/mock/gomock"
)

func TestCheckBlacklist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCheckBlacklist := check_mock.NewMockIBlackListRepository(ctrl)
	mockCheckBlacklist.EXPECT().Check(gomock.Any(), gomock.Any()).Return( "Fraude detectada", nil).AnyTimes()
	usecase := new_black_list_usecase.NewCheckBlacklist(mockCheckBlacklist)

	mesage, err := usecase.Execute(10, "event_id")
	if err != nil {
		t.Errorf("Expected nil, got %v", err.Error())
	}

	if mesage != "Fraude detectada" {
		t.Errorf("Expected 'Fraude detectada', got %v", mesage)
	}
}
