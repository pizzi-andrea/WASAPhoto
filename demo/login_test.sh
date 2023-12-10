#!/bin/bash
clear
username=("debian" "kali" "opensuse" "ubuntu" "fedora" "mangeia" "knopix" "garuda")
echo	"[Login Test]"
for u in "${username[@]}"
do
	curl -X 'POST' \
		'http://localhost:3000/session' \
		-H 'accept: application/json' \
		-H 'Content-Type: application/json' \
		-d '{"name": "'"$u"'"}'
	echo ":$u"
done

echo "[end Login test]"

echo "[start following test]"


for i in {1..10}
do
	from=$(expr 1 + $RANDOM % "${#username[@]}")
	to=$(expr 1 + $RANDOM % "${#username[@]}")

	if [ $from != $to ]
	then
		curl -X 'PUT' \
			"http://localhost:3000/users/$to/followers/$from" \
			-H 'accept: application/json' \
			-H "Authorization: Bearer $from"
	fi
done

echo "[end following test]"

echo "[start post photo]"

for i in {1..120}
do
	from=$(expr 1 + $RANDOM % "${#username[@]}")

	curl -X 'POST' \
	"http://localhost:3000/users/$from/myPhotos/" \
	-H 'accept: application/json' \
	-H 'Content-Type: multipart/form-data' \
	-H "Authorization: Bearer $from " \
	-F 'descriptionImg="Una foto spazialeeeee!" '\
	-F 'imageData=@tiny.png'
done



