GOOS=linux GOARCH=amd64 go build .

scp adorable-avatars-go-server sreekar339@139.59.93.218:~/adorable-avatars-go-server-tmp

ssh -t sreekar339@139.59.93.218 '
echo "1"
sudo systemctl stop adorable-avatars-go-server
echo "2"
rm -rf ~/adorable-avatars-go-server/adorable-avatars-go-server
echo "3"
mv ~/adorable-avatars-go-server-tmp ~/adorable-avatars-go-server/adorable-avatars-go-server
echo "4"
sudo systemctl start adorable-avatars-go-server
echo "5"
'