package helper

var version = "1.0.0" // cannot be accessed from outside
var Application = "golang"

func sayGoodbye(name string) string {
	return "Goodbye, " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
