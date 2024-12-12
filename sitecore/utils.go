package sitecore

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"guitarrich/xmcloud/model"
)

func GetEnvVar(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error getting env var: %v\n", err)
	}
	return os.Getenv(key)
}

func ToString(field interface{}) string {
	if field == nil {
		return ""
	}
	return fmt.Sprintf("%s", field)
}

func AddIfNotEmpty(name string, value string) string {
	if value == "" {
		return ""
	}
	return fmt.Sprintf("%s=\"%s\"", name, value)

}
func IsEditMode(data model.SitecoreContext) bool {
	fmt.Printf("SitecoreContext.PageState::%s\n", data.PageState)
	return data.PageEditing
}
