package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/docker/go-plugins-helpers/volume"
)

type s3fsDriver struct {
	defaultS3fsopts string
	Driver
}

func (p *s3fsDriver) Validate(req *volume.CreateRequest) error {
	return nil
}

func (p *s3fsDriver) MountOptions(req *volume.CreateRequest) ([]string, error) {
	s3fsopts, s3fsoptsInOpts := req.Options["o"]
	bucket, bucketInOpts := req.Options["bucket"]
	folder, folderInOpts := req.Options["folder"]

	if !bucketInOpts {
		return nil, errors.New("driver option 'bucket' is mandatory")
	}

	var s3fsoptsArray []string
	if s3fsoptsInOpts && s3fsopts != "" {
		s3fsoptsArray = append(s3fsoptsArray, strings.Split(s3fsopts, ",")...)
	} else if p.defaultS3fsopts != "" {
		s3fsoptsArray = append(s3fsoptsArray, strings.Split(p.defaultS3fsopts, ",")...)
	}
	bucketOption := "bucket=" + bucket
	if folderInOpts {
		bucketOption = bucketOption + ":/" + folder
	}
	s3fsoptsArray = append(s3fsoptsArray, bucketOption)

	return []string{"-o", strings.Join(s3fsoptsArray, ",")}, nil
}

func (p *s3fsDriver) PreMount(req *volume.MountRequest) error {
	return nil
}

func (p *s3fsDriver) PostMount(req *volume.MountRequest) {
}

func buildDriver() *s3fsDriver {
	defaultsopts := os.Getenv("DEFAULT_S3FSOPTS")
	d := &s3fsDriver{
		Driver:          *NewDriver("s3fs", false, "s3fs", "local"),
		defaultS3fsopts: defaultsopts,
	}
	d.Init(d)
	return d
}

func spawnSyslog() {
	cmd := exec.Command("rsyslogd", "-n")
	cmd.Start()
}

func main() {
	spawnSyslog()
	//log.SetFlags(0)
	d := buildDriver()
	defer d.Close()
	d.ServeUnix()
}
