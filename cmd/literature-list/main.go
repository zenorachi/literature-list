package main

import (
	"fmt"
	"github.com/zenorachi/literature-list/internal/config"
	"github.com/zenorachi/literature-list/internal/service"
	"time"
)

func main() {
	//body := models.CyberleninkaRequest{
	//	Mode: "articles",
	//	Size: 1,
	//	Q:    "Панкратов",
	//}
	//kek, _ := json.Marshal(body)
	//
	//resp, _ := http.Post("https://cyberleninka.ru/api/search", "application/json", bytes.NewBuffer(kek))
	//fmt.Println(resp)
	cfg := &config.Config{
		HTTP: config.HTTPConfig{},
		GIN:  config.GINConfig{},
		CyberleninkaClient: config.ClientConfig{
			BaseURL: "https://cyberleninka.ru",
			Timeout: 10 * time.Second,
		},
		ELibraryClient: config.ClientConfig{},
	}
	s := service.New(cfg)
	res, _ := s.SearchLiterature([]string{"Пакратов И. Ю.", "Пушкин А. С."})
	for _, v := range res {
		fmt.Printf("Title: %s\n, Contained: %v\n", v.Title, v.IsContained)
	}
}
