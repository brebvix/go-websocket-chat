package main

type Command struct {
	action string
	params []string
	client *Client
}

func (co Command) run() (message string) {
	print(co.action)
	switch co.action {
	case "/name":
		return co.name(co.params[1])
	}

	return ""
}

func (co Command) name(name string) (message string) {
	var resultMessage = string(co.client.name) + " change name to " + name

	co.client.name = []byte(name)

	return resultMessage
}
