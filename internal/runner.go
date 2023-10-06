package internal

import (
	"bufio"
	"fmt"
	"os"
	"sorohimm/envsync/internal/params"
	"sort"
	"strings"
)

func NewRunner() *Runner {
	return &Runner{}
}

type Runner struct{}

func (o *Runner) Run() {
	var p params.Params
	params.New(&p)

	o.run(&p)
}

func (o *Runner) run(p *params.Params) {
	deploy, err := loadDeployEnv(p.DeployPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	secrets, err := loadSecrets(p.Secrets)
	if err != nil {
		fmt.Println(err)
		return
	}

	writeEnv(deploy, secrets, p.OutputPath)
}

func writeEnv(deploy *Deploy, secrets Secrets, outputPath string) {
	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Create file error: %v\n", err)
		return
	}
	defer file.Close()

	var envs []string

	for envKey, envVal := range deploy.Env {
		line := envKey + "=" + envVal
		envs = append(envs, line)
	}

	for _, services := range deploy.Vault.Services {
		for _, serviceEnvs := range services {
			for dplSecretKey, dplSecretVal := range serviceEnvs {
				if secretVal, ok := secrets[dplSecretVal]; ok {
					line := dplSecretKey + "=" + secretVal
					envs = append(envs, line)
				}
			}
		}
	}

	sort.Strings(envs)

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	_, err = writer.WriteString(strings.Join(envs, "\n"))
	if err != nil {
		fmt.Printf("Write env error: %s", err)
	}
}
