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

	mockCheckBlacklist := check_mock.NewMockBlackListRepositoryInterface(ctrl)
	mockAddLog := check_mock.NewMockAuditLoggerInterface(ctrl)

	mockCheckBlacklist.EXPECT().Check(gomock.Any(), gomock.Any()).Return(true, "Fraude detectada").AnyTimes()
	mockAddLog.EXPECT().LogAction(gomock.Any(), gomock.Any(), gomock.Any(), gomock.AssignableToTypeOf(&map[string]interface{}{})).AnyTimes()
	usecase := new_black_list_usecase.NewCheckBlacklist(mockCheckBlacklist, mockAddLog)

	blocked, mesage := usecase.Execute(10, "event_id")
	if blocked != true {
		t.Errorf("Expected true, got %v", blocked)
	}
	if mesage != "Fraude detectada" {
		t.Errorf("Expected 'Fraude detectada', got %v", mesage)
	}
}
