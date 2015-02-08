package main

import (
	"github.com/Unknwon/goconfig"
	"log"
	"os"
	"os/exec"
	"strings"
)

type NmapParam struct {
	path       string
	target     string
	port_range string
	sub_param  string
	stdout     string
	stderr     string
}

func NewNampParam() *NmapParam {
	return &NmapParam{
		path:       "nmap",
		target:     "192.168.0.1",
		port_range: "80,443",
		sub_param:  "-O",
		stdout:     "stdout.log",
		stderr:     "stderr.log"}
}

func main() {
	np := NewNampParam()
	np.LoadConf("conf.ini")
	np.Run()
}

func (this *NmapParam) SaveConf(conf string) {
	cfg, err := goconfig.LoadFromData([]byte(``))
	if err != nil {
		log.Panicln(err)
	}
	cfg.SetValue("default", "nmap_path", this.path)
	cfg.SetValue("default", "target", this.target)
	cfg.SetValue("default", "port_range", this.port_range)
	cfg.SetValue("default", "sub_param", this.sub_param)
	cfg.SetValue("default", "stdout", this.stdout)
	cfg.SetValue("default", "stderr", this.stderr)
	goconfig.SaveConfigFile(cfg, conf)
}

func (this *NmapParam) LoadConf(conf string) {
	cfg, err := goconfig.LoadConfigFile(conf)
	if err != nil {
		log.Println("conf is not exist, so create it")
		this.SaveConf(conf)
		return
	}
	this.path = cfg.MustValue("default", "nmap_path", "nmap")
	this.target = cfg.MustValue("default", "target", "192.168.0.1")
	this.port_range = cfg.MustValue("default", "port_range", "80,443")
	this.sub_param = cfg.MustValue("default", "sub_param", "-O")
	this.stdout = cfg.MustValue("default", "stdout", "stdout.log")
	this.stderr = cfg.MustValue("default", "stderr", "stderr.log")
}

func (this *NmapParam) Run() {
	fdout, err := os.Create(this.stdout)
	if err != nil {
		log.Panicln(err)
	}
	defer fdout.Close()

	fderr, err := os.Create(this.stderr)
	if err != nil {
		log.Panicln(err)
	}
	defer fderr.Close()

	argArray := strings.Split(this.sub_param, " ")
	argArray = append(argArray, "-p", this.port_range, this.target)

	cmd := exec.Command(this.path, argArray...)

	cmd.Stderr = fderr
	cmd.Stdout = fdout

	log.Println(cmd.Args)
	cmd.Run()
}
