package service

import (
	"errors"
	"strings"
)

// Service структура, реализующая бизнес-логику
type Service struct {
	prod Producer
	pres Presenter
}

// NewService конструктор Service
func NewService(prod Producer, pres Presenter) *Service {
	return &Service{prod: prod, pres: pres}
}

func (s *Service) MaskString(input string) string {
	return strings.ReplaceAll(input, "secret", "******")
}

// Run основной метод сервиса
func (s *Service) Run() error {
	// Получаем данные от Producer
	data, err := s.prod.Produce()
	if err != nil {
		return err
	}

	// Проверяем, что данные не пустые
	if len(data) == 0 {
		return errors.New("no data to process")
	}

	// Обрабатываем данные
	var processedData []string
	for _, line := range data {
		processedData = append(processedData, s.MaskString(line))
	}

	// Передаем обработанные данные Presenter
	return s.pres.Present(processedData)
}
