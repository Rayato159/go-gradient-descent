import { Double } from "mongodb"

export const up = async (db, client) => {
    const session = client.startSession();
    try {
        await session.withTransaction(async () => {
            await db.createCollection("train_data", {
                validator: {
                    $jsonSchema: {
                        bsonType: "object",
                        title: "Train Data Object Validation",
                        required: ["x", "y"],
                        properties: {
                            x: {
                                bsonType: "double",
                                description:
                                    "'x' must be a double and is required",
                            },
                            y: {
                                bsonType: "double",
                                description:
                                    "'y' must be a double and is required",
                            },
                        },
                    },
                },
            });

            await db.createCollection("test_data", {
                validator: {
                    $jsonSchema: {
                        bsonType: "object",
                        title: "Test Data Object Validation",
                        required: ["x", "y"],
                        properties: {
                            x: {
                                bsonType: "double",
                                description:
                                    "'x' must be a double and is required",
                            },
                            y: {
                                bsonType: "double",
                                description:
                                    "'y' must be a double and is required",
                            },
                        },
                    },
                },
            });

            await db.createCollection("records", {
                validator: {
                    $jsonSchema: {
                        bsonType: "object",
                        title: "Weight Object Validation",
                        required: ["weights", "error"],
                        properties: {
                            weights: {
                                bsonType: ["double"],
                                description:
                                    "'weights' must be a [double] and is required",
                            },
                            error: {
                                bsonType: "double",
                                description:
                                    "'error' must be a double and is required",
                            },
                        },
                    },
                },
            });

            await db.collection("train_data").insertMany([
                {
                    x: Double(1.0),
                    y: Double(2.0),
                },
                {
                    x: Double(2.0),
                    y: Double(3.0),
                },
                {
                    x: Double(3.0),
                    y: Double(4.0),
                },
                {
                    x: Double(4.0),
                    y: Double(6.0),
                },
                {
                    x: Double(5.0),
                    y: Double(7.0),
                },
                {
                    x: Double(6.0),
                    y: Double(9.0),
                },
                {
                    x: Double(7.0),
                    y: Double(11.0),
                },
                {
                    x: Double(8.0),
                    y: Double(13.0),
                },
            ]);

            await db.collection("test_data").insertMany([
                {
                    x: Double(9.0),
                    y: Double(15.0),
                },
                {
                    x: Double(10.0),
                    y: Double(17.0),
                },
            ]);
        });
    } finally {
        await session.endSession();
    }
};

export const down = async (db, client) => {
    const session = client.startSession();
    try {
        await session.withTransaction(async () => {
            await db.collection("data").drop();
            await db.collection("records").drop();
        });
    } finally {
        await session.endSession();
    }
};
