package main

import "fmt"

type Robot struct {
	model       string
	serialId    int
	workCounter int
}

func (r Robot) String() string {
	return fmt.Sprintf("Robot %s serialID %d", r.model, r.serialId)
}

func (r *Robot) Work(tasks []string) string {
	res := fmt.Sprintf("%s work:", r)
	for _, task := range tasks {
		res += "\n I do " + task
	}
	r.workCounter += len(tasks)
	return res
}

type Person struct {
	name     string
	homework string
	children []*Person
}

func (p Person) DoHomework() string {
	return p.homework
}

func (p Person) Children() []*Person {
	return p.children
}

func (p Person) Work(tasks []string) string {
	s := p.name + " work:"
	for _, task := range tasks {
		s += "\n I do " + task
	}
	return s
}

func (p Person) String() string {
	return p.name
}

type Worker interface {
	Work(tasks []string) string
}

type Company struct {
	personal []Worker
}

func (c *Company) Hire(newbie Worker) {
	c.personal = append(c.personal, newbie)
}

func (c Company) Progress(id int, tasks []string) (res string) {
	return c.personal[id].Work(tasks)
}

func main() {
	pers := Person{}
	comp := Company{}
	comp.Hire(pers)
}
