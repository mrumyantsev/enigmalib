package machine

type MachinesSpec struct {
	Machines []MachineSpec `json:"machines"`
}

type MachineSpec struct {
	Name            string          `json:"name"`
	Reflectors      []ReflectorSpec `json:"reflectors,omitempty"`
	Rotors          []RotorSpec     `json:"rotors,omitempty"`
	ReflectorsCount int             `json:"reflectorsCount,omitempty"`
	RotorsCount     int             `json:"rotorsCount,omitempty"`
	PlugboardsCount int             `json:"plugboardsCount,omitempty"`
}

type ReflectorSpec struct {
	Name string `json:"name"`
}

type RotorSpec struct {
	Name          string `json:"name"`
	Positions     byte   `json:"positions"`
	RingPositions byte   `json:"ringPositions"`
}
