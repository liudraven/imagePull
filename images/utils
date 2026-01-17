package images

import (
	"github.com/containers/image/v5/signature"
	"github.com/containers/image/v5/types"
)

func systemContext(harborUser, harborPasswd string) *types.SystemContext {
	ctx := &types.SystemContext{
		ArchitectureChoice:          "amd64",
		OSChoice:                    "linux",
		VariantChoice:               "",
		DockerRegistryUserAgent:     "",
		DockerInsecureSkipTLSVerify: types.NewOptionalBool(true),
		DockerCertPath:              "",
		DockerAuthConfig: &types.DockerAuthConfig{
			Username: harborUser,
			Password: harborPasswd,
		},
	}
	return ctx
}

func getPolicyContext() (*signature.PolicyContext, error) {
	policy := &signature.Policy{Default: []signature.PolicyRequirement{signature.NewPRInsecureAcceptAnything()}}
	return signature.NewPolicyContext(policy)
}


