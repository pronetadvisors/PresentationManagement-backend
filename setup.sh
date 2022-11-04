##Formatting
##sudo ./setup.sh <DB_Password> <Branch Name> <APP URL> 

###Update all packages###
apt update && apt upgrade -y

###install git###
apt install git -y

###Install go###
echo "********************************"
echo "Installing Go..."
wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
. /etc/profile
echo "************DONE****************"

###Install Node/NPM/PM2/Yarn###
echo "Installing Node..."
curl -fsSL https://deb.nodesource.com/setup_lts.x | sudo -E bash - &&\
sudo apt-get install -y nodejs
echo "************DONE****************"

echo "Updating NPM..."
npm i -g npm@latest
echo "************DONE****************"

echo "Installing PM2..."
npm i -g pm2
pm2 startup
echo "************DONE****************"

echo "Installing Yarn..."
npm i -g yarn
echo "************DONE****************"


###Install MySQL###
echo "Installing MYSQL..."
apt install mysql-server -y
systemctl enable --now mysql
touch mysql_setup_commands.txt
echo "ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '$1';" > mysql_setup_commands.txt
echo "mysql_secure_installation" >> mysql_setup_commands.txt
echo "************DONE****************"


###Fetch Code###
echo "Removing old code..."
rm PresentationManagement-* -r

echo "Fetching code from Github..."
git clone -b $2 https://github.com/pronetadvisors/PresentationManagement-frontend
git clone https://github.com/pronetadvisors/PresentationManagement-backend

echo "Injecting secrets..."
cd PresentationManagement-backend
echo "DB_HOST=localhost" > .env
echo "DB_DRIVER=mysql" >> .env
echo "DB_USER=root" >> .env
echo "DB_PASSWORD=$1" >> .env
echo "DB_NAME=presman_prod" >> .env
echo "DB_PORT=3306" >> .env
echo "BUCKET_PATH=./storage/" >> .env
echo "APP_URL=$3" >> .env
echo "PORT=8080" >> .env

go build main.go
pm2 start main
pm2 save
cd ..

cd PresentationManagement-frontend
yarn
yarn build
rm /var/www/dist -r
mv dist/ /var/www
cd ..

echo "************DONE****************"



###Configure Nginx###
echo "Installing nginx..."
apt install nginx -y
echo "************DONE****************"

###CLEAN UP###
echo "Cleaning up..."
rm go1.19.2.linux-amd64.tar.gz
sudo apt autoremove
echo "************DONE****************"