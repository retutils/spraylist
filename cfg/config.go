package cfg

import (
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
)

//Config holds configuration
type Config struct {
	Policy   *Policy
	Rules    *Rules
	Replace  map[string]string
	Month    map[string][]string
	MonthRus map[string][]string
	Company  *Company
	Common   *Common
}

//Policy for password generation
type Policy struct {
	RequiredChargroups int      //The password should fulfill at least $RequiredChargroups of the following for character group
	CharacterGroup     []string //Char group preset in the passwords
	Lehgth             int      //Min passwords length
	Expire             int      //Password expiration in days
}

//Rules for password generation
type Rules struct {
	Transforms   []string
	Dictionaries []string
	Suffixes     []string
	Delimeters   []string
}

//Company names
type Company struct {
	Names []string
}

//Common config
type Common struct {
	Words []string
}

//Seazon keep pair of
type Seazon struct {
	Month, Year, YearShort string
}

//MonthsYear return mounths, years dict
func (c *Config) MonthsYear() (months []Seazon) {
	now := time.Now()
	ttl := 0 - c.Policy.Expire
	month := now.AddDate(0, 0, ttl).Month()
	months = make([]Seazon, 0)
	for ttl <= 0 {
		getAt := now.AddDate(0, 0, ttl)
		nextMonth := getAt.Month()
		if nextMonth != month || ttl == 0-c.Policy.Expire {
			year := strconv.Itoa(getAt.Year())
			m := Seazon{Month: getAt.Month().String(), Year: year, YearShort: year[2:]}
			months = append(months, m)
		}
		ttl++
		month = nextMonth
	}
	return months
}

//New configuration form file
func New(path string) (c *Config, err error) {
	c = new(Config)
	if _, err := toml.DecodeFile(path, c); err != nil {
		return c, err
	}
	return c, err
}
