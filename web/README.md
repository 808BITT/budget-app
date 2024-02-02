## Setup 
```bash
$ git clone https://github.com/808BITT/template-solid.git project-name
$ cd project-name
```
Clones the repository and changes the directory to the project folder.<br>

## Install
```bash
$ npm install 
```
Installs all dependencies.<br>

## Dev
```bash
$ npm run dev
```
Starts the development server.<br>
Runs the app in the development mode.<br>
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.


## Build
```bash
$ npm run build
```
Builds the app for production to the `/dist/src` folder.<br>
This can be configured in the `vite.config.ts`,<br>


## Deploy
You can deploy the `dist` folder to any static host provider (Azure, AWS, Linode, etc.)
<br>
```bash
$ scp -r public user@server:/
```
Copies the `dist` folder to the server.<br>
