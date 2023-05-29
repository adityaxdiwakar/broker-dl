package main

type tomlConfig struct {
	KafkaUrl      string `toml:"kafka_url"`
	NumThreads    int    `toml:"num_threads"`
	RemoteDetails remoteDetails
}

type remoteDetails struct {
	Host     string
	Username string
	Password string
}
