#!/bin/bash
yum install -y yum-utils   device-mapper-persistent-data   lvm2
yum-config-manager     --add-repo     https://download.docker.com/linux/centos/docker-ce.repo
yum install docker-ce-18.09.9 docker-ce-cli-18.09.9 containerd.io -y
systemctl start docker
systemctl enable docker
yum -y install bash-completion
source /etc/profile.d/bash_completion.sh
mkdir -p /etc/docker
tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://v16stybc.mirror.aliyuncs.com"]
}
EOF
systemctl daemon-reload
systemctl restart docker