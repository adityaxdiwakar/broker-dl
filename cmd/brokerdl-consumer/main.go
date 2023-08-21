package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/BurntSushi/toml"
	brokerdl "github.com/adityaxdiwakar/broker-dl"
)

var conf tomlConfig

func init() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatalf("error: could not parse configuration: %v\n", err)
	}
}

func genRemoteExec(location, name string) string {
	name = strings.ReplaceAll(name, " ", "\\ ")
	name = strings.ReplaceAll(name, "'", "\'")

	remoteLoc := fmt.Sprintf("sftp://%s:%s@%s%s/%s",
		conf.RemoteDetails.Username, conf.RemoteDetails.Password,
		conf.RemoteDetails.Host, location, name)

	if strings.HasSuffix(name, "mkv") {
		return fmt.Sprintf("pget -n %d -c %s", conf.NumThreads, remoteLoc)
	}

	return fmt.Sprintf("mirror --use-pget-n=%d -c %s", conf.NumThreads, remoteLoc)
}

func main() {
	r := brokerdl.GetKafkaReader(conf.KafkaUrl)

	fmt.Println("Starting consuming...")
	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}

		var notification brokerdl.DownloadNotification
		json.Unmarshal(m.Value, &notification)

		fmt.Printf("Processing file %s\n", notification.Name)

		/* remote command to execute for transfer */
		remoteCommand := genRemoteExec(notification.Location, notification.Name)

		/* begin transfer and wait for completion */
		cmd := exec.Command("lftp", "-c", remoteCommand)
		cmd.Dir = conf.Locations.Incompletes
		err = cmd.Run()
		if err != nil {
			fmt.Printf("Skipping %s due to download failure\n", notification.Name)
			continue
		}

		from := fmt.Sprintf("%s%s", conf.Locations.Incompletes, notification.Name)
		to := fmt.Sprintf("%s%s", conf.Locations.Completes, notification.Name)
		if err = os.Rename(from, to); err != nil {
			fmt.Printf("Skipping %s due to move failure\n", notification.Name)
			continue
		}

		fmt.Printf("Marked %s as completed\n", notification.Name)
		r.CommitMessages(ctx, m)
	}
}
