package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hokou      string
	Xingzuo    string
	House      string
	Car        string
}

func FromJsonObj(obj interface{}) (Profile, error) {
	var profile Profile
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &profile)
	return profile, err
}
