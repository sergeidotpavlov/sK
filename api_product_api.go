/*
* SystemK
*
* API <br/>  Documentation:    <ul>      <li><a href=\"https://atlassian.net/wiki/</a></li>      <li><a href=\"https://git.net/f\">Git</a></li>      <li><a href=\"https://humansinc.atlassian.net/browse/ORM\">Jira</a></li>    </ul>  <br/>  <a href=\"git .yaml\">API Artifact</a><br/>
*
* API version: 1.0.0
 */
package mware

import (
	"net/http"
)

func CreateUserV12(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func DeleteUserV12(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetUserV12(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateUserV12(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
