package generate

import (
	"fmt"
	"strings"

	"github.com/spraylist/cfg"
	"github.com/spraylist/policy"
	"github.com/spraylist/transform"
)

//Dictonary data
type Dictonary struct {
	Result map[string]struct{} //We use map to ommit duplicates
	Conf   *cfg.Config
	Policy *policy.Policy
	Months []cfg.Seazon
	Years  map[string]struct{}
}

//New Dictonary from config file
func New(path string) (dict *Dictonary, err error) {
	dict = new(Dictonary)
	dict.Conf, err = cfg.New(path)
	if err != nil {
		return nil, err
	}
	dict.Policy, err = policy.New(dict.Conf.Policy.RequiredChargroups, dict.Conf.Policy.Lehgth, dict.Conf.Policy.CharacterGroup...)
	if err != nil {
		return nil, err
	}
	dict.Result = make(map[string]struct{}, 0)
	dict.Years = make(map[string]struct{}, 0)
	dict.Months = dict.Conf.MonthsYear()
	return dict, nil
}

//Make seazon Dictonary according to config
func (dict *Dictonary) makeDict() {
	for _, month := range dict.Months {
		dict.Years[month.Year] = struct{}{}
		dict.Years[month.YearShort] = struct{}{}
		if dict.hasDict("Month") {
			for _, monthEng := range dict.Conf.Month[month.Month] {
				dict.addMonth(monthEng, month.Year, month.YearShort)
			}
		}
		if dict.hasDict("MonthRus") {
			for _, monthRus := range dict.Conf.MonthRus[month.Month] {
				rusEng := transform.RussToEngKey(monthRus)
				dict.addMonth(rusEng, month.Year, month.YearShort)

			}
		}
		if dict.hasDict("Common") {
			for _, word := range dict.Conf.Common.Words {
				dict.addMonth(word, month.Year, month.YearShort)

			}
		}

	}
	return
}

func (dict *Dictonary) addMonth(month, year, yearShort string) {
	dict.Result[month] = struct{}{}
	dict.Result[transform.AppendSuffix(month, year)] = struct{}{}
	dict.Result[transform.AppendSuffix(month, yearShort)] = struct{}{}
	if len(dict.Conf.Rules.Delimeters) > 0 {
		for _, delim := range dict.Conf.Rules.Delimeters {
			withdelim := transform.AppendSuffix(month, delim)
			dict.Result[transform.AppendSuffix(withdelim, year)] = struct{}{}
			dict.Result[transform.AppendSuffix(withdelim, yearShort)] = struct{}{}
		}
	}
}

func (dict *Dictonary) addTransforms() error {
	for word := range dict.Result {
		for _, v := range dict.Conf.Rules.Transforms {
			switch strings.ToLower(v) {
			case "lower":
				dict.Result[transform.Lower(word)] = struct{}{}
			case "upper":
				dict.Result[transform.Uper(word)] = struct{}{}
			case "title":
				dict.Result[transform.Uper(word)] = struct{}{}
			case "togle":
				dict.Result[transform.Toggle(word)] = struct{}{}
			default:
				return fmt.Errorf("inforect trnsform %s", v)
			}
		}
	}
	return nil
}

func (dict *Dictonary) addSuffixes() {
	for word := range dict.Result {
		for _, sfx := range dict.Conf.Rules.Suffixes {
			dict.Result[transform.AppendSuffix(word, sfx)] = struct{}{}
		}
	}
}
func (dict *Dictonary) addReplaces() {
	for word := range dict.Result {
		for replace, replaceTo := range dict.Conf.Replace {
			dict.Result[transform.Replace(word, replace, replaceTo)] = struct{}{}
		}
	}
}

func (dict *Dictonary) hasDict(name string) bool {
	for _, v := range dict.Conf.Rules.Dictionaries {
		if v == name {
			return true
		}
	}
	return false
}
