// package main

// type Logic interface {
// 	Process(data string) string
// }

// type LogicProvider struct{}

// func (lp LogicProvider) Process(data string) string {
// 	return data
// }

// type Client struct {
// 	L Logic
// }

// func (c Client) Program() {
// 	data := "get from somewhere"
// 	c.L.Process(data)
// }

// func main() {
// 	c := Client{
// 		L: LogicProvider{},
// 	}
// 	c.Program()
// }
