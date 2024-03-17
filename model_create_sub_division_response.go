/*
* SystemK
*
* API <br/>  Documentation:    <ul>      <li><a href=\"https://atlassian.net/wiki/</a></li>      <li><a href=\"https://git.net/f\">Git</a></li>      <li><a href=\"https://humansinc.atlassian.net/browse/ORM\">Jira</a></li>    </ul>  <br/>  <a href=\"git .yaml\">API Artifact</a><br/>
*
* API version: 1.0.0
 */
package swagger

type CreateSubDivisionResponse struct {
	// уникальный идентификатор подразделения
	Id string `db:"id"`
	// уникальное наименование подразделения
	SubDivision string `db:"SubDivision"`
	// компания
	CompanyId string `db:"CompanyId"`
}
