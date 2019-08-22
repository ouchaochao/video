#!/bin/sh

function push(){
	echo "请输入commit的内容："
	read commit
	git add .
	git commit -m "$commit"
	git push origin master
	echo ""
	echo "Push done!"
}

function pull(){
	git pull origin master
}

echo "请输入序号："
echo "	1.push to github"
echo "	2.pull from github"
read input

if ((input == 1));then
	push
elif ((input == 2));then
	pull
fi
