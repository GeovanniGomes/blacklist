package usecase

import (
	"testing"

	new_black_list_usecase "github.com/GeovanniGomes/blacklist/internal/application/usecase"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	add_mock "github.com/GeovanniGomes/blacklist/tests/unittests/mocks"

	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestBlackList_AddUserBlackList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCheckBlacklist := add_mock.NewMockIBlackListRepository(ctrl)
	mockCheckBlacklist.EXPECT().Add(gomock.Any()).Return(nil).AnyTimes()

	usecase := new_black_list_usecase.NewAddBlacklist(mockCheckBlacklist)
	eventId := "1222"
	blocked, err := usecase.Execute(10, &eventId, "fradue", "10101010112111", nil)

	require.NotNil(t, err)
	require.NotNil(t, blocked)
	require.Equal(t, "event id is not a valid uuid", err.Error())
	eventId =uuid.NewV4().String()
	blocked, err = usecase.Execute(10, &eventId, "fradue", "10101010112111", nil)

	require.Nil(t, err)
	require.Equal(t, entity.PERMANENT, blocked.GetBlockedType())
	require.Equal(t, entity.SPECIFIC, blocked.GetScope())

	blocked, err = usecase.Execute(10, nil, "fradue", "10101010112111", nil)
	require.Nil(t, err)
	require.Equal(t, entity.PERMANENT, blocked.GetBlockedType())
	require.Equal(t, entity.GLOBAL, blocked.GetScope())
}
