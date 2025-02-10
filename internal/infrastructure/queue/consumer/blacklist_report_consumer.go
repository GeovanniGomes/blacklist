package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/xuri/excelize/v2"
)

type BlacklistReportConsumer struct {
	queue                contracts.IQueue
	blacklist_repository repositoty.IBlackListRepository
}

func NewBlacklistReportConsumer(queue contracts.IQueue, blacklist_repository repositoty.IBlackListRepository) *BlacklistReportConsumer {
	return &BlacklistReportConsumer{queue: queue, blacklist_repository: blacklist_repository}
}

func (c *BlacklistReportConsumer) HandleMessage() func([]byte) error {
	return func(message []byte) error {
		log.Printf("Processando mensagem da blacklist: %s", message)

		var msg map[string]string
		if err := json.Unmarshal(message, &msg); err != nil {
			return err
		}

		startDate, err := time.Parse(time.RFC3339, msg["start_date"])
		if err != nil {
			return err
		}
		startDate = startDate.Truncate(24 * time.Hour)

		endDate, err := time.Parse(time.RFC3339, msg["end_date"])
		if err != nil {
			return err
		}

		endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, endDate.Location())

		blacklistToReports, err := c.blacklist_repository.FetchBlacklistEntries(startDate, endDate)
		if err != nil {
			return err
		}

		f := excelize.NewFile()
		sheetName := "Eventos"
		index, _ := f.NewSheet(sheetName)

		headers := []string{"Data Criacao", "ID Evento", "ID Usuario", "Scopo", "Motivo"}
		columnMap := make(map[string]string)

		for i, h := range headers {
			col := string(rune('A' + i))
			cell := fmt.Sprintf("%s1", col)
			f.SetCellValue(sheetName, cell, h)
			columnMap[h] = col
		}

		// Inserindo dados na planilha
		for i, evento := range blacklistToReports {
			row := i + 2 // Começa na linha 2 (1 é o cabeçalho)

			// Preencher células dinamicamente usando o mapeamento
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", columnMap["Data Criacao"], row), evento.GetCreatedAt().Format("2006-01-02 15:04:05"))
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", columnMap["ID Evento"], row), evento.GetEventId())
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", columnMap["ID Usuario"], row), evento.GetUserIdentifier())
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", columnMap["Scopo"], row), evento.GetScope())
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", columnMap["Motivo"], row), evento.GetReason())
		}

		// Definir a planilha ativa
		f.SetActiveSheet(index)

		// Salvar o arquivo Excel
		if err := f.SaveAs("eventos.xlsx"); err != nil {
			log.Printf("Erro ao salvar o arquivo: %v", err)
		}

		fmt.Println("Arquivo Excel criado com sucesso: eventos.xlsx")
		return nil // Caso precise retornar erro, pode tratar aqui
	}
}
