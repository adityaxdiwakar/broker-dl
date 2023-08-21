package main

type tomlConfig struct {
	KafkaUrl      string `toml:"kafka_url"`
	NumThreads    int    `toml:"num_threads"`
	RemoteDetails remoteDetails
	Locations     locations
	DebugLevel    string `toml:"debug_level"`
}

type remoteDetails struct {
	Host     string
	Username string
	Password string
}

type locations struct {
	Incompletes string
	Completes   string
}
