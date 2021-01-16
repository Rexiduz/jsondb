package jsondb

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

/*********
* MODELS
*********/

// User user-model
type User struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Age  byte   `json:"age"`
}

// ManagerUser store helper
type ManagerUser struct {
	store Store
}

/*********
* METHODS
*********/

func (m ManagerUser) getDataAsBytes(list []User) []byte {
	byt, err := json.MarshalIndent(list, "", "    ")

	if err != nil {
		fmt.Println("Error [getDataAsBytes]", err)
		return []byte(m.store.DefaultValue)
	}

	return byt
}

func (m ManagerUser) updateUserList(list []User) {
	byt := m.getDataAsBytes(list)
	m.store.Write(byt)
}

func (m ManagerUser) getUserList() []User {
	byt, err := m.store.Read()
	if err != nil {
		byt = []byte(m.store.DefaultValue)
	}

	var list []User

	e2 := json.Unmarshal(byt, &list)

	if e2 != nil {
		fmt.Println("Error ", err)
	}

	return list
}

/*********
* VARIABLES
*********/
var storage = Store{Path: "static/db.json", DefaultValue: `[]`}

/*********
* FUNCTIONS
*********/

// GenerateID random unit32
func GenerateID() uint32 {
	buf := make([]byte, 8)
	rand.Read(buf) // Always succeeds, no need to check error
	return binary.LittleEndian.Uint32(buf)
}

// AddUser add user data to db
func AddUser(name string, age byte) {
	m := ManagerUser{store: storage}

	list := m.getUserList()
	list = append(list, User{ID: GenerateID(), Name: name, Age: age})

	m.updateUserList(list)
}

// PrintList show all user data
func PrintList() {
	m := ManagerUser{store: storage}
	byt := m.getDataAsBytes(m.getUserList())

	fmt.Println(string(byt))
}
