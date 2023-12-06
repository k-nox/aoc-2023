package util

import (
	"fmt"
	"os"
	"text/template"
)

type Data struct {
	Day  string
	Part string
}

func GenerateMain(data Data) {
	tmplFile := "main.tmpl"

	tmpl, err := template.New(tmplFile).ParseFiles(fmt.Sprintf("util/%s", tmplFile))
	if err != nil {
		panic(err)
	}

	var f *os.File
	f, err = os.Create(fmt.Sprintf("day%s/main.go", data.Day))
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}

func GenerateSolution(data Data) {
	tmplFile := "solution.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(fmt.Sprintf("util/%s", tmplFile))
	if err != nil {
		panic(err)
	}

	var f *os.File
	f, err = os.Create(fmt.Sprintf("day%s/part%s/part%s.go", data.Day, data.Part, data.Part))
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}

func Generate(day int) {
	dayStr := fmt.Sprintf("%02d", day)
	CreateDayDir(dayStr)
	CreateSolutionDirs(dayStr)
	CreateInputDir(dayStr)
	CreateInputFiles(dayStr)
	GenerateMain(Data{
		Day: dayStr,
	})
	GenerateSolution(Data{
		Day:  dayStr,
		Part: "one",
	})
	GenerateSolution(Data{
		Day:  dayStr,
		Part: "two",
	})
}

func CreateDayDir(dayStr string) {
	err := os.Mkdir(fmt.Sprintf("day%s", dayStr), 0750)
	if err != nil {
		panic(err)
	}
}

func CreateSolutionDirs(dayStr string) {
	err := os.Mkdir(fmt.Sprintf("day%s/partone", dayStr), 0750)
	if err != nil {
		panic(err)
	}
	err = os.Mkdir(fmt.Sprintf("day%s/parttwo", dayStr), 0750)
	if err != nil {
		panic(err)
	}
}

func CreateInputDir(dayStr string) {
	err := os.Mkdir(fmt.Sprintf("day%s/input", dayStr), 0750)
	if err != nil {
		panic(err)
	}
}

func CreateInputFiles(dayStr string) {
	inp, err := os.Create(fmt.Sprintf("day%s/input/input.txt", dayStr))
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(inp)
	sample, err := os.Create(fmt.Sprintf("day%s/input/sample.txt", dayStr))
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(sample)
}
