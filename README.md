Site Go
=======

Simple [Golang](https://golang.org/) website with [Echo Framework](https://echo.labstack.com/) and [Hero templates](https://shiyanhui.github.io/hero) with [Boostrap 4](https://getbootstrap.com/).

The goal is to create a very simple but fully functional [customisable](customise.md) stateless website which can be used to test multiple [deployments](deployments.md), like [Vagrant](vagrant.md) VMs, Ansible Playbooks, Terraform scripts, Docker containers and Kubernetes clusters with minimal overhead and no runtimes.

Features:

* Static build on multiple platforms (see Release page);
* Working with as little as only the executable;
* Configurable with command line parameters and environment variables;
* Use logging (to console, but can be captured to files);
* [Custom images and pages](customise.md);

Installation
------------

You may need to install project dependencies, like:

```
go get -u github.com/labstack/echo/...
go get -u github.com/shiyanhui/hero/hero
go get -u golang.org/x/tools/cmd/goimports
```

Compile from source
-------------------

Compile Hero templates

```
hero -source="./template"
```

Compile SiteGo for Linux on Windows (you may need to close command prompt and reopen it)

```
setx GOOS linux
go build -o sitego main.go
```

Or on Windows (you may need to set GOOS environment variable back to 'windows'):

```
go build -o sitego.exe main.go
```

Check your Golang environment variables with *go env*

Usage
-----

To run the application

```
sitego[.exe] [params]
```

Get all parameters

```
sitego -help
```

Command line parameters (optional):

```
Usage of sitego:
  -b    hide Echo banner
  -i    logo image (default false)
  -l    log to console
  -p string
        port (default "8000")
  -s string
        site name (default "Site Go")
  -version
        show version
```

Example:

```
sitego -s "My Test" -p 8080 -l
``` 

Environment variables

* SITE_NAME
* SITE_PORT
* SITE_LOG = missing or "0" means logging is not active
* SITE_LOGO = missing or "0" means logo image is not displayed

IMPORTANT: If both command line parameters and environment variables are defined the parameters take precedence. 
For instance, if *SITE_LOG = 1* and *sitego -l=0*, the logging is not active.

Deployments
-----------

* [Vagrant](vagrant.md)
* [Ansible Playbook](ansible.md) for AWS and Azure
* [Terraform](terraform.md) scripts for AWS and Azure
* [Docker](docker.md)
* [Kubernetes](kubernetes.md)

Notes
-----

* Host name is displayed on the footer;
* Source-code include generated Hero templates;
* If you create new environment variables you may need to reboot your computer on Windows; Or you can use *setx SITE_LOGO 1* in command prompt, then close the window and open it again before restarting *sitego.exe*;
* ASP.NET Core version is also available at [github.com/OviCrisan/SiteNet](https://github.com/ovicrisan/SiteNet);

Versions

* 1.0 - initial release

