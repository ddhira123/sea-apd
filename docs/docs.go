package docs

import (
	"github.com/swaggo/swag"
	"io/ioutil"
	"log"
	"os"
)

type s struct{}

func (s *s) ReadDoc() string {
	file, err := os.Open("docs/doc.json")
	if err != nil {
		log.Panic(err)
	}
	str, _ := ioutil.ReadAll(file)
	return string(str)
}
func init() {
	swag.Register(swag.Name, &s{})
}