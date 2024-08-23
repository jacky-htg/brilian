package lib

var MyStr string = "Selamat "

type Salam struct {
	Message string
	Type    string
}

func (u *Salam) Sapa() string {
	// db.Close()
	return u.Message + " " + u.Type
}
