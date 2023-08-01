#!/bin/bash

find * -type f -name "*.sh" -exec basename \{} .sh \;  | sort -nr | sed -e 's/\.sh$//'
