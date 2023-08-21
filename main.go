package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type LogicOpenScenario struct {
	XMLName              xml.Name             `xml:"LogicalOpenSCENARIO"`
	FileHeader           FileHeader           `xml:"FileHeader"`
	Properties           Properties           `xml:"Properties"`
	ParameterDeclaration ParameterDeclaration `xml:"ParameterDeclaration"`
}

type FileHeader struct {
	XMLName     xml.Name `xml:"FileHeader"`
	Description string   `xml:"description,attr"`
	Author      string   `xml:"author,attr"`
	RevMajor    string   `xml:"revMajor,attr"`
	Date        string   `xml:"date,attr"`
	RevMinor    string   `xml:"revMinor,attr"`
}

type Properties struct {
	XMLName xml.Name `xml:"Properties"`
	Export  []Export
	Random  []Random
}
type Export struct {
	XMLName   xml.Name    `xml:"Export"`
	Directory []Directory `xml:"Directory"`
}

type Directory struct {
	XMLName xml.Name `xml:"Directory"`
	Path    string   `xml:"path,attr"`
}

type Random struct {
	XMLName xml.Name `xml:"Random"`
	Seed    []Seed
}
type Seed struct {
	XMLName xml.Name `xml:"Seed"`
	Value   int      `xml:"value,attr"`
}

type ParameterDeclaration struct {
	XMLName   xml.Name    `xml:"ParameterDeclaration"`
	Parameter []Parameter `xml:"Parameter"`
}

type Parameter struct {
	XMLName   xml.Name  `xml:"Parameter"`
	Name      string    `xml:"name,attr"`
	Datatype  string    `xml:"datatype,attr"`
	Query     string    `xml:"query,attr"`
	Generator Generator `xml:"Generator"`
}

type Generator struct {
	XMLName xml.Name `xml:"Generator"`
	Type    string   `xml:"type,attr"`
	Samples int      `xml:"samples,attr"`
	Max     float64  `xml:"max,attr"`
	Min     float64  `xml:"min,attr"`
	Step    float64  `xml:"step,attr"`
}

func ParseLosc(xmlFile *os.File) {
	var result LogicOpenScenario
	err := xml.NewDecoder(xmlFile).Decode(&result)

	if err != nil {
		return
	}
	logical := result.ParameterDeclaration
	for i := 0; i < len(logical.Parameter); i++ {
		fmt.Printf("param name=%s , type=%v , q=%v \n", logical.Parameter[i].Name, logical.Parameter[i].Datatype, logical.Parameter[i].Query)
		fmt.Printf("max=%v , min=%v, type=%v, sample=%v, step=%v \n", logical.Parameter[i].Generator.Max, logical.Parameter[i].Generator.Min, logical.Parameter[i].Generator.Type, logical.Parameter[i].Generator.Samples, logical.Parameter[i].Generator.Step)
	}
}

func main() {
	//fileName := "second_vairator.losc"
	fileName := "LS39.losc"
	xmlFile, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Couldn't read"))
	}
	ParseLosc(xmlFile)

}
