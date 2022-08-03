package providerschema

import (
	"io/ioutil"
	"os"
)

func GetPublishedSchema() (string, error) {
	jsonFile, err := os.Open("./schema/provider.json")
	if err != nil {
		return "", err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return "", err
	}

	// because this is an io operation, the schema should be cached upon startup
	return string(byteValue), nil
}
