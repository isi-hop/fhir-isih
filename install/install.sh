#!/bin/bash

clear

#definir IP du serveur
IPADR=$(ip addr | grep "state UP" -A2 | tail -n1 | cut -f1 -d"/" | cut -f6 -d" ")
#definir hosname
HOSTNAME=$(hostname)
#definir le dossier d'installation 
INSTALLPATH=$(dirname $PWD)

#mettre a jour la distrib avant toutes actions...
echo "****************************"
echo "*   Mise à jour de l'OS    *"
echo "*   et des applications    *"
echo "* mot de passe SUDO requis *"
echo "****************************"
sudo apt update
sudo apt -y upgrade
#verifier la version de l'OS et la version du noyau
#afficher directement à l'écran
echo "**********************"
echo "*  version de l'OS   *"
echo "*  version du noyau  *"
echo "**********************"
cat /etc/os-release
hostnamectl
echo "############################################################"
#verifier la quantite de memoire 
echo "**********************"
echo "*  Quantite memoire  *"
echo "**********************"
free -h
echo "############################################################"
#verifier l espace disque 
echo "**********************"
echo "*   Espace disque    *"
echo "*     disponible     *"
echo "**********************"
df -h
echo "############################################################"
echo "**********************"
echo "*  Installation des  *"
echo "* Applications tiers *"
echo "**********************"
test=$(sudo dpkg --list | grep curl)
if [ -z "$test" ]
then 
#installer curl nécessaire pour la suite des download
sudo apt -y install curl
fi

#test=$(sudo dpkg --list | grep docker)
test=$(ps -ax | grep /usr/bin/dockerd | wc -l)
if [ "$test" -lt 2 ]
then 
#installer docker
curl -fsSL https://get.docker.com -o docker-install.sh
sudo sh docker-install.sh
fi

test=$(sudo dpkg --list | grep docker-compose)
if [ -z "$test" ]
then 
#installer docker-compose
sudo apt -y install docker-compose
fi

echo "**********************"
echo "* Positionnement des *"
echo "* droits pour Docker *"
echo "* et Docker-compose  *"
echo "**********************"
#ajouter les droits pour l'utilisateur
sudo groupadd docker
sudo usermod -aG docker $USER
sudo systemctl restart docker

echo "********************************************"
echo "FIN DE L INSTALLATION-REBOOT NECESSAIRE"
echo " HOSTNAME = $HOSTNAME"
echo " USER = $USER"
echo " IP MACHINE = $IPADR"
echo " INSTALLPATH = $INSTALLPATH"
echo "********************************************"

#pause
echo " "
echo "Reboot su serveur"
echo " "
echo "############################################################"
read -p "Appuyer sur la touche ENTREE pour continuer " REPLY
sudo reboot