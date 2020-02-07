package config

import (
    "os"
    "github.com/arturoguerra/d2arena/internal/structs"
)

func LoadgRPC() *structs.GRPC {
    return &structs.GRPC{
        FaceitHost: os.Getenv("GRPC_FACEIT_HOST"),
        FaceitPort: os.Getenv("GRPC_FACEIT_PORT"),
        ProfilesHost: os.Getenv("GRPC_PROFILES_HOST"),
        ProfilesPort: os.Getenv("GRPC_PROFILES_PORT"),
    }
}
