package usecase

import (
	new_black_list_usecase "github.com/GeovanniGomes/blacklist/internal/application/usecase"
	check_mock "github.com/GeovanniGomes/blacklist/tests/unittests/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestRemoveBlacklist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCheckBlacklist := check_mock.NewMockBlackListRepositoryInterface(ctrl)
	mockCheckBlacklist.EXPECT().Remove(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	usecase:= new_black_list_usecase.NewRemoveBlacklist(mockCheckBlacklist)
	err := usecase.Execute(10, "event_id")

	if err != nil {
		t.Errorf("Expected err nil, got %v", err)
	}
}
