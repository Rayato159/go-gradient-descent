<h1>🤖 Machine Learning as a Services</h1>
<p>This project is to code a linear regressgion model as a service (Just a noob model)</p>

<h2>🧑🏼‍💻 Established by</h2>
<ul>
    <li><strong>Ruangyot Nanchiang</strong></li>
</ul>

<h2>🚀 Version</h2>
<ul>
    <li><strong>1.0.0</strong></li>
</ul>

<h2>API List</h2>
<ul>
    <li><strong>Get Data</strong></li>
    <li><strong>Insert a new data</strong></li>
    <li><strong>Train Data</strong></li>
    <li><strong>Predict a result</strong></li>
    <li><strong>Clear data</strong></li>
</ul>

<h2>Solution</h2>
<img src="pkg\plots\figure_1.png">

<h2>🐳 MongoDb Docker Run Container</h2>

pull mongo image
```bash
docker pull mongo
```
start container
```bash
docker run -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=123456 -p 27017:27017 -v data:/data --name go-ml-v1-db -d mongo:latest
```
configuration (before into the container)
```bash
docker exec -it go-ml-v1-db bash
```
configuration (on the container)
```bash
mongosh mongodb://root:123456@127.0.0.1:27017
```
```bash
use go-ml-v1-db
```
```bash
show dbs
```

output
```bash
admin          100.00 KiB
config         108.00 KiB
local          72.00 KiB
```

<h2>🍃 MongoDb Migration</h2>

```bash
cd ./package/database/mongodb/albums-migrations
```

install dependencies
```bash
npm install
```

up or down
```bash
migrate-mongo [up|down]
```