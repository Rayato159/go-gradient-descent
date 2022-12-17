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
<p>***All example has been added into the postman collection and environment, you can check it</p>
<ul>
    <li><strong>Get data</strong></li>
    <li><strong>Insert a new data</strong></li>
    <li><strong>Train data</strong></li>
    <li><strong>Predict a result</strong></li>
    <li><strong>Clear data</strong></li>
    <li><strong>Clear weights</strong></li>
</ul>

<h2>Solution</h2>
<img src="pkg\plots\train.png">

<h2>Test</h2>
<img src="pkg\plots\test.png">

<h2>Machine Learning Details</h2>

```bash
Model:          Linear Regression
Optimizer:      Gradient Descent
Loss Function:  Error Sum of Squares
Max Iterations: 1000
```

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

<h2>⚙️ Start Project</h2>

Clone the project
```bash
git clone https://github.com/Rayato159/go-gradient-descent.git
cd ./go-gradient-descent
```

Start Go service
```bash
air init
```

Copy and paste the code below and repalce for some section in air.toml
```bash
[build]
  cmd = "go build -o ./tmp/main.exe ./app/main.go"
  exclude_dir = ["assets", "tmp", "vendor", "testdata", "pkg\\databases\\mongodb\\albums-migrations"]

[misc]
  clean_on_exit = true
```

.env file
```bash
#Stage
STAGE=dev

#App
FIBER_HOST=127.0.0.1
FIBER_PORT=3000
APP_VERSION=1.0.0
FIBER_REQUEST_TIMEOUT=120

#Database
#MongoDb
MONGODB_HOST=127.0.0.1
MONGODB_PORT=27017
MONGODB_DATABASE=go_ml_v1_db
MONGODB_USERNAME=root
MONGODB_PASSWORD=123456
```

Start the project!, The services are divided in two parts

First start the Go service
```
air -c air.toml
```

Second start the python service
```
cd ./pkg/plots
python main.py
```

If you get an error go find and install a module in this lists

```bash
matplotlib
numpy
flask
pymongo
dotenv
```

Have fun! 😄