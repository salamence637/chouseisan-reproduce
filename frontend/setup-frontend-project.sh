#!/bin/sh
npm i -g create-react-app
npx create-react-app react-app
cd react-app
npm i react-router-dom react-bootstrap bootstrap framer-motion
rm -r src
mv ../src .