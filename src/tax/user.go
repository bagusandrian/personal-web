package tax

import (
	"log"
)

func (m *Module) RegisterUser(name string) (result UserData, err error) {
	rows, err := m.queries.InsertUser.Query(name)
	if err != nil {
		log.Printf("[tax][RegisterUser] queries InsertUser err: %+v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Name)
		if err != nil {
			log.Printf("[tax][RegisterUser] failed scan data : %+v", err)
			return
		}
	}
	return
}

func (m *Module) GetListUser() (result []UserData, err error) {
	var user UserData
	rows, err := m.queries.GetlistUsers.Query()
	if err != nil {
		log.Printf("[tax][GetListUser] queries GetlistUsers err: %+v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Printf("[tax][GetListUser] failed scan data : %+v", err)
			return
		}
		result = append(result, user)
	}
	return
}
