package machine

type MachinesSpec struct {
	Machines []MachineSpec `json:"machines"`
}

type MachineSpec struct {
	Name           string          `json:"name"`
	Reflectors     []ReflectorSpec `json:"refls"`
	Rotors         []RotorSpec     `json:"rots"`
	RotorsCount    int             `json:"rotsCount"`
	IsHasPlugboard bool            `json:"isHasPb"`
}

type ReflectorSpec struct {
	Name string `json:"name"`
}

type RotorSpec struct {
	Name      string `json:"name"`
	Poses     byte   `json:"poses"`
	RingPoses byte   `json:"ringPoses"`
}
