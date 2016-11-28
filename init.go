package main

import (
    "flag"
    "os"
    "strconv"
)

// ...
var (
    CUDDY_USER_LABEL   string
    CUDDY_GROUP_LABEL  string
    DOCKER_API_VERSION float64
    DOCKER_CERT_PATH   string
    DOCKER_TLS_VERIFY  bool
    DOCKER_HOST        string
)

func init() {
    // ...
    if _default := os.Getenv("CUDDY_GROUP_LABEL"); _default == "" {
        CUDDY_GROUP_LABEL = "io.cuddy.group"
    }

    // ...
    if _default := os.Getenv("CUDDY_USER_LABEL"); _default == "" {
        CUDDY_USER_LABEL = "io.cuddy.user"
    }

    // ...
    if _default := os.Getenv("DOCKER_API_VERSION"); _default == "" {
        DOCKER_API_VERSION = 1.24
    }

    // ...
    if _default := os.Getenv("DOCKER_CERT_PATH"); _default == "" {
        DOCKER_CERT_PATH = ""
    }

    // ...
    if _default := os.Getenv("DOCKER_HOST"); _default == "" {
        DOCKER_HOST = "unix:///var/run/docker.sock"
    }

    // ...
    if _default := os.Getenv("DOCKER_TLS_VERIFY"); _default == "" {
        DOCKER_TLS_VERIFY = false
    }

    // ...
    flag.StringVar(&CUDDY_GROUP_LABEL, "cuddy-group-label", CUDDY_GROUP_LABEL, "Label containing group config")
    flag.StringVar(&CUDDY_USER_LABEL, "cuddy-user-label", CUDDY_USER_LABEL, "Label containing user config")
    flag.Float64Var(&DOCKER_API_VERSION, "docker-api-version", DOCKER_API_VERSION, "Docker API version")
    flag.StringVar(&DOCKER_CERT_PATH, "docker-cert-path", DOCKER_CERT_PATH, "Path to TLS files")
    flag.StringVar(&DOCKER_HOST, "docker-host", DOCKER_HOST, "Daemon socket to connect to")
    flag.BoolVar(&DOCKER_TLS_VERIFY, "docker-tls-verify", DOCKER_TLS_VERIFY, "Use TLS and verify the remote")
    flag.Parse();

    // ...
    os.Setenv("DOCKER_API_VERSION", strconv.FormatFloat(DOCKER_API_VERSION, 'f', 2, 64))
    os.Setenv("DOCKER_CERT_PATH", DOCKER_CERT_PATH)
    os.Setenv("DOCKER_HOST", DOCKER_HOST)
    os.Setenv("DOCKER_TLS_VERIFY", strconv.FormatBool(DOCKER_TLS_VERIFY))
}