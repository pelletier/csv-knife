package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

func parseCommaSeparatedInts(s string) []int {
	var ints []int
	for _, index := range strings.Split(s, ",") {
		i, err := strconv.Atoi(strings.Trim(index, " "))
		if err != nil {
			log.Fatal("Error during comma-separated column list argument:", err)
		}
		ints = append(ints, i)
	}
	return ints
}

func bitFieldForColumnsList(cols []int, invert bool) uint64 {
	var field uint64

	for _, idx := range cols {
		field |= (uint64(1) << uint(idx))
	}

	if invert {
		field = ^field
	}
	return field
}

// columns management
var columnsWhitelistArg string
var columnsBlacklistArg string

// reader settings
var readerDelimiter string
var readerComment string
var readerFieldsPerRecord int
var readerLazyQuotes bool
var readerTrimLeadingSpace bool

// writer settings
var writerDelimiter string
var writerUseCRLF bool

func init() {
	// columns management
	flag.StringVar(&columnsWhitelistArg, "k", "", "remove all but the given comma separated columns")
	flag.StringVar(&columnsBlacklistArg, "d", "", "keep all but the given comma separated columns")

	// reader settings
	flag.StringVar(&readerDelimiter, "rd", ",", "field delimiter for the input stream")
	flag.StringVar(&readerComment, "rc", "\"", "comment delimiter for the input stream")
	flag.IntVar(&readerFieldsPerRecord, "rf", 0, "number of fields per record (0 = all equal to the first row, -1 = no check)")
	flag.BoolVar(&readerLazyQuotes, "rl", false, "allow quotes not to be closed")
	flag.BoolVar(&readerTrimLeadingSpace, "rt", true, "trim leading spaces in fields")

	// writer settings
	flag.StringVar(&writerDelimiter, "wd", ",", "field delimiter for the output stream")
	flag.BoolVar(&writerUseCRLF, "wc", false, "use CRLF as a new line character in the output stream")
}

func main() {
	flag.Parse()
	if len(columnsWhitelistArg) > 0 && len(columnsBlacklistArg) > 0 {
		log.Fatal("You cannot provide both -d and -k.")
	}

	var columnsField uint64
	if len(columnsWhitelistArg) > 0 {
		columnsField = bitFieldForColumnsList(parseCommaSeparatedInts(columnsWhitelistArg), false)
	} else if len(columnsBlacklistArg) > 0 {
		columnsField = bitFieldForColumnsList(parseCommaSeparatedInts(columnsBlacklistArg), true)
	}

	stdinReader := bufio.NewReader(os.Stdin)
	csvReader := csv.NewReader(stdinReader)
	r, _ := utf8.DecodeRuneInString(readerDelimiter)
	csvReader.Comma = r
	r, _ = utf8.DecodeRuneInString(readerComment)
	csvReader.Comment = r
	csvReader.FieldsPerRecord = readerFieldsPerRecord
	csvReader.LazyQuotes = readerLazyQuotes
	csvReader.TrimLeadingSpace = readerTrimLeadingSpace

	stdoutWriter := bufio.NewWriter(os.Stdout)
	csvWriter := csv.NewWriter(stdoutWriter)
	r, _ = utf8.DecodeRuneInString(writerDelimiter)
	csvWriter.Comma = r
	csvWriter.UseCRLF = writerUseCRLF

	buff := make(chan []string, 2)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for row := range buff {
			b := row[:0]
			for index, field := range row {
				mask := uint64(1 << uint(index))
				if columnsField&mask == mask {
					b = append(b, field)
				}
			}
			err := csvWriter.Write(b)
			if err != nil {
				log.Fatal("Unexpected write error:", err)
			}
		}
		csvWriter.Flush()
		err := csvWriter.Error()
		if err != nil {
			log.Fatal("Unexpected flush error:", err)
		}
		wg.Done()
	}()

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Unexpected read error:", err)
		}
		buff <- row
	}

	close(buff)
	wg.Wait()
}
