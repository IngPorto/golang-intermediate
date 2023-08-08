package employee

type Employee struct {
	Id       int
	Name     string
	Vacation bool
}

// reciver function
// Un struct de tipo Employee va a poseer un método llamado SetId
func (e *Employee) SetId(id int) {
	e.Id = id
}

func (e *Employee) SetName(n string) {
	e.Name = n
}

func (e *Employee) GetId() int {
	return e.Id
}

func (e *Employee) GetName() string {
	return e.Name
}