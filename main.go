package main

import (
	"app/source"
	"app/target"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type configuration struct {
	SourceDSN string `mapstructure:"source_dsn"`
	TargetDSN string `mapstructure:"target_dsn"`
}

func main() {
	config := loadConfig()

	sourceDB, err := gorm.Open(sqlserver.Open(config.SourceDSN), &gorm.Config{})
	if err != nil {
		panic("no pudo conectar con la base de datos de origen")
	}

	targetDB, err := gorm.Open(postgres.Open(config.TargetDSN), &gorm.Config{})
	if err != nil {
		panic("no pudo conectar con la base de datos de destino")
	}

	sourceRepository := source.CreateRepository(sourceDB)
	targetService := target.CreateService(target.CreateRepository(targetDB))

	clientList, err := sourceRepository.ListAll()
	if err != nil {
		panic("No pudo extraer la lista de clientes de la base de datos origen")
	}

	for _, c := range clientList {
		fmt.Printf("Source CardID %s", c.CardID)
		_, err := targetService.SaveClientIfNotExists(c.FirstName, c.LastName, c.CardType, c.CardID)
		if err != nil {
			panic(fmt.Sprintf("No se pudo traspasar el cliente %s", c.CardID))
		}
	}
}

func loadConfig() *configuration {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	conf := &configuration{}
	err = viper.Unmarshal(conf)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	return conf
}
