#!/bin/bash
curl -s https://01.alem.school/assets/superhero/all.json | jq '.[] | select(."id"==170) | .name' | tr -d '"'
curl -s https://01.alem.school/assets/superhero/all.json | jq '.[] | select(."id"==170) | .powerstats.power'
curl -s https://01.alem.school/assets/superhero/all.json | jq '.[] | select(."id"==170) | .appearance.gender' | tr -d '"'
