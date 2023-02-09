package dictionary_test

import (
	"github.com/mesh-master/foundation/pkg/ancillary/struc/dictionary"
	"reflect"
	"strings"
	"testing"
	"time"
)

type (
	UppercaseString string
	Parentheses     string
	ExtraInfo       struct {
		Enrollment time.Time
	}
)

type PersonInterface interface {
	GetName() string
	GetAge() uint
}

type StudentInterface interface {
	GetYear() uint
	GetFaculty() string
}

type Person struct {
	dictionary.Dictionary
	Name UppercaseString `name:"name"`
	Age  uint            `name:"age"`
}

// Student extends the Person dictionary by embedding its pointer.
type Student struct {
	dictionary.DictionaryInterface
	Year      uint        `name:"year"`
	Faculty   Parentheses `name:"faculty"`
	ExtraInfo ExtraInfo   `name:"extra"`
}

func (p Person) GetName() string {
	return string(p.Name)
}

func (p Person) GetAge() uint {
	return p.Age
}

func (s Student) GetYear() uint {
	return s.Year
}

func (s Student) GetFaculty() string {
	return string(s.Faculty)
}

const (
	name    = "Andrei Samuilik"
	age     = 39 // alas not a constant in real life
	faculty = "chemical"
)

func TestImport(t *testing.T) {
	now := time.Now()
	person := &Person{Name: name, Age: age}
	upperCaseImp := func(target dictionary.DictionaryInterface, name, alias string, v reflect.Value) error {
		v.SetString(strings.ToUpper(v.String()))
		return nil
	}
	parenthesesImp := func(target dictionary.DictionaryInterface, name, alias string, v reflect.Value) error {
		v.SetString("(" + v.String() + ")")
		return nil
	}
	extraImp := func(target dictionary.DictionaryInterface, name, alias string, v reflect.Value) error {
		src := ExtraInfo{Enrollment: now}
		v.Set(reflect.ValueOf(src))
		return nil
	}
	student := &Student{
		Year:    1,
		Faculty: faculty,
	}
	student.DictionaryInterface = person
	dictionary.RegisterTypeHandlers(reflect.TypeOf((*UppercaseString)(nil)).Elem(), upperCaseImp, nil)
	dictionary.RegisterTypeHandlers(reflect.TypeOf((*Parentheses)(nil)).Elem(), parenthesesImp, nil)
	dictionary.RegisterTypeHandlers(reflect.TypeOf((*ExtraInfo)(nil)).Elem(), extraImp, nil)
	err := dictionary.Dictionary{}.Import(student)
	if err != nil {
		t.Fatalf("import: %v", err)
	}
	var expected string
	expected = strings.ToUpper(name)
	p := student.DictionaryInterface.(PersonInterface)
	if p.GetName() != expected {
		t.Fatalf("expected %s, got %s", expected, p.GetName())
	}
	expected = "(" + faculty + ")"
	if student.GetFaculty() != expected {
		t.Fatalf("expected %s, got %s", expected, student.GetFaculty())
	}
	if student.ExtraInfo.Enrollment != now {
		t.Fatalf("expected %v, got %v", now, student.ExtraInfo.Enrollment)
	}
}
