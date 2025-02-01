package entity_test

import (
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/domain/value_objects"
	"testing"
	"time"
	"github.com/stretchr/testify/require"
)

// Helper function to create a new Category
func createCategory(categoryName string) *value_objects.Category {
	category := value_objects.Category{}
	newCategory, _ := category.NewCategory(categoryName)
	return newCategory
}

func createEvent(title, description string, date time.Time, category *value_objects.Category) *entity.Event {
	return entity.NewEvent(title, description, date, *category)
}

func TestEvent_Enable(t *testing.T) {
	// Setup test fixtures
	category := createCategory(value_objects.SOCCER)
	dateEvent := time.Now()
	event := createEvent("Jogo do Flamengo", "Campeonato carioca", dateEvent, category)
	// Act
	err := event.Enable()

	// Assertions
	require.Nil(t, err)
	require.NotNil(t, event.GetId())
	require.Equal(t, event.GetStatus(), entity.ENABLED)
	require.Equal(t, event.GetDate(), dateEvent)
	require.Equal(t, event.GetCategory().GetName(), value_objects.SOCCER)
	require.NotNil(t, event.GetCategory().GetCode())
	require.Equal(t, event.GetDescription(), "Campeonato carioca")
	require.Equal(t, event.GetTitle(), "Jogo do Flamengo")
	require.Equal(t, event.GetIsActive(), false)

	event.ChangeDateEvent(time.Now().AddDate(0, 0, 10))
	error_enable := event.Enable()
	require.Nil(t, error_enable)

	event.ChangeDateEvent(time.Now().AddDate(0, 0, -5))
	error_enable = event.Enable()
	require.Equal(t, "it is not possible to enable an event with a past date", error_enable.Error())

}

func TestEvent_Disable(t *testing.T) {
	// Setup test fixtures
	category := createCategory(value_objects.SOCCER)
	dateEvent := time.Now()
	event := createEvent("Jogo do Flamengo", "Campeonato carioca", dateEvent, category)
	// Act
	err := event.Disable()

	// Assertions
	require.Nil(t, err)
	require.NotNil(t, event.GetId())
	require.Equal(t, event.GetStatus(), entity.DISABLED)
	require.Equal(t, event.GetDate(), dateEvent)
	require.Equal(t, event.GetCategory().GetName(), value_objects.SOCCER)
	require.NotNil(t, event.GetCategory().GetCode())
	require.Equal(t, event.GetDescription(), "Campeonato carioca")
	require.Equal(t, event.GetTitle(), "Jogo do Flamengo")

	event.ChangeDateEvent(time.Now().AddDate(0, 0, 10))
	error_disable := event.Disable()
	require.Equal(t, "it is not possible to disable an event with a past date", error_disable.Error())

	event.ChangeDateEvent(time.Now().AddDate(0, 0, -5))
	error_disable = event.Disable()
	require.Nil(t, error_disable)
}
func TestEvent_NewEvent_Date_Invalid(t *testing.T) {
	category := createCategory(value_objects.CARNIVAL)
	dateEvent := time.Now().AddDate(0, 0, -5)

	event := createEvent("Jogo do Flamengo", "Campeonato carioca", dateEvent, category)
	_, err := event.IsValid()
	require.NotNil(t, err)

	require.Equal(t, "Event date cannot be less than the current date", err.Error())
}

func TestEvent_ChangeCategory(t *testing.T) {
	// Setup test fixtures
	category := createCategory(value_objects.SOCCER)
	dateEvent := time.Now()
	event := createEvent("Jogo do Flamengo", "Campeonato carioca", dateEvent, category)

	// Act
	newCategory := createCategory(value_objects.CARNIVAL)
	event.ChangeCatrgory(*newCategory)

	// Assertions
	require.Equal(t, event.GetCategory().GetName(), value_objects.CARNIVAL)
	require.NotEqual(t, event.GetCategory().GetCode(), value_objects.SOCCER)
	require.Equal(t, event.GetDescription(), "Campeonato carioca")
	require.Equal(t, event.GetTitle(), "Jogo do Flamengo")
}
