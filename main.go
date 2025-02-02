package main

import (
	"flag"
	"fmt"
	"log"
	"project/service"
)

func main() {
	// Получаем путь к файлу для чтения и записи из аргументов
	inputFile := flag.String("input", "", "input file")
	outputFile := flag.String("output", "output.txt", "output file")
	flag.Parse()

	if *inputFile == "" {
		log.Fatal("Input file is required. Use -input to specify the file.")
	}

	producer := service.NewFileProducer(*inputFile)
	presenter := service.NewFilePresenter(*outputFile)

	svc := service.NewService(producer, presenter)

	if err := svc.Run(); err != nil {
		log.Fatalf("Service failed: %v", err)
	}

	fmt.Println("Processing completed successfully.")
}
