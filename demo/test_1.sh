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
sleep 5
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
sleep 5
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
sleep 5
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

curl -X 'PUT' \
  'http://localhost:3000/users/5/followers/2' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 2'

curl -X 'PUT' \
  'http://localhost:3000/users/3/followers/2' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 2'



curl -X 'DELETE' \
  'http://localhost:3000/users/1/followers/2' \
  -H 'accept: */*' \
  -H 'Authorization: Bearer 2'
  

  

echo "\n========end test /users/:uid/followers/followerId================\n"
echo "\n========start test /users/:uid/followers================\n"
curl -X 'GET' \
  'http://localhost:3000/users/1/followers/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 1'
echo "\t\n\t\n"
curl -X 'GET' \
  'http://localhost:3000/users/2/followers/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 2'
echo "\t\n\t\n"
curl -X 'GET' \
  'http://localhost:3000/users/1/followers/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 2'
echo "\t\n\t\n"
curl -X 'GET' \
  'http://localhost:3000/users/3/followers/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 3'
sleep 5
echo "\n========end test /users/:uid/followers/================\n"
echo "\n========start test /users/:uid/banned/:bannedId================\n"


curl -X 'PUT' \
  'http://localhost:3000/users/5/banned/4' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer 1'\
  -d '{
  "username": "marioross"
}'


curl -X 'GET' \
  'http://localhost:3000/users/5/banned/?limit=50' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 5'

curl -X 'DELETE' \
  'http://localhost:3000/users/3/banned/1' \
  -H 'accept: */*' \
  -H 'Authorization: Bearer 3'



echo "\n========end test /users/:uid/banned/:bannedId================\n"

echo "\n=======Start upload Photo ===================================\n"


curl -X 'POST' \
  'http://localhost:3000/users/1/myPhotos/' \
  -H 'accept: application/json' \
  -H 'Content-Type: multipart/form-data' \
  -H 'Authorization: Bearer 1' \
  -F 'descriptionImg="Una foto spazialeeeee!" '\
  -F 'imageData=@big.png'

sleep 5

echo "\n========start test /users/:uid/MyPhotos/:photoId/comments/=====\n"

curl -X 'POST' \
  'http://localhost:3000/users/1/myPhotos/1/comments/' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "author": {
    "username": "marioross"
  },
  "text": "😀 i like you photo! 😀"
}'

curl -X 'POST' \
  'http://localhost:3000/users/1/myPhotos/1/comments/' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "author": {
    "username": "marioross"
  },
  "text": "😀 i like you photo! 😀"
}'
