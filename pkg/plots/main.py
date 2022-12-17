import os
import random
import time

from flask import Flask, send_file
from pymongo import MongoClient
from matplotlib import pyplot as plt
import numpy as np

from dotenv import load_dotenv

class Database:
    def __init__(self):
        self.host = os.getenv("MONGODB_HOST")
        self.port = os.getenv("MONGODB_PORT")
        self.database = os.getenv("MONGODB_DATABASE")
        self.username = os.getenv("MONGODB_USERNAME")
        self.password = os.getenv("MONGODB_PASSWORD")

    def get_db(self) -> MongoClient:
        # Provide the mongodb atlas url to connect python to mongodb using pymongo
        CONNECTION_STRING = f"mongodb://{self.username}:{self.password}@{self.host}:{self.port}"

        # Create a connection using MongoClient. You can import MongoClient or use pymongo.MongoClient
        db = MongoClient(CONNECTION_STRING)

        # Create the database for our example (we will use the same database throughout the tutorial
        return db[f"{self.database}"]

# Load database
load_dotenv("../../.env.dev")
db = Database().get_db()

def rand_data():
    from sklearn.model_selection import train_test_split
    random.seed(time.time_ns())
    
    x = np.linspace(start=-10, stop=10, num=100)

    slope = random.uniform(-1, 1)
    y_intercept = random.uniform(-1, 1)
    error = [random.uniform(-1, 1) for x in range(100)]

    y = (slope*x + y_intercept) + error
    y_true = (slope*x + y_intercept)

    x_train, x_test, y_train, y_test = train_test_split(x, y, test_size=0.2, random_state=random.randint(0, 100))

    train_doc = []
    test_doc = []
    for i, j in zip(x_train, y_train):
        train_doc.append({
            "x": i,
            "y": j
        })
    for i, j in zip(x_test, y_test):
        test_doc.append({
            "x": i,
            "y": j
        })

    train_data_coll = db["train_data"]
    test_data_coll = db["test_data"]

    train_data_coll.insert_many(train_doc)
    test_data_coll.insert_many(test_doc)

    print("success, all data has been inserted")

def get_plot():
    latest_record = db["records"].find_one(sort=[("$natural", -1)])
    test_data = db["test_data"].find()

    params, err = latest_record["weights"], latest_record["error_test"]

    x_test, y_test, y_true = [], [], []
    for i in test_data:
        x_test.append(float(i["x"]))
        y_test.append(float(i["y"]))

    slope, y_intercept = float(params[0]), float(params[1])
    for i in x_test:
        y_true.append(params[0]*i + params[1])

    op = "+" if y_intercept > 0 else ""
    op = "+" if y_intercept > 0 else ""
    plt.figure(figsize=(5,4))
    plt.title(f"f(x) = {slope:.6f}x {op} {y_intercept:.6f}")
    plt.xlabel("x")
    plt.ylabel("y")
    plt.scatter(x_test, y_test, color="#FF597B")
    plt.plot(x_test, y_true, color="#2B3A55")
    plt.savefig('test_1.png')

# Service API
app = Flask(__name__)

@app.route('/insert-rand-data', methods=["POST"])
def insert_rand_data():
    # rand_data()
    return {
        "message": "success, data has been inserted"
    }, 201

@app.route('/get-plot', methods=["GET"])
def stream_file():
    get_plot()
    return send_file("./test_1.png", mimetype="image/png"), 200

if __name__ == "__main__":
    app.run()