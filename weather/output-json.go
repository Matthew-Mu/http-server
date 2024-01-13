package weather

import (
	"encoding/json"
	"fmt"
)

type WeatherTable struct {
	Table []Weather
}

func ConvertBytesToJson(b []byte) WeatherTable {

	var structArray []Weather
	err := json.Unmarshal([]byte(b), &structArray)
	if err != nil {
		fmt.Println(err.Error())
	}
	wTable := WeatherTable{
		Table: structArray,
	}

	return wTable
}
