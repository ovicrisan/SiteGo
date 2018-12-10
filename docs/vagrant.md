Vagrant
=======

```
mkdir vagrant_sitego
cd vagrant_sitego
vagrant init bento/ubuntu-18.04
vim Vagrantfile
```
You can manually add / edit these lines:

```
   config.vm.network "forwarded_port", guest: 8000, host: 8080
   config.vm.provider "virtualbox" do |vb|
     vb.memory = "512"
   end
   config.vm.provision "shell", run: "always", inline: "/vagrant/sitego -i"
```

Or you can just download Vagrantfile as well as other resources (logo and favicon)

```
wget https://github.com/ovicrisan/SiteGo/blob/master/deployment/Vagrantfile
wget https://github.com/ovicrisan/SiteGo/releases/download/1.0/sitego
chmod +x sitego
mkdir static
wget -P static https://github.com/ovicrisan/SiteGo/blob/master/static/logo.png
wget -P static https://github.com/ovicrisan/SiteGo/blob/master/static/favicon.ico
```

Then just build and start VM with *vagrant up*

Open *localhost:8080* on your browser to see the website (including logo).
