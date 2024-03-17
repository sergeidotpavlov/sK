/*
 * SystemK
 *
 * API <br/>  Documentation:    <ul>      <li><a href=\"https://atlassian.net/wiki/</a></li>      <li><a href=\"https://git.net/f\">Git</a></li>      <li><a href=\"https://humansinc.atlassian.net/browse/ORM\">Jira</a></li>    </ul>  <br/>  <a href=\"git .yaml\">API Artifact</a><br/>
 *
 * API version: 1.0.0
 */
package swagger

type CreateUserResponse struct {
	// уникальный идентификатор клиента
	Id string `db:"id"`
	// логин
	Username string `db:"username"`
	// Фамилия*
	FirstName string `db:"firstName"`
	// Имя*
	LastName string `db:"lastName"`
	// Отчество
	SecondName string `db:"secondName"`
	// дата рождения unix time в микросекундах
	BirhDate int32 `db:"birhDate"`
	// Электронная почта*
	Email string `db:"email"`
	// Пароль храним как хеш
	Password string `db:"password"`
	// телефон
	Phone string `db:"phone"`
	// User Status
	UserStatus string `db:"statusId"`
	// Подразделение
	SubDivision string `db:"subdivision"`
}
