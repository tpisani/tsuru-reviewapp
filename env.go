package main

import (
	"fmt"
)

type EnvVar struct {
	Name    string
	Value   string
	Public  bool
	Inherit bool
}

type Env map[string]EnvVar

func MergeEnvs(baseEnv Env, newEnv Env) (Env, []error) {
	var errs []error
	env := make(Env)

	for name, envVar := range newEnv {
		value := envVar.Value

		if envVar.Inherit {
			v, exists := baseEnv[name]
			if !exists {
				errs = append(errs, fmt.Errorf("%s does not exist on base env", name))
				continue
			}

			if !v.Public {
				errs = append(errs, fmt.Errorf("%s is a private variable", name))
				continue
			}

			value = v.Value
		}

		env[name] = EnvVar{
			Name:    name,
			Value:   value,
			Public:  true,
			Inherit: false,
		}
	}

	return env, errs
}
