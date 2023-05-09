package main

type UserData struct {
	Name string
}

func main() {
	var info UserData
	info.Name = "WilburXu"
	_ = GetUserInfo(info)
}

func GetUserInfo(userInfo UserData) *UserData {
	return &userInfo
}

// go run -gcflags '-m -l' main.go
