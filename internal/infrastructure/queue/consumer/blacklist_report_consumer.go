package consumer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/xuri/excelize/v2"
)

type BlacklistReportConsumer struct {
	queue               contracts.IQueue
	blacklistRepository repositoty.IBlackListRepository
	clientUpload        contracts.IFileSystem
}

func NewBlacklistReportConsumer(
	queue contracts.IQueue,
	blacklistRepository repositoty.IBlackListRepository,
	clientUpload contracts.IFileSystem,

) *BlacklistReportConsumer {
	return &BlacklistReportConsumer{
		queue:               queue,
		clientUpload:        clientUpload,
		blacklistRepository: blacklistRepository,
	}
}

func (c *BlacklistReportConsumer) HandleMessage() func([]byte) error {
	return func(message []byte) error {
		log.Printf("Process message blacklist: %s", message)
		currentDate := time.Now()
		year, month, day := currentDate.Date()
		dateGenetarion := fmt.Sprintf("%v_%v_%v_%v", year, month, day, currentDate.Second())

		var msg map[string]interface{}
		if err := json.Unmarshal(message, &msg); err != nil {
			return err
		}

		data, ok := msg["data"].(map[string]interface{})
		if !ok {
			log.Println("Error accessing the field 'data'")
			return nil
		}

		startDate, err := c.parseDate(data["start_date"].(string))
		if err != nil {
			return err
		}

		endDate, err := c.parseDate(data["end_date"].(string))
		if err != nil {
			return err
		}

		blacklistToReports, err := c.blacklistRepository.FetchBlacklistEntries(startDate, endDate)

		if len(blacklistToReports) == 0 {
			log.Print("Empty list to generate report")
		}

		if err != nil {
			return err
		}

		f := excelize.NewFile()
		sheetName := "Blacklist"
		index, _ := f.NewSheet(sheetName)

		headers := []string{"Data Criacao", "ID Evento", "ID Usuario", "Scopo", "Motivo"}
		columnMap := make(map[string]string)

		for i, h := range headers {
			col := string(rune('A' + i))
			cell := fmt.Sprintf("%s1", col)
			f.SetCellValue(sheetName, cell, h)
			columnMap[h] = col
		}

		for i, evento := range blacklistToReports {
			row := i + 2

			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", columnMap["Data Criacao"], row), evento.GetCreatedAt().Format("2006-01-02 15:04:05"))
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", columnMap["ID Evento"], row), evento.GetEventId())
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", columnMap["ID Usuario"], row), evento.GetUserIdentifier())
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", columnMap["Scopo"], row), evento.GetScope())
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", columnMap["Motivo"], row), evento.GetReason())
		}

		f.SetActiveSheet(index)

		var buf bytes.Buffer
		if err := f.Write(&buf); err != nil {
			log.Printf("Erro ao escrever o arquivo no buffer: %v", err)
			return err
		}

		fileName := fmt.Sprintf("%v.xlsx", dateGenetarion)
		c.clientUpload.Upload("file-blacklist", fileName, buf)
		return nil
	}
}

func (c *BlacklistReportConsumer) parseDate(value string) (time.Time, error) {
	startDate, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}, err
	}
	return startDate, nil
}
