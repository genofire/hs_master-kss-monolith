# Getting Started

## 1. Install Docker 

There are several ways to install docker depending on your host operating system (OS). 
Only the [Docker Community Edition (Docker CE)](https://www.docker.com/community-edition) from the stable channel is required.

### Windows 10 Proffesional or Enterprise Edition

If you have Windows 10 Proffesional or Enterprise then Docker will work "nativley" using the internal Hyper-V virtualization technology. 
Follow these [installation instructions](https://store.docker.com/editions/community/docker-ce-desktop-windows?tab=description) 
and check if all is working fine according to [this little guide](https://docs.docker.com/docker-for-windows/).

* Docker tools should be readily available in any shell/terminal 

### Other Windows Edition

Docker does not run "natively" on other Windows versions and editions. 
However, Docker provides the [Docker Toolbox](https://docs.docker.com/toolbox/overview/) to install the neccassry tooling and 
let Docker use the VirtualBox virtualization technology. Follow these [installation instructions](https://docs.docker.com/toolbox/toolbox_install_windows/).

* It will also install a `bash` like Docker Quickstart Terminal
* The Toolbox version requires `bash` like terminals, hence one can also use the Git Bash but needs to setup the environment by running 
   `eval $("docker-machine.exe" env)`
   

### Mac

Here, the choice between newer more "native" installations or use of Docker Toolbox also depends on the version of Mac (Requires OSX Yosemite 10.10.3 or above). 
Then it will be depply integrated with the MacOS Hypervisor framework, networking and filesystem.

Otherwise also install Docker Toolbox following [these instructions](https://docs.docker.com/toolbox/toolbox_install_mac/).

### Linux/Unix

There should be no issue if the distribution uses a recent kernel.


## 2. Run Monolith

1.  Set up an environment variable `LAB_MONOLITH_DB_IP` and assign the docker host's IP.

    For docker machine type `docker-machine env` and the value of `$DOCKER_HOST`,
    otherwise `localhost` is probably correct.

2.  Set up an environment variable `LAB_MONOLITH_DB_PASSWORD` and assign any password. 

3.  `./start-db.sh`

    Starts a PostgreSQL database using the password from `LAB_MONOLITH_DB_PASSWORD`.

4.  `gradle bootRun`
5.  Open http://localhost:8080
6. Happy exploring and di-secting

