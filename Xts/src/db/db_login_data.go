package db


const Register string = "insert into login_data(userid,password)values(?, ?)"

const FindPassword string = " select password from login_data where userid=? "

type login_data struct {
	Pid int `db:"Pid"`
	Userid string `db:"userid"`
	Password string `db:"password"`
}
