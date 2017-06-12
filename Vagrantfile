# -*- mode: ruby -*-
# vi: set ft=ruby :
# See https://github.com/discourse/discourse/blob/master/docs/VAGRANT.md
#
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box= "ubuntu/trusty64"

  config.vm.network :private_network, ip: "192.168.33.35"  
  
  config.vm.provider :virtualbox do |v|
    v.customize ["modifyvm", :id, "--memory", "512"]
    v.customize ["modifyvm", :id, "--cpus", "1"]
    v.customize ["modifyvm", :id, "--natdnsproxy1", "on"]
    v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
    v.customize ["modifyvm", :id, "--nictype1", "virtio"]
  end

  config.vm.synced_folder ".", "/home/vagrant/ECS-GO",
    mount_options: ["dmode=775,fmode=664"] 
 
  config.vm.provision "shell", path: "scripts/script.sh"  
end