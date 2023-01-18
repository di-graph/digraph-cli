package formatting

import (
	"fmt"
	"regexp"
)

const sanitize_replacement_string string = "REDACTED"

var sensitiveFieldNames = []string{
	"(account)*(.[a-z0-9_]*)*key",
	"(api)*(.[a-z0-9_]*)*key",
	"(auth)*(.[a-z0-9_]*)*key",
	"azurerm-account-key",
	"(client)*(.[a-z0-9_]*)*key",
	"contrasena",
	"contraseña",
	"(database)*(.[a-z0-9_]*)*key",
	"(database)*(.[a-z0-9_]*)*pass",
	"(database)*(.[a-z0-9_]*)*password",
	"(db)*(.[a-z0-9_-]+)*key",
	"(db)*(.[a-z0-9_]*)*pass",
	"(db)*(.[a-z0-9_]*)*password",
	"fetch-tfstate-headers",
	"(key)*(.[a-z0-9_]*)*pass",
	"(key)*(.[a-z0-9_]*)*password",
	"passwd",
	"password",
	"(priv)*(.[a-z0-9_]*)*key",
	"(private)*(.[a-z0-9_]*)*key",
	"pwd",
	"secret",
	"(service)*(.[a-z0-9_]*)*key",
	"tfc-token",
	"username",
}

func SanitizeValuesByKey(content []byte) ([]byte, error) {
	for i := range sensitiveFieldNames {
		filter := sensitiveFieldNames[i]
		r, err := regexp.Compile(fmt.Sprintf(`(?:\"|\')(?P<key>(%s)+)(?:\"|\')(?:\:\s*)(?:\"|\')?(?P<value>[\w\s-\[\]]*)(?:\"|\')?`, filter))
		if err != nil {
			fmt.Printf("error in regex %s", err.Error())
			return nil, err
		}
		template := fmt.Sprintf("\"$key\":\"%s\"", sanitize_replacement_string)
		content = r.ReplaceAll(content, []byte(template))
	}
	return content, nil
}
