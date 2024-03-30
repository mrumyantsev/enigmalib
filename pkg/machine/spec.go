package machine

type MachinesSpec struct {
	Machines []MachineSpec `json:"machines"`
}

type MachineSpec struct {
	Name           string          `json:"name"`
	Reflectors     []ReflectorSpec `json:"refls,omitempty"`
	Rotors         []RotorSpec     `json:"rots,omitempty"`
	RotorsCount    int             `json:"rotsCount,omitempty"`
	IsHasPlugboard bool            `json:"isHasPb,omitempty"`
}

type ReflectorSpec struct {
	Name string `json:"name"`
}

type RotorSpec struct {
	Name      string `json:"name"`
	Poses     byte   `json:"poses"`
	RingPoses byte   `json:"ringPoses"`
}
