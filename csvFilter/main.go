package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func handleErr(m string, err error) {
	if err != nil {
		log.Fatalf("%s\n%v", m, err)
	}
}

func verifyColumns (headers []string, names []string) ([]int, error){
	var indexes []int
	for _, name := range names {
		name = strings.TrimSpace(name)
		i := slices.Index(headers, name)
		if  i < 0  {
			return indexes, errors.New("Column name does not exist")
		}
		indexes = append(indexes, i)
	}

	return indexes, nil
}

func main() {
	var fname string
	var columns string
	var indexes string
	var regex string
	var lines int
	var file io.Reader
	var err error

	flag.StringVar(&columns, "c", "", "Comma separated header names")
	flag.StringVar(&indexes, "i", "", "Comma separated header indexes")
	flag.StringVar(&regex, "r", "", "Regex pattern to search on the file")
	flag.IntVar(&lines, "l", 0, "Limit the number of lines in the output")

	flag.Parse()
	fname = flag.Arg(0)
	if fname == "" {
		file = os.Stdin
	} else {
		file, err = os.Open(fname)
		handleErr("", err)
	}
	switch {
		case regex != "":
			csv, err := io.ReadAll(file)
			handleErr("Unable to read file", err)
			rp := regexp.MustCompile(regex)
			if match := rp.FindAll(csv, -1); match != nil {
				for i,l := range match {
					fmt.Println(string(l))
					if lines > 0 && i >= lines - 1 {
						break
					}
				}
			}
		case columns != "":
			cols := strings.Split(columns, `,`)
			csvR := csv.NewReader(file)
			headers, err := csvR.Read()
			handleErr("Error reading line from CSV file", err)
			idx, err := verifyColumns(headers, cols)
			handleErr("Parse error", err)
			for _, i := range idx {
				fmt.Print(headers[i] + "   ")
			}

			data,err := csvR.ReadAll()
			handleErr("Reading csv data", err)
			for c,line := range data{
				fmt.Println()
				for _,i := range idx{
					fmt.Print(line[i] + "   ")
				}
				if lines > 0 && c >= lines - 1 {
					break
				}
			}
		case indexes != "":
			cols := strings.Split(indexes, `,`)
			csvR := csv.NewReader(file)
			headers, err := csvR.Read()
			handleErr("Error reading line from CSV file", err)
			for _, i := range cols {
				j, err := strconv.Atoi(strings.TrimSpace(i))
				handleErr("Invalid idx", err)
				fmt.Print(headers[j] + "   ")
			}
			data,err := csvR.ReadAll()
			handleErr("Reading csv data", err)
			for c,line := range data{
				fmt.Println()
				for _,i := range cols{
					j, err := strconv.Atoi(strings.TrimSpace(i))
					handleErr("Invalid idx", err)
					fmt.Print(line[j] + "   ")
				}
				if lines > 0 && c >= lines - 1 {
					break
				}
			}


	}


}
