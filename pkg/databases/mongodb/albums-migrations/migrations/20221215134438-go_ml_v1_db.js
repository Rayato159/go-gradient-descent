import { Double } from "mongodb";

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
                        required: ["weights", "error", "timestamp"],
                        properties: {
                            weights: {
                                bsonType: "array",
                                items: {
                                    bsonType: "double",
                                },
                                description:
                                    "'weights' must be a [double] and is required",
                            },
                            error: {
                                bsonType: "double",
                                description:
                                    "'error' must be a double and is required",
                            },
                            timestamp: {
                                bsonType: "date",
                            },
                        },
                    },
                },
            });
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
