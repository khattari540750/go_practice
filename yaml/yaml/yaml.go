package yaml

import (
	my_yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

// structたち
type Data struct {
	Users []User `yaml:"users"`
}

type User struct {
	Name             string           `yaml:"common"`
	FullName         fullName         `yaml:"full_name"`
	Sex              string           `yaml:"sex"`
	SelfIntroduction selfIntroduction `yaml:"self_introduction"`
	ImageURLs        []string         `yaml:"image_urls"`
	Shemale          bool             `yaml:"shemale"`
}

type fullName struct {
	FirstName string `yaml:"first_name"`
	LastName  string `yaml:"last_name"`
}

type selfIntroduction struct {
	Long  string `yaml:"long"`
	Short string `yaml:"short"`
}


func ReadYamlFile(fileName string) (Data, error) {
	buf, err := ioutil.ReadFile(fileName)
	var d Data
	err = my_yaml.Unmarshal(buf, &d)
	return d, err
}


// func main() {

// 	k, e :=  ReadYamlFile("sample.yaml")
// 	if e != nil {
// 		return
// 	}
// 	fmt.Println(k.Users[0].Sex)
// 	fmt.Println(k.Users[1].Sex)
// }
