package management

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kkirsche/kratos/executable"
	"github.com/kkirsche/kratos/utilities/logs"
)

// UpdateKratosGithub is used to allow kratos to update itself in place using
// a public repository.
// theoretically this may be picked up by a AV
func UpdateKratosGithub(username string, repository string) {
	var releaseInfo KratosRelease

	currentVersion := executable.Version
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", username, repository)
	resp, err := http.Get(url)
	if err != nil {
		logs.PrintlnError("failed to retrieve the latest kratos release with error", err)
		return
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&releaseInfo)

	fmt.Println(currentVersion, releaseInfo.TagName)
}
