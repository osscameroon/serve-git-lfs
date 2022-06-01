# SERVE-GIT-LFS

This is a small golang service than can serve for you [git-lfs](https://git-lfs.github.com/) large files.
You just have to provide the configuration file properly and you're good to go.

## DEMO

https://user-images.githubusercontent.com/22576758/171308009-a99fb94a-b455-4966-9c5a-4de24dd5edb8.mp4


## HOW DOES IT WORKS

**sglfs** will parse the configuration file you are providing, then try to pull directories from
repositories you provided, if those directories/repos are not valids, it will try to clone those.
There is a small loop inside it for refreshing your shared folder depending on the configurations.


## HOW TO INSTALL

After cloning the repository : `git clone https://github.com/osscameroon/serve-git-lfs`...

**/!\ NOTE**: Always make sure to have a clean shared directory before running sglfs `rm ./shared/*` or at least keep the same gitconfig user !

- First you need to copy the configuration example file and set your parameters:
  ```bash
  cp example.conf.yml conf.yml
  ```
  The structure is quite simple, an array of storage elements:
  ```yaml
  storage:
    - path: "audio-files"
      url: "https://github.com/osscameroon/podcasts"
      branch: "master"
    - path: "videos"
      url: "https://github.com/osscameroon/docjocoding"
      branch: "main"
  ```
  - _**path**_: The directory you want to get from the repo.

    EX: `sub_dir1/sub_dir2`
  - _**url**_: The url of the repository.

    EX: `https://github.com/djunior/repoUltime`
  - _**branch**_: The branch name from which you want to clone the repo.

    EX: `master` or `release`, depends on you.

- Second, you can build the docker image
  ```bash
  make docker-build
  # or with no cache
  make docker-build-no-cache
  ```
  or just build the golang service with `make build`


## HOW TO LAUNCH

Just do `make run-dev` or if it's already build : `make run-prod`.

You can also use docker for that:
```bash
make docker-run
```

**/!\ NOTE**: Either you start the service with docker, either you start with the default CLI, since **./shared** is a volume, you should not do git operations with the docker container at the same time

Now, a server should be running on port :3000 !
```
2022/06/01 07:03:26 git-lfs[--version]
2022/06/01 07:03:26 git-lfs/3.1.4 (GitHub; linux amd64; go 1.18.2)
2022/06/01 07:03:26 git[--version]
2022/06/01 07:03:26 git version 2.34.2
2022/06/01 07:03:26 [-] sglfs Listening on :3000...
2022/06/01 07:03:26 git[init podcasts]
2022/06/01 07:03:26 Initialized empty Git repository in /shared/podcasts/.git/
2022/06/01 07:03:26 git[remote add origin https://github.com/osscameroon/podcasts]
2022/06/01 07:03:26 
2022/06/01 07:03:26 git[config core.sparsecheckout true]
2022/06/01 07:03:26 
2022/06/01 07:03:26 git-lfs[install]
2022/06/01 07:03:26 Updated Git hooks.
Git LFS initialized.
2022/06/01 07:03:34 git[pull origin master]
2022/06/01 07:03:34 
2022/06/01 07:03:34 [-] Sleeping state for 24 Hours...
```
