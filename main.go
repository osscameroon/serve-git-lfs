package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type StorageUnit struct {
	Path   string `yaml:"path"`
	Url    string `yaml:"url"`
	Branch string `yaml:"branch"`
}

type Config struct {
	Storage []StorageUnit `yaml:"storage"`
}

var SHAREDDIR = "shared"
var CONFPATH = "conf.yml"

func readConf() Config {
	filename, _ := filepath.Abs(CONFPATH)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Panic(err)
	}
	return config
}

func execCommand(bin string, targetDir string, args ...string) {
	cmd := exec.Command(bin, args...)
	// because we execute all our commands inside the shared directory
	cmd.Dir = targetDir

	stdout, err := cmd.Output()

	if err != nil {
		log.Panic(err)
	}

	log.Print(bin, args)
	// Print the output
	log.Print(string(stdout))
}

func appendToFile(path string, content string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		log.Println(err)
	}
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func gitClone(path string, url string, repoName string, branch string) {
	// We want to clone only the subdirectory not the whole project
	execCommand("git", SHAREDDIR, "init", repoName)

	commands := [...]string{
		"git remote add origin " + url,
		"git config core.sparsecheckout true",
		"git-lfs install",
	}
	for _, cmdd := range commands {
		binn := strings.Split(cmdd, " ")[0]
		args := strings.Split(cmdd, " ")[1:]
		execCommand(binn, SHAREDDIR+"/"+repoName, args...)
	}
	// this will help us cloning only the directory we want from
	// the repo and nothing else
	appendToFile(SHAREDDIR+"/"+repoName+"/.git/info/sparse-checkout", path)
	execCommand("git", SHAREDDIR+"/"+repoName, "pull", "origin", branch)
}

func gitPull(path string, url string, branch string) {
	// we parse to get the repo name
	urlParts := strings.Split(url, "/")
	repoName := urlParts[len(urlParts)-1]

	// a simple chef if the dir exists or not
	// this will help us decide either we pull or clone
	res, err := exists(SHAREDDIR + "/" + repoName)
	if err != nil {
		log.Panic(err)
	}

	if res {
		execCommand("git", SHAREDDIR+"/"+repoName, "pull", "origin", branch)
	} else {
		gitClone(path, url, repoName, branch)
	}
}

func refreshData() {
	// we loop over configuration points
	// we pull if the directories exists and we clone if not
	for _, store := range readConf().Storage {
		gitPull(store.Path, store.Url, store.Branch)
		log.Println("[-] Sleeping state for 24 Hours...")
		time.Sleep(24 * time.Hour)
	}
}

func main() {
	res, err := exists(CONFPATH)
	if err != nil {
		log.Panic(err)
	}
	if res {
		execCommand("git-lfs", ".", "--version")
		execCommand("git", ".", "--version")
		go refreshData()

		fs := http.FileServer(http.Dir(SHAREDDIR))
		http.Handle("/", fs)

		log.Print("[-] sglfs Listening on :3000...")
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("[x] Please set Up your configuration file !")
	}
}
