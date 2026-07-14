#!/bin/bash

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

if [ -z "$1" ]; then
    echo -e "${RED}❌ Commit message do!${NC}"
    echo "Usage: ./push.sh 'Day 1: Hello World'"
    exit 1
fi

echo -e "${YELLOW}📁 Changed files:${NC}"
git status --short

echo ""
echo -e "${YELLOW}➕ Adding all files...${NC}"
git add .

echo ""
echo -e "${YELLOW}💾 Committing: $1${NC}"
git commit -m "$1"

echo ""
echo -e "${YELLOW}🚀 Pushing to GitHub...${NC}"
git push origin main

echo ""
echo -e "${GREEN}✅ GitHub pe push ho gaya!${NC}"
echo -e "${GREEN}🔗 https://github.com/vkssngh/go-learning-exercises${NC}"