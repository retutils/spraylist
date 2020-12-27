package generate

import (
	"io"
	"log"
	"os"
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
	dict.Months = dict.Conf.MonthsYear()
	return dict, nil
}

//MakeDictonary according to config
func (dict *Dictonary) MakeDictonary() {
	for _, month := range dict.Months {
		for _, v := range dict.Conf.Rules.Dictionaries {
			switch v {
			case "Month":
				for _, monthEng := range dict.Conf.Month[month.Month] {
					dict.addMonth(monthEng, month.Year, month.YearShort)
				}
			case "MonthRus":
				for _, monthRus := range dict.Conf.MonthRus[month.Month] {
					rusEng := transform.RussToEngKey(monthRus)
					dict.addMonth(rusEng, month.Year, month.YearShort)
				}
			case "Common":
				for _, word := range dict.Conf.Common.Words {
					dict.addMonth(word, month.Year, month.YearShort)
				}
			case "Company":
				for _, word := range dict.Conf.Company.Names {
					dict.addMonth(word, month.Year, month.YearShort)
				}
			}
		}
	}
	return
}

func (dict *Dictonary) addMonth(month, year, yearShort string) {
	dict.addWithTransforms(month)
	dict.addWithTransforms(transform.AppendSuffix(month, year))
	dict.addWithTransforms(transform.AppendSuffix(month, yearShort))
	if len(dict.Conf.Rules.Delimeters) > 0 {
		for _, delim := range dict.Conf.Rules.Delimeters {
			withdelim := transform.AppendSuffix(month, delim)
			dict.addWithTransforms(transform.AppendSuffix(withdelim, year))
			dict.addWithTransforms(transform.AppendSuffix(withdelim, yearShort))
		}
	}
}

func (dict *Dictonary) addWithTransforms(word string) {
	buff := make([]string, 0)
	for _, v := range dict.Conf.Rules.Transforms {
		switch strings.ToLower(v) {
		case "lower":
			lower := transform.Lower(word)
			dict.Result[lower] = struct{}{}
			buff = append(buff, lower)
		case "upper":
			upper := transform.Uper(word)
			dict.Result[upper] = struct{}{}
			buff = append(buff, upper)
		case "title":
			title := transform.ToTitle(strings.ToLower(word))
			dict.Result[title] = struct{}{}
			buff = append(buff, title)
		case "togle":
			togle := transform.Toggle(word)
			dict.Result[togle] = struct{}{}
			buff = append(buff, togle)
		case "suffixes":
			sfxbuff := make([]string, 0)
			for _, w := range buff {
				dict.addSuffixes(w)
				sfxbuff = append(sfxbuff, w)
			}
			buff = append(buff, sfxbuff...)
		case "replace":
			for _, w := range buff {
				dict.addReplaces(w)
			}
		default:
			log.Fatalf("inforect transform %s", v)
		}
	}
	return
}

func (dict *Dictonary) addSuffixes(word string) {
	for _, sfx := range dict.Conf.Rules.Suffixes {
		dict.Result[transform.AppendSuffix(word, sfx)] = struct{}{}
	}
}
func (dict *Dictonary) addReplaces(word string) {
	for replace, replaceTo := range dict.Conf.Replace {
		dict.Result[transform.Replace(word, replace, replaceTo)] = struct{}{}
	}
}

func (dict *Dictonary) Write() {
	for word := range dict.Result {
		if dict.Policy.Match(word) {
			io.WriteString(os.Stdout, word)
			io.WriteString(os.Stdout, "\n")
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
