package lib

import "fmt"

type Unit struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Player  string `json:"player"`
	Spawned bool   `json:"spawned"`
	Died    bool   `json:"dies"`
}

func (u Unit) Self() bool     { return u.Player == "0" }
func (u Unit) Leader() bool   { return u.Id == 1 || u.Id == 2 }
func (u Unit) String() string { return fmt.Sprintf("%s | %t | %t", u.Type, u.Spawned, u.Died) }

type Units []Unit

func (us Units) Len() int           { return len(us) }
func (us Units) Swap(i, j int)      { us[i], us[j] = us[j], us[i] }
func (us Units) Less(i, j int) bool { return us[i].Id < us[j].Id }
