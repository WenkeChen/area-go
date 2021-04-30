package utils

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strconv"
)

func GetOffset(pageString string) (int, int, error) {
	pageInt, err := strconv.Atoi(pageString)
	if err != nil {
		return 0, 0, errors.Wrap(err, "page格式有误")
	}
	sizeInt := viper.GetInt("app.pageSize")
	return sizeInt, (pageInt - 1) * sizeInt, nil
}
