package main


type User struct {
	Name string
	Age  int
}

func main() {

	/* 	user := User{
	   		Name: "Muhammad",
	   		Age: 0,
	   	}

	   	temp, err := template.New("test").Parse("My name is {{.Name}} and{{if .Age}}  I am {{.Age}} years old{{else}} Your age is unknown {{end}}\n")
	   	if err != nil{
	   		panic(err)
	   	}
	   	err = temp.Execute(os.Stdout, user)
	   	if err != nil{
	   		panic(err)
	   	} */

	/* data := []string{"Muhammad", "Ahmad", "Tom"}

	temp := `{{range .}}Hello {{.}}! {{end}}`

	t, err := template.New("test").Parse(temp)
	if err != nil{
		panic(err)
	}

	t.Execute(os.Stdout, data)
	*/

	/* user := User{
		Name: "Muhammad",
		Age:  20,
	}

	temp := `{{with .}}
	Name : {{.Name}}
	Age : {{.Age}}
	{{end}}` */




}
