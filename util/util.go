package util

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

// SettingStruct is class for default settings
type SettingStruct struct {
	ConectionString string `yaml:"ConectionString"`
	SuperAdmin      string `yaml:"SuperAdmin"`
	Password        string `yaml:"Password"`
	Protocol        string `yaml:"Protocol"`
	Domain          string `yaml:"Domain"`
	AppName         string `yaml:"AppName"`
	Minify          bool
}

/*
	S3Bucket             string `yaml:"Bucket"`
	S3Region             string `yaml:"Region"`
	S3AwsAccessKeyID     string `yaml:"AwsAccessKeyID"`
	S3AwsSecretAccessKey string `yaml:"AwsSecretAccessKey"`
	SendgridAPIKey       string `yaml:"SendgridAPIKey"`
	SendgridSenderEmail  string `yaml:"SendgridSenderEmail"`
	SendgridSenderName   string `yaml:"SendgridSenderName"`
*/

// Settings of application
var Settings SettingStruct

var logger *log.Logger

func init() {

	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Log(err)
	}
	logger = log.New(f, "prefix", log.LstdFlags)

	Log("utility module initiating")
	defer Log("utility module initiated")

	f, err = os.Open("../cubizy.txt")
	if err != nil {
		Log(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	Settings = SettingStruct{}

	scanner.Scan()
	Settings.ConectionString = strings.TrimSpace(scanner.Text())
	scanner.Scan()
	Settings.SuperAdmin = strings.TrimSpace(scanner.Text())
	scanner.Scan()
	Settings.Password = strings.TrimSpace(scanner.Text())
	scanner.Scan()
	Settings.Protocol = strings.TrimSpace(scanner.Text())
	scanner.Scan()
	Settings.Domain = strings.TrimSpace(scanner.Text())
	scanner.Scan()
	Settings.AppName = strings.TrimSpace(scanner.Text())
	scanner.Scan()
	minify := strings.TrimSpace(scanner.Text())
	if minify == "minify" {
		Settings.Minify = true
	} else {
		Settings.Minify = false
	}
}

func GenVar(length int) string {
	str := ""
	for i := 0; i < length; i++ {
		if str != "" {
			str += ", "
		}
		str += "?"
	}
	return str
}

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Log(arg ...interface{}) {
	logger.Println(arg...)
	log.Println(arg...)
}

func Panic(arg ...interface{}) {
	logger.Println(arg...)
	log.Panicln(arg...)
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// DoEvery is function to run any function periodically after given Duration
func DoEvery(d time.Duration, f func(time.Time)) {
	if d == 0 {
		return
	}
	for x := range time.Tick(d) {
		f(x)
	}
}
