package swag

type Envs envs

type envs struct {
	Router map[string]string
}

func GlobalEnvs(envs map[string]string) {
	Envs.Router = envs
}
