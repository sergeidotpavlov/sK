/*
 * SystemK
 *
 * API <br/>  Documentation:    <ul>      <li><a href=\"https://atlassian.net/wiki/</a></li>      <li><a href=\"https://git.net/f\">Git</a></li>      <li><a href=\"https://humansinc.atlassian.net/browse/ORM\">Jira</a></li>    </ul>  <br/>  <a href=\"git .yaml\">API Artifact</a><br/>
 *
 * API version: 1.0.0
 */
package swagger

type CreateWorkPositionResponse struct {
	// уникальный идентификатор
	Id string `db:"id"`
	// уникальное наименование должности
	WorkPosition string `db:"name"`
	// id подразделения
	SubDivisionId string `db:"subDivisionId"`
}
