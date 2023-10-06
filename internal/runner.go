package internal

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/sorohimm/envsync/internal/params"
)

type Envs []string

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
	cfg, err := loadConfig(p.ConfigPath)
	if err != nil {
		fmt.Println("Warning: config is not loaded!")
	}

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

	envs := makeEnvs(deploy, secrets, cfg.Replace)

	err = writeEnv(envs, p.OutputPath)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Done.")
}

func writeEnv(envs Envs, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create file error: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	_, err = writer.WriteString(strings.Join(envs, "\n"))
	if err != nil {
		return fmt.Errorf("write env error: %w", err)
	}

	return nil
}

func makeEnvs(deploy *Deploy, secrets Secrets, replace Replace) []string {
	var envsRaw = make(map[string]string)

	for envKey, envVal := range deploy.Env {
		envsRaw[envKey] = envVal
	}

	for _, services := range deploy.Vault.Services {
		for _, serviceEnvs := range services {
			for dplSecretKey, dplSecretVal := range serviceEnvs {
				if secretVal, ok := secrets[dplSecretVal]; ok {
					envsRaw[dplSecretKey] = secretVal
				}
			}
		}
	}

	if len(replace) != 0 {
		for replaceKey, replaceVal := range replace {
			envsRaw[replaceKey] = replaceVal
		}
	}

	var envs []string
	for k, v := range envsRaw {
		envs = append(envs, k+"="+v)
	}

	sort.Strings(envs)

	return envs
}
