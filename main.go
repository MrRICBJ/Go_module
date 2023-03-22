package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Read struct {
}

type Recipes struct {
	XMLName xml.Name `xml:"recipes"`
	Cakes   []Cake   `xml:"cake"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Ingredient struct {
	Name  string  `json:"ingredient_name" xml:"itemname"`
	Count float64 `json:"ingredient_count" xml:"itemcount"`
	Unit  string  `json:"ingredient_unit,omitempty" xml:"itemunit"`
}

func (r *Read) ReadJson(fileName string) ([]Cake, error) {
	data, err := allByte(fileName)
	if err != nil {
		return nil, err
	}
	var cake []Cake
	err = json.Unmarshal(data, &cake)
	if err != nil {
		return nil, err
	}
	return cake, nil
}

func (r *Read) ReadXml(fileName string) (*Recipes, error) {
	data, err := allByte(fileName)
	if err != nil {
		return nil, err
	}
	var cake *Recipes
	err = xml.Unmarshal(data, cake)
	if err != nil {
		return nil, err
	}
	return cake, nil
}

func allByte(fileName string) ([]byte, error) {
	data, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "Couldn't open the file", err)
	}
	defer data.Close()
	byteVal, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "Couldn't read the file", err)
	}
	return byteVal, err
}

func main() {
	filename := flag.String("f", "", "имя файла для чтения")
	flag.Parse()
}
