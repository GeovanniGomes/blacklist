package entity_test

import (
	"testing"
	"time"

	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/domain/value_objects"
	"github.com/GeovanniGomes/blacklist/tests/fixtures"
	"github.com/stretchr/testify/require"
)



func TestEvent_NewEvent_Date_Invalid(t *testing.T) {
	factoryEntity := entity.FactoryEntity{}
	dateEvent := time.Now().AddDate(0, 0, -5)

	event, _ := factoryEntity.FactoryNewEvent(nil, "Jogo do Flamengo", "Campeonato carioca", dateEvent, nil, value_objects.CARNIVAL, true)
	err := event.IsValid()
	require.NotNil(t, err)

	require.Equal(t, "Event date cannot be less than the current date", err.Error())
}

func TestEvent_ChangeCategory(t *testing.T) {
	factoryEntity := entity.FactoryEntity{}
	dateEvent := time.Now()
	event, err := factoryEntity.FactoryNewEvent(nil, "Jogo do Flamengo", "Campeonato carioca", dateEvent, nil, value_objects.SOCCER, true)
	require.Nil(t, err)
	newCategory := fixtures.CreateCategory(value_objects.CARNIVAL)
	event.ChangeCatrgory(*newCategory)

	// Assertions
	require.Equal(t, event.GetCategory().GetName(), value_objects.CARNIVAL)
	require.NotEqual(t, event.GetCategory().GetCode(), value_objects.SOCCER)
	require.Equal(t, event.GetDescription(), "Campeonato carioca")
	require.Equal(t, event.GetTitle(), "Jogo do Flamengo")
}
