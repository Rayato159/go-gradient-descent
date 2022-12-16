import numpy as np
import random
import time
from pymongo import MongoClient
from dotenv import load_dotenv
import os

class Configs:
    def __init__(self):
        self.database = Database()

class Database:
    def __init__(self):
        self.host = os.getenv("MONGODB_HOST")
        self.port = os.getenv("MONGODB_PORT")
        self.database = os.getenv("MONGODB_DATABASE")
        self.username = os.getenv("MONGODB_USERNAME")
        self.password = os.getenv("MONGODB_PASSWORD")

def main(cfg: Configs):
    from sklearn.model_selection import train_test_split
    
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

    db = mongo_conn(cfg)
    train_data_coll = db["train_data"]
    test_data_coll = db["test_data"]

    train_data_coll.insert_many(train_doc)
    test_data_coll.insert_many(test_doc)

    print("success, all data has been inserted")

    from matplotlib import pyplot as plt
    
    op = "+" if y_intercept > 0 else ""
    op = "+" if y_intercept > 0 else ""
    plt.figure(figsize=(5,4))
    plt.title(f"f(x) = {slope:.6f}x {op} {y_intercept:.6f}")
    plt.xlabel("x")
    plt.ylabel("y")
    plt.scatter(x, y, color="#FF597B")
    plt.plot(x, y_true, color="#2B3A55")
    plt.show()

def mongo_conn(cfg: Configs) -> MongoClient:
    # Provide the mongodb atlas url to connect python to mongodb using pymongo
    CONNECTION_STRING = f"mongodb://{cfg.database.username}:{cfg.database.password}@{cfg.database.host}"

    # Create a connection using MongoClient. You can import MongoClient or use pymongo.MongoClient
    db = MongoClient(CONNECTION_STRING)

    # Create the database for our example (we will use the same database throughout the tutorial
    return db[f"{cfg.database.database}"]

if __name__ == "__main__":
    random.seed(time.time_ns())
    load_dotenv("../../.env.dev")

    # Load configs
    cfg = Configs()

    main(cfg)