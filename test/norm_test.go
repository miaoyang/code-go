package test

import (
	"fmt"
	"log"
	"testing"
)

type Man struct {
	height int
	weight int
	sex    int
	age    int
}

func (this *Man) NewMan(height, weight, sex, age int) *Man {
	return &Man{
		height: height,
		weight: weight,
		sex:    sex,
		age:    age,
	}
}

func (this *Man) PrintMan() {
	log.Printf("[Man] height: %d,weight: %d, sex: %d, age: %d\n", this.height, this.weight, this.sex, this.age)
}

type Student struct {
	Man
	score int
	level int
	name  string
}

func NewStudent(score, level int, name string) *Student {
	return &Student{
		score: score,
		level: level,
		name:  name,
	}
}

type Teacher struct {
	Man
	class string
	name  string
}

func NewTeacher(class, name string) *Teacher {
	return &Teacher{
		class: class,
		name:  name,
	}
}

func (this *Teacher) PrintTeacher() {
	log.Printf("[Teacher] class: %s, name: %s\n", this.class, this.name)
}

func TestSlice(t *testing.T) {
	a := make([]int, 5, 10)
	a = append(a, 1, 2, 3, 4, 5, 6, 7)
	log.Println(a)
	log.Println(a[6:len(a)])
}

func TestMan(t *testing.T) {
	stu := Student{}
	stu.weight = 120
	stu.height = 180
	stu.age = 20
	stu.sex = 1
	stu.PrintMan()

	teacher := NewTeacher("二年级", "二年级一班")
	teacher.PrintTeacher()
}

type Context interface {
	GetName() string
}

func (this *Teacher) GetName() string {
	return "Teacher"
}

func (this *Student) GetName() string {
	return "Student"
}

func GetNameTest(ctx Context) {
	name := ctx.GetName()
	fmt.Println(name)
}

func TestTeacherInterface(t *testing.T) {
	teacher := NewTeacher("三年级", "三年级一班")
	GetNameTest(teacher)
}
