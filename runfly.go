package main

import (
    "log"
    "os"
    "os/exec"
    "fmt"
)

func main() {

	// command to start the pcf pipeline using fly command
	fmt.Printf("Apply Configuration changes for setting up pipeline (y/N) :")
    cmd := exec.Command("fly","-tlocal","set-pipeline","-pgo-pcf","-c/home/ubuntu/vasanth/aws/pipeline.yml","-l/home/ubuntu/vasanth/aws/params.yml")
	cmd.Stdin = os.Stdin
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	// command to unpause the pipeline
	un_pause_cmd := exec.Command("fly","-tlocal","unpause-pipeline", "-pgo-pcf")
	un_pause_cmd.Stdin = os.Stdin
	un_pause_out, un_pause_err := un_pause_cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("unpause_cmd.Run() failed with %s\n", un_pause_err)
	}
	fmt.Printf("combined out:\n%s\n", string(un_pause_out))

	trigger_cmd := exec.Command("fly","-tlocal","trigger-job", "-jgo-pcf/bootstrap-terraform-state")
	trigger_cmd.Stdin = os.Stdin
	trigger_out, trigger_err := trigger_cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("trigger_cmd.Run() failed with %s\n", trigger_err)
	}
	fmt.Printf("combined out:\n%s\n", string(trigger_out))
}