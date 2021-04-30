package service

import (
	"AreaGo/model"
)

func GetOptions(fields []string) []model.Option {
	var options []model.Option
	model.Db.Where("option_key IN (?)", fields).Find(&options)
	return options
}

func GetOption(key string) string {
	var option model.Option
	model.Db.Where("option_key = ?", key).Find(&option)
	return option.OptionVal
}

func SetOption(key, val string) []error {
	//var option model.Option
	//errs := model.Db.Where(model.Option{OptionKey: key}).FirstOrCreate(&option, model.Option{
	//	OptionKey: key,
	//	OptionVal: val,
	//}).GetErrors()
	return nil
}

func BuildUpOptions(options []model.Option) map[string]interface{} {
	var optionMap = make(map[string]interface{})
	for _, t := range options {
		optionMap[t.OptionKey] = t.OptionVal
	}
	return optionMap
}
