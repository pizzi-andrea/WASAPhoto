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
echo "\n========end test /users/:uid/================\n"
echo "\n========star test /users/:uid/followers/followerId================\n"


curl -X 'PUT' \
  'http://localhost:3000/users/2/followers/3' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 1'
echo "\t\n\t\n"
curl -X 'PUT' \
  'http://localhost:3000/users/3/followers/4' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 1'
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
echo "\n========start test /users/:uid/followers/================\n"
