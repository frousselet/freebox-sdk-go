package freebox

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 * Request authorization. The Freebox owner will have to validate the app on the Freebox LCD.
 *
 * @param  {string} app_id			string The ID of the App
 * @param  {string} app_name		string The name of the App
 * @param  {string} app_version	string The version of the App
 * @param  {string} device_name	string The name of the client device
 */
func RequestAuthorization(app_id string, app_name string, app_version string, device_name string) string {
	resp, err := http.Get("https://api.github.com/users/tensorflow")
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	fmt.Print(string(body))

	return ("temp")
}
