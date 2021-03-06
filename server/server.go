package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Data struct {
	N  string
	Ch chan int
}

type IdData struct {
	id int
}

type ObjData struct {
	obj interface{}
}

type Teacher struct {
	Firstname string
	Lastname  string
	Id        int
	Subject   string
	Classroom int
}

type Student struct {
	Firstname string
	Lastname  string
	Id        int
	Class     int
}

type Stuff struct {
	Firstname string
	Lastname  string
	Id        int
	Phone     string
	Classroom int
}

type School interface {
	PrintAll()
	GetId() int
}

var school []School
var dat Data

func Add(s *[]School, per School) {
	*s = append(*s, per)
}

func Remove(s []School, id int) {
	for i := 0; i < len(s); i++ {
		if s[i].GetId() == id {
			s[i] = nil
		}
	}
}

func (t *Teacher) GetId() int {
	return t.Id
}

func (stud *Student) GetId() int {
	return stud.Id
}

func (st *Stuff) GetId() int {
	return st.Id
}

func listAll(s []School) {
	for i := 0; i < len(s); i++ {
		if s[i] != nil {
			s[i].PrintAll()
		}
	}
}
func (st *Stuff) PrintAll() {
	fmt.Println("Stuff : ", st.Firstname, st.Lastname, st.Id, st.Phone, st.Classroom)
}
func (stud *Student) PrintAll() {
	fmt.Println("Student : ", stud.Firstname, stud.Lastname, stud.Id, stud.Class)
}

func (t *Teacher) PrintAll() {
	fmt.Println("Teacher : ", t.Firstname, t.Lastname, t.Id, t.Subject, t.Classroom)
}

func ListStud(s []School) {
	for i := 0; i < len(s); i++ {
		switch s[i].(type) {
		case *Student:
			s[i].PrintAll()
		}
	}
}

func ListTeacher(s []School) {
	for i := 0; i < len(s); i++ {
		switch s[i].(type) {
		case *Teacher:
			s[i].PrintAll()

		}
	}
}
func main() {
	l, err := net.Listen("tcp", "127.0.0.1:12667")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	dat.Ch = make(chan int, 1)
	dat.Ch <- 1

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

	<-dat.Ch

	fmt.Println(string(buf[:n]))
	dat.N = string(buf[:n])
	switch dat.N {
	case "delete":
		fmt.Println("Input ID : ")
		var IdDat IdData
		fmt.Println(string(buf[:n]))
		IdDat.id = int(binary.BigEndian.Uint64(buf[:n]))
		Remove(school, IdDat.id)
	case "add":
		fmt.Println("Input person : ")
		var ObjDat ObjData
		fmt.Println(string(buf[:n]))
		ObjDat.obj = School(interface{}(buf[:n]))
		Add(&school, ObjDat.obj)
	case "list all":
		listAll(school)
	case "list students":
		ListStud(school)
	case "list teachers":
		ListTeacher(school)
	default:
		fmt.Println("Incorect input ")
	}
	/*	data := []byte("Connection great")
		conn.Write(data)*/

	dat.Ch <- 1

	conn.Close()
}
