GOOS=linux GOARCH=amd64 go build .

scp adorable-avatars-go-server sreekar339@139.59.93.218:~/adorable-avatars-go-server/adorable-avatars-go-server

ssh -t sreekar339@139.59.93.218 '
sudo systemctl restart adorable-avatars-go-server
'