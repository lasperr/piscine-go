#!/bin/bash
id=$(echo ${HERO_ID})
json=$(curl -s https://01.alem.school/assets/superhero/all.json) 
echo $json | jq '.[] | select(."id" == '$id') | .connections.relatives' | tr -d '"'
