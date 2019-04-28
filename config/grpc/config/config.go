package config

type Config struct {
	Micro *Micro `json:"micro,omitempty"`
	Extra *Extra `json:"other,omitempty"`
}

type Micro struct {
	Info
}

type Extra struct {
	Info
}

type Info struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Message string `json:"message,omitempty"`
	Age     int    `json:"age,omitempty"`
}
