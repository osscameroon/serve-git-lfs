# SERVE-GIT-LFS

This is a small golang service than can serve for you git-lfs large files.
You just have to provide the configuration file properly and you're good to go.

## DEMO

https://user-images.githubusercontent.com/22576758/171308009-a99fb94a-b455-4966-9c5a-4de24dd5edb8.mp4


## How to install

**/!\NOTE**: always make sure to have a clean shared directory before running sglfs `rm ./shared/*`

- First you need to copy the configuratio example and set our parameters:
  ```bash
  cp example.conf.yml conf.yml
  ```
  The structure is quite simple:
  ```yaml
  storage:
    - path: "audio-files"
      url: "https://github.com/osscameroon/podcasts"
      branch: "master"
    - path: "videos"
      url: "https://github.com/osscameroon/docjocoding"
      branch: "main"
  ```

- The you can build the docker image
  ```bash
  make docker-build
  # or with no cache
  make docker-build-no-cache
  ```

## How to launch

Just run `go run main.go` or build it manually.

You can also use docker here:
```
make docker-run
```

A server should be running on port :3000 !
