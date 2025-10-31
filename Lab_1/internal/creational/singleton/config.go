package singleton

type configuration struct {
    Value string
}

var instance *configuration

func GetConfig() *configuration {
    if instance == nil {
        instance = &configuration{Value: "Bloomify Default Config"}
    }
    return instance
}
