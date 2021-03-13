package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Action struct {
	Action  string `json:"action"`
	ObjName string `json:"object"`
}

type Teacher struct {
	ID        string   `json:"id"`
	Salary    float64  `json:"salary"`
	Subject   string   `json:"subject"`
	Classroom []string `json:"classroom"`
	Person    struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

type Stuff struct {
	ID        string   `json:"id"`
	Salary    float64  `json:"salary"`
	Classroom []string `json:"classroom"`
	Person    struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

type Student struct {
	ID     string  `json:"id"`
	Salary float64 `json:"salary"`
	Class  string  `json:"class"`
	Person struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

var obj_sl []GeneralObject

func GetId(o GeneralObject) string {
	switch o.(type) {
	case *Teacher:
		t := o.(Teacher)
		return t.ID
	case *Stuff:
		s := o.(Stuff)
		return s.ID
	case *Student:
		st := o.(Student)
		return st.ID

	}
	return ""
}

func (t Teacher) GetCreateAction() DefinedAction {
	return &CreateTeacher{}
}
func (t Teacher) GetUpdateAction() DefinedAction {
	return &UpdateTeacher{}
}
func (t Teacher) GetReadAction() DefinedAction {
	return &ReadTeacher{}
}
func (t Teacher) GetDeleteAction() DefinedAction {
	return &DeleteTeacher{}
}

func (s Stuff) GetCreateAction() DefinedAction {
	return &CreateStuff{}
}
func (s Stuff) GetUpdateAction() DefinedAction {
	return &UpdateStuff{}
}
func (s Stuff) GetReadAction() DefinedAction {
	return &ReadStuff{}
}
func (s Stuff) GetDeleteAction() DefinedAction {
	return &DeleteStuff{}
}

func (st Student) GetCreateAction() DefinedAction {
	return &CreateStudent{}
}
func (st Student) GetUpdateAction() DefinedAction {
	return &UpdateStudent{}
}
func (st Student) GetReadAction() DefinedAction {
	return &ReadStudent{}
}
func (st Student) GetDeleteAction() DefinedAction {
	return &DeleteStudent{}
}

type DefinedAction interface {
	GetFromJSON([]byte)
	Process()
}

type GeneralObject interface {
	GetCreateAction() DefinedAction
	GetUpdateAction() DefinedAction
	GetReadAction() DefinedAction
	GetDeleteAction() DefinedAction
}

type CreateTeacher struct {
	T Teacher `json:"data"`
}

func (action *CreateTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action CreateTeacher) Process() {
	obj_sl = append(obj_sl, action.T)
}

type UpdateTeacher struct {
	T Teacher `json:"data"`
}

func (action *UpdateTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action UpdateTeacher) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if action.T.ID == GetId(obj_sl[i]) {
			obj_sl[i] = action.T
		}
	}
}

type ReadTeacher struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (action *ReadTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action ReadTeacher) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if action.Data.ID == GetId(obj_sl[i]) {
			fmt.Println(obj_sl[i])
		}
	}
}

type DeleteTeacher struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (action *DeleteTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action DeleteTeacher) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if GetId(obj_sl[i]) == action.Data.ID {
			obj_sl[i] = nil
		}
	}
}

type CreateStuff struct {
	S Stuff `json:"data"`
}

func (action *CreateStuff) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action CreateStuff) Process() {
	obj_sl = append(obj_sl, action.S)
}

type UpdateStuff struct {
	S Stuff `json:"data"`
}

func (action *UpdateStuff) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action UpdateStuff) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if action.S.ID == GetId(obj_sl[i]) {
			obj_sl[i] = action.S
		}
	}
}

type ReadStuff struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (action *ReadStuff) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action ReadStuff) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if action.Data.ID == GetId(obj_sl[i]) {
			fmt.Println(obj_sl[i])
		}
	}
}

type DeleteStuff struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (action *DeleteStuff) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action DeleteStuff) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if GetId(obj_sl[i]) == action.Data.ID {
			obj_sl[i] = nil
		}
	}
}

type CreateStudent struct {
	St Student `json:"data"`
}

func (action *CreateStudent) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action CreateStudent) Process() {
	obj_sl = append(obj_sl, action.St)
}

type UpdateStudent struct {
	St Student `json:"data"`
}

func (action *UpdateStudent) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action UpdateStudent) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if action.St.ID == GetId(obj_sl[i]) {
			obj_sl[i] = action.St
		}
	}
}

type ReadStudent struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (action *ReadStudent) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action ReadStudent) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if action.Data.ID == GetId(obj_sl[i]) {
			fmt.Println(obj_sl[i])
		}
	}
}

type DeleteStudent struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (action *DeleteStudent) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action DeleteStudent) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if GetId(obj_sl[i]) == action.Data.ID {
			obj_sl[i] = nil
		}
	}
}

var Ch = make(chan int, 1)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:12667")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	Ch <- 1

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go HandleConnection(conn)
	}

}

func HandleConnection(conn net.Conn) {
	buf := make([]byte, 2000)
	n, err := conn.Read(buf)
	if err != nil {
		conn.Close()
		return
	}

	<-Ch

	//	for scanner.Scan() {
	var act Action
	err = json.Unmarshal(buf[:n], &act)
	if err != nil {
		fmt.Println(err)
		return
	}

	var obj GeneralObject
	switch act.ObjName {
	case "Teacher":
		obj = &Teacher{}
	case "Stuff":
		obj = &Stuff{}
	case "Student":
		obj = &Student{}

	}
	var toDo DefinedAction
	switch act.Action {
	case "create":
		toDo = obj.GetCreateAction()
		toDo.Process()
	case "update":
		toDo = obj.GetUpdateAction()
		toDo.Process()
	case "read":
		toDo = obj.GetReadAction()
		toDo.Process()
	}
	//toDo.GetFromJSON(file)

	toDo.Process()
	//}
	Ch <- 1

	conn.Close()

	/*	data := []byte("Connection great")
		conn.Write(data)*/

}
