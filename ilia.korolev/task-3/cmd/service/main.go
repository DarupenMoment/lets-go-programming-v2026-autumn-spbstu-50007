package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"gopkg.in/yaml.v3"
)

const (
	folderMode   = 0o755
	resultMode   = 0o644
	floatBitSize = 64
)

var (
	errEmptyConfigFlag    = errors.New("config flag is required")
	errInvalidConfig      = errors.New("input-file and output-file are required")
	errUnsupportedCharset = errors.New("unsupported charset")
	errEmptyValutes       = errors.New("no valutes found in xml")
	errEmptyCharCode      = errors.New("empty char code in xml")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ResultValute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfgData, err := os.ReadFile(filepath.Clean(*configPath))
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(cfgData, &cfg); err != nil {
		panic(err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic(errInvalidConfig)
	}

	xmlFile, err := os.Open(filepath.Clean(cfg.InputFile))
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = xmlFile.Close()
	}()

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch strings.ToLower(strings.TrimSpace(charset)) {
		case "windows-1251", "cp1251":
			return transform.NewReader(input, charmap.Windows1251.NewDecoder()), nil
		default:
			return nil, fmt.Errorf("%w: %s", errUnsupportedCharset, charset)
		}
	}

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		panic(err)
	}

	if len(valCurs.Valutes) == 0 {
		panic(errEmptyValutes)
	}

	results := make([]ResultValute, 0, len(valCurs.Valutes))
	for _, v := range valCurs.Valutes {
		numCode, err := strconv.Atoi(strings.TrimSpace(v.NumCode))
		if err != nil {
			panic(err)
		}

		charCode := strings.TrimSpace(v.CharCode)
		if charCode == "" {
			panic(errEmptyCharCode)
		}

		valStr := strings.ReplaceAll(strings.TrimSpace(v.Value), ",", ".")
		valFloat, err := strconv.ParseFloat(valStr, floatBitSize)
		if err != nil {
			panic(err)
		}

		results = append(results, ResultValute{
			NumCode:  numCode,
			CharCode: charCode,
			Value:    valFloat,
		})
	}

	sort.SliceStable(results, func(i, j int) bool {
		return results[i].Value > results[j].Value
	})

	outDir := filepath.Dir(cfg.OutputFile)
	if err := os.MkdirAll(outDir, folderMode); err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(cfg.OutputFile, jsonData, resultMode); err != nil {
		panic(err)
	}
}
