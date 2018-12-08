echo "1"
sudo systemctl stop adorable-avatars-go-server
echo "2"
rm -rf ~/adorable-avatars-go-server/adorable-avatars-go-server
echo "3"
mv ~/adorable-avatars-go-server-tmp ~/adorable-avatars-go-server/adorable-avatars-go-server
echo "4"
sudo systemctl start adorable-avatars-go-server
echo "5"