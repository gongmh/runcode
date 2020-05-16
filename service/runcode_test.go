package service

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func TestDockerLs(t *testing.T) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

	return
}

func TestRunShellCode(t *testing.T) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	reader, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stdout, reader)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}

func TestRunShellFile(t *testing.T) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "golang:alpine",
		Cmd:   []string{"sh", "-c", "sh /go/1.sh"},
		//Cmd: []string{"sh", "-c", "which go"},
		Tty: true,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	var goDemoCode = "package main\n\nfunc main() { \n    println(\"hello golang\")\n}\n"
	shellContent := fmt.Sprintf("#!/bin/bash \necho \"%s\" > 1.go \n/usr/local/go/bin/go run 1.go\nls", strings.ReplaceAll(goDemoCode, `"`, `\"`))

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	err = tw.WriteHeader(&tar.Header{
		Name: "1.sh",                   // filename
		Mode: 0777,                     // permissions
		Size: int64(len(shellContent)), // filesize
	})
	fmt.Println(err)
	tw.Write([]byte(shellContent))
	tw.Close()

	//fmt.Println(buf.String())

	opt := types.CopyToContainerOptions{
		//AllowOverwriteDirWithFile: true,
	}
	err = cli.CopyToContainer(ctx, resp.ID, "/go/", &buf, opt)
	fmt.Println(err)

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	reader, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stdout, reader)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}

func TestRunGolangCode(t *testing.T) {
	var goDemoCode = "package main\n\nfunc main() { \n    println(\"hello golang\")\n}\n"
	//fileName := "file"

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	//fmt.Println(fmt.Sprintf("ls && cat> ./text.txt<EOF '%s' EOF && ls -la ./", goDemoCode))

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "golang:alpine",
		//Cmd:   []string{"sh", "-c", "go build /tmp/1.go && ./1 2>> tmp && cat tmp"},
		//Cmd: []string{"sh", "-c", "go run 1.go"},
		//Cmd:          []string{"go", "run", "/tmp/1.go"},
		Cmd:          []string{"sh", "-c", "ls && cat 1.go && ls -la ./"},
		Tty:          true,
		AttachStderr: true,
		AttachStdout: true,
		//Shell:        []string{"sh", "-c", "ls && cat /tmp/1.go && ls -la /tmp"},
		Entrypoint: []string{},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	err = tw.WriteHeader(&tar.Header{
		Name: "1.go",                 // filename
		Mode: 0777,                   // permissions
		Size: int64(len(goDemoCode)), // filesize
	})
	fmt.Println(err)
	tw.Write([]byte(goDemoCode))
	tw.Close()

	//fmt.Println(buf.String())

	opt := types.CopyToContainerOptions{
		//AllowOverwriteDirWithFile: true,
	}
	err = cli.CopyToContainer(ctx, resp.ID, "/go/", &buf, opt)
	fmt.Println(err)

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	reader, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		panic(err)
	}

	//_, err = io.Copy(os.Stdout, reader)
	//if err != nil && err != io.EOF {
	//	log.Fatal(err)
	//}

	bs, err := ioutil.ReadAll(reader)
	fmt.Println(string(bs))
}
