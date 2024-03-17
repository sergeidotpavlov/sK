/*
 * SystemK
 *
 * API <br/>  Documentation:    <ul>      <li><a href=\"https://atlassian.net/wiki/</a></li>      <li><a href=\"https://git.net/f\">Git</a></li>      <li><a href=\"https://humansinc.atlassian.net/browse/ORM\">Jira</a></li>    </ul>  <br/>  <a href=\"git .yaml\">API Artifact</a><br/>
 *
 * API version: 1.0.0
 */
package mware

type UserParams struct {
	// логин
	Username string `json:"username,omitempty"`
	// Фамилия*
	FirstName string `json:"firstName,omitempty"`
	// Имя*
	LastName string `json:"lastName,omitempty"`
	// Отчество
	SecondName string `json:"secondName,omitempty"`
	// дата рождения unix time в микросекундах
	BirhDate int32 `json:"birhDate,omitempty"`
	// Электронная почта*
	Email string `json:"email,omitempty"`
	// Пароль храним как хеш
	Password string `json:"password,omitempty"`
	// телефон
	Phone string `json:"phone,omitempty"`
	// User Status
	UserStatus string `json:"userStatus,omitempty"`
	// Подразделение
	SubDivision string `json:"subDivision,omitempty"`
}
