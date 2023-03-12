package conf

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-yaml/yaml"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var Env = flag.String("env", "local", "Please run app with environment -> ./app -env environment")
var global = "global"

// Local settings from config/{environment}.yaml
var Local Environment

// Global settings from config/global.yaml
var Global Environment

type Environment struct {
	Name *string
}

// [environment][setting][value]
var configData map[*string]map[string]interface{}
var re *regexp.Regexp

func init() {
	flag.Parse()
	configData = make(map[*string]map[string]interface{})
	re = regexp.MustCompile("^\\s*([\\w-]*)\\s*:\\s*(.*)\\s*")
	Global.Name = &global
	if len(os.Args) > 1 {
		Local.Name = Env
	} else {
		panic("Please run app with environment -> ./app -env environment")
	}
}

// GetEnv Return current environment, dev is default
func GetEnv() string {
	return *Local.Name
}

// GetString setting as string
func (e Environment) GetString(setting string) (result string) {
	environmentMap := fetchenvironment(e)
	val, _ := environmentMap[setting]

	parse(val, &result)
	return
}

// GetUint get setting as uint
func (e Environment) GetUint(setting string) uint {
	var result uint

	environmentMap := fetchenvironment(e)
	val, _ := environmentMap[setting]
	parse(val, &result)

	return result
}

// GetUint64 get setting as uint64
func (e Environment) GetUint64(setting string) uint64 {
	var result uint64

	environmentMap := fetchenvironment(e)
	val, _ := environmentMap[setting]
	parse(val, &result)

	return result
}

// GetInt get setting as int
func (e Environment) GetInt(setting string) int {
	var result int

	environmentMap := fetchenvironment(e)
	val, _ := environmentMap[setting]
	parse(val, &result)

	return result
}

// GetInt64 get setting as int64
func (e Environment) GetInt64(setting string) int64 {
	var result int64

	environmentMap := fetchenvironment(e)
	val, _ := environmentMap[setting]
	parse(val, &result)

	return result
}

// GetFloat64 get setting as float64
func (e Environment) GetFloat64(setting string) float64 {
	var strVal string

	environmentMap := fetchenvironment(e)
	val, _ := environmentMap[setting]
	parse(val, &strVal)

	parsedVal, _ := strconv.ParseFloat(strVal, 64)
	return parsedVal
}

// GetBool get setting as boolean
func (e Environment) GetBool(setting string) bool {
	var strVal string

	environmentMap := fetchenvironment(e)
	val, _ := environmentMap[setting]
	parse(val, &strVal)

	parsedVal, _ := strconv.ParseBool(strVal)
	return parsedVal
}

// GetSlice get setting as slice of strings
func (e Environment) GetSlice(setting string) (result []string) {
	environmentMap := fetchenvironment(e)
	val, _ := environmentMap[setting]

	result = strings.Split(val.(string), ",")
	return
}

func fetchenvironment(e Environment) map[string]interface{} {
	environmentMap, ok := configData[e.Name]
	// singleton
	if !ok {
		e.importSettingsFromFile()
		environmentMap, _ = configData[e.Name]
	}
	return environmentMap
}

func (e Environment) importSettingsFromFile() {
	config := make(map[string]interface{})
	name := *e.Name
	file, err := os.ReadFile("config/" + name + ".yaml")
	if err != nil {
		panic(fmt.Sprintf("Open config file fail: config/%s.yaml. Please run application as ./app [dev] ", *e.Name))
		return
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(fmt.Sprintf("Parse config file fail: config/%s.yaml %s", *e.Name, err.Error()))
		return
	}
	configData[e.Name] = config
}

func parse(in interface{}, out interface{}) {
	bytes, _ := json.Marshal(in)
	_ = json.Unmarshal(bytes, &out)
}
