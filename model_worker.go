/*
 * SystemK
 *
 * API <br/>  Documentation:    <ul>      <li><a href=\"https://atlassian.net/wiki/</a></li>      <li><a href=\"https://git.net/f\">Git</a></li>      <li><a href=\"https://humansinc.atlassian.net/browse/ORM\">Jira</a></li>    </ul>  <br/>  <a href=\"git .yaml\">API Artifact</a><br/>
 *
 * API version: 1.0.0
 */
package mware

type Worker struct {
	// id должности
	WorkPositionId string `json:"WorkPositionId"`
	// id работника
	UserId string `json:"UserId,omitempty"`
}
