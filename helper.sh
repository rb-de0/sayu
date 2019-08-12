#!/bin/bash

for file in ./*.jpg; do
  filename=$(basename "$file")
  sayu $filename 50
done