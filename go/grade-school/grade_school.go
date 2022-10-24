package school

import "sort"

// School is school type
type School []Grade

// Grade is grade type
type Grade struct {
	level    int
	students []string
}

func New() *School {
	return new(School)
}

func (s *School) Add(student string, level int) {
	for i, grade := range *s {
		if grade.level == level {
			// todo search if already exists
			(*s)[i].students = append((*s)[i].students, student)
			return
		}
	}
	*s = append(*s, Grade{level: level, students: []string{student}})
}

func (s *School) Grade(level int) []string {
	for _, grade := range *s {
		if grade.level == level {
			return grade.students
		}
	}
	return []string{}
}

func (s *School) Enrollment() (grades []Grade) {
	for i := range *s {
		sort.Slice((*s)[i].students, func(a, b int) bool { return (*s)[i].students[a] < (*s)[i].students[b] })
	}
	sort.Slice(*s, func(a, b int) bool { return (*s)[a].level < (*s)[b].level })
	return *s
}
