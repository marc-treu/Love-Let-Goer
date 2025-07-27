package configs

var DB map[string][]string

func ConfigureDB() {
	DB = make(map[string][]string)
}
