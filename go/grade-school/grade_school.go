package school

// School is school type
type School map[string]int

// Grade is grade type
type Grade struct {
	level    int
	students []string
}

func New() *School {
	return new(School)
}

func (s *School) Add(student string, grade int) {
	(*s)[student] = grade
}

func (s *School) Grade(level int) (students []string) {
	for student, l := range *s {
		if level == l {
			students = append(students, student)
		}
	}
	return
}

func (s *School) Enrollment() (grades []Grade) {

	panic("Please implement the Enrollment function")
}
