echo "[START TEST]\n"
echo "\n=========Start /session/ test============"
curl -X 'POST' \
  'http://localhost:3000/session' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "tux"
}'

curl -X 'POST' \
  'http://localhost:3000/session' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "markus"
}'

curl -X 'POST' \
  'http://localhost:3000/session' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "vainor"
}'

curl -X 'POST' \
  'http://localhost:3000/session' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "opensuse"
}'

curl -X 'POST' \
  'http://localhost:3000/session' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "nitrux"
}'

curl -X 'POST' \
  'http://localhost:3000/session' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "fedora"
}'

curl -X 'POST' \
  'http://localhost:3000/session' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "timux"
}'
sleep 2
echo "\n=======end /session/ test=============\n"
echo "\n=======start /users/ test=============\n"


curl -X 'GET' \
  'http://localhost:3000/users/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 12'
echo "\t\n\t\n"
curl -X 'GET' \
  'http://localhost:3000/users/?username=v' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 3'
echo "\t\n\t\n"
curl -X 'GET' \
  'http://localhost:3000/users/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 2'
echo "\t\n\t\n"
curl -X 'GET' \
  'http://localhost:3000/users/?username=a' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 1'
sleep 2
echo "\n========end test /users/================\n"
echo "\n========start test /users/:uid================\n"


curl -X 'GET' \
  'http://localhost:3000/users/2/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 1'
echo "\t\n\t\n"
curl -X 'GET' \
  'http://localhost:3000/users/3/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 1'
echo "\t\n\t\n"
curl -X 'GET' \
  'http://localhost:3000/users/1/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 2'
echo "\t\n\t\n"
curl -X 'GET' \
  'http://localhost:3000/users/3/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 4'
echo "\t\n\t\n"
sleep 2
echo "\n========end test /users/:uid/================\n"
echo "\n========star test /users/:uid/followers/followerId================\n"


curl -X 'PUT' \
  'http://localhost:3000/users/2/followers/3' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 3'
echo "\t\n\t\n"
curl -X 'PUT' \
  'http://localhost:3000/users/3/followers/4' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 4'
echo "\t\n\t\n"
curl -X 'PUT' \
  'http://localhost:3000/users/1/followers/2' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 2'
echo "\t\n\t\n"
curl -X 'PUT' \
  'http://localhost:3000/users/3/followers/4' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 4'
 
 echo \n dell \n 
  
curl -X 'DELETE' \
  'http://localhost:3000/users/1/followers/2' \
  -H 'accept: */*' \
  -H 'Authorization: Bearer 2'
  

  

echo "\n========end test /users/:uid/followers/followerId================\n"
echo "\n========start test /users/:uid/followers================\n"
curl -X 'GET' \
  'http://localhost:3000/users/2/followers/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 1'
echo "\t\n\t\n"
curl -X 'GET' \
  'http://localhost:3000/users/3/followers/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 1'
echo "\t\n\t\n"
curl -X 'GET' \
  'http://localhost:3000/users/1/followers/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 2'
echo "\t\n\t\n"
curl -X 'GET' \
  'http://localhost:3000/users/3/followers/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 4'
sleep 2
echo "\n========end test /users/:uid/followers/================\n"
echo "\n========start test /users/:uid/banned/:bannedId================\n"

curl -X 'PUT' \
  'http://localhost:3000/users/2/banned/3' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer 2'\
  -d '{
  "username": "marioross"
}'
curl -X 'PUT' \
  'http://localhost:3000/users/2/banned/5' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer 2'\
  -d '{
  "username": "marioross"
}'

curl -X 'PUT' \
  'http://localhost:3000/users/3/banned/1' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer 3'\
  -d '{
  "username": "marioross"
}'

curl -X 'PUT' \
  'http://localhost:3000/users/5/banned/4' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer 1'\
  -d '{
  "username": "marioross"
}'

echo \nput\n
curl -X 'GET' \
  'http://localhost:3000/users/2/banned/?limit=50' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 2'

echo \ndel\n

curl -X 'DELETE' \
  'http://localhost:3000/users/3/banned/1' \
  -H 'accept: */*' \
  -H 'Authorization: Bearer 3'


  

echo "\n========start test /users/:uid/myPhotos/================\n"

curl --output outFile -X 'POST' \
  'http://localhost:3000/users/2/myPhotos/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 2' \
  -H 'Content-Type: multipart/form-data' \
  -F 'imageData=@big.png;type=image/png'\
  -F 'descriptionImg="Foto spaziale!'\

curl  -X 'POST' \
  'http://localhost:3000/users/3/myPhotos/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 3' \
  -H 'Content-Type: multipart/form-data' \
  -F 'imageData=@tiny.png;type=image/png'\
  -F 'descriptionImg="Foto spaziale!'\

#curl --output outFile -X 'POST' \
#  'http://localhost:3000/users/5/myPhotos/' \
#  -H 'accept: application/json' \
#  -H 'Authorization: Bearer 4' \
#  -H 'Content-Type: multipart/form-data' \
#  -F 'imageData=@tiny.png;type=image/png'\
#  -F 'descriptionImg="Foto spaziale!'\

sleep 2
echo \n=============================\(start test /users/:uid/myPhotos/ \)===============================\n 


curl -X 'GET' \
  'http://localhost:3000/users/3/myPhotos/?limit=50&sortBy=dataUpdate&ordering=desc' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 2'


  sleep 2 


  echo \n=============================\(start test /users/:uid/myPhotos/:photoId/comments/ \)===============================\n 

  curl -X 'POST' \
  'http://localhost:3000/users/1/myPhotos/1/comments/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 2' \
  -H 'Content-Type: application/json' \
  -d '{
  "user": {
    "username": "markus"
  },
  "text": "😀 i like you photo! 😀"
}'

curl -X 'POST' \
  'http://localhost:3000/users/1/myPhotos/1/comments/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 4' \
  -H 'Content-Type: application/json' \
  -d '{
  "user": {
    "username": "opensuse"
  },
  "text": "😀"
}'

curl -X 'POST' \
  'http://localhost:3000/users/1/myPhotos/1/comments/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 7' \
  -H 'Content-Type: application/json' \
  -d '{
  "user": {
    "username": "timux"
  },
  "text": "Go😀d"
}'

curl -X 'POST' \
  'http://localhost:3000/users/1/myPhotos/1/comments/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 5' \
  -H 'Content-Type: application/json' \
  -d '{
  "user": {
    "username": "marioross"
  },
  "text": "best photo"
}'
