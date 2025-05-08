package service

import (
	"errors"
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

// функция ищет подстроку "http://" и заменяет все символы после неё до первого пробела на звездочки.
func maskLinks(input string) string {
	text := []byte(input)
	poisk := []byte("http://")

	for i := 0; i <= len(text)-len(poisk); i++ {
		match := true
		for j := 0; j < len(poisk); j++ {
			if text[i+j] != poisk[j] {
				match = false
				break
			}
		}

		if match {
			i += len(poisk)
			for i < len(text) && text[i] != ' ' {
				text[i] = '*'
				i++
			}
		}
	}

	return string(text)
}

func (s *Service) MaskString(input string) string {
	return maskLinks(input)
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
