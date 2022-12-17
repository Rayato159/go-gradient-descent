from rand_sample import Database
from flask import Flask, send_file
from matplotlib import pyplot as plt
import numpy as np
from dotenv import load_dotenv

# Load database
load_dotenv("../../.env.dev")
db = Database().get_db()

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

app = Flask(__name__)

@app.route('/get-plot', methods=["GET"])
def stream_file():
    get_plot()
    return send_file("./test_1.png", mimetype="image/png") 

if __name__ == "__main__":
    app.run()