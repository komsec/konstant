# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  config.vm.box_check_update = false
  config.vm.define "n1"
  # config.vm.define "n2"
  config.vm.box = "centos/7"
  # config.vm.synced_folder "../data", "/vagrant_data"
  config.vm.synced_folder "../", "/code"

  net_prefix = ENV['NET_PREFIX'] || "192.168.100.0"
  config.vm.network "private_network", :type => :dhcp, :ip => net_prefix, :netmask => "255.255.255.0"
  # Enable provisioning with a shell script. Additional provisioners such as
  # Puppet, Chef, Ansible, Salt, and Docker are also available. Please see the
  # documentation for more information about their specific syntax and use.
  # config.vm.provision "shell", inline: <<-SHELL
  #   apt-get update
  #   apt-get install -y apache2
  # SHELL
end
