package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// / MockProducer реализует интерфейс Producer для тестирования.
type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) Produce() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

// MockPresenter реализует интерфейс Presenter для тестирования.
type MockPresenter struct {
	mock.Mock
}

func (m *MockPresenter) Present(data []string) error {
	args := m.Called(data)
	return args.Error(0)
}

// проверка работы Run при успешном выполнении.
func TestService_Run_Success(t *testing.T) {
	producer := new(MockProducer)
	presenter := new(MockPresenter)

	inputData := []string{
		"Visit http://example.com for details",
		"No links here",
		"Multiple http://test.com in one http://second.com line",
	}

	expectedData := []string{
		"Visit http://*********** for details",
		"No links here",
		"Multiple http://******** in one http://********** line",
	}

	producer.On("Produce").Return(inputData, nil)
	presenter.On("Present", expectedData).Return(nil)

	svc := NewService(producer, presenter)
	err := svc.Run()

	assert.NoError(t, err)
	producer.AssertExpectations(t)
	presenter.AssertExpectations(t)
}

// проверка работы Run, когда продюсер возвращает пустые данные.
func TestService_Run_EmptyData(t *testing.T) {
	producer := new(MockProducer)
	presenter := new(MockPresenter)

	producer.On("Produce").Return([]string{}, nil)

	svc := NewService(producer, presenter)
	err := svc.Run()

	assert.EqualError(t, err, "no data to process")
	producer.AssertExpectations(t)
	presenter.AssertNotCalled(t, "Present", mock.Anything)
}

// проверка обработки ошибки от продюсера.
func TestService_Run_ProducerError(t *testing.T) {
	producer := new(MockProducer)
	presenter := new(MockPresenter)

	producer.On("Produce").Return(nil, errors.New("producer error"))

	svc := NewService(producer, presenter)
	err := svc.Run()

	assert.EqualError(t, err, "producer error")
	producer.AssertExpectations(t)
	presenter.AssertNotCalled(t, "Present", mock.Anything)
}

// проверка обработки ошибки от презентера.
func TestService_Run_PresenterError(t *testing.T) {
	producer := new(MockProducer)
	presenter := new(MockPresenter)

	inputData := []string{
		"Visit http://example.com",
	}

	expectedData := []string{
		"Visit http://***********",
	}

	producer.On("Produce").Return(inputData, nil)
	presenter.On("Present", expectedData).Return(errors.New("presenter error"))

	svc := NewService(producer, presenter)
	err := svc.Run()

	assert.EqualError(t, err, "presenter error")
	producer.AssertExpectations(t)
	presenter.AssertExpectations(t)
}
