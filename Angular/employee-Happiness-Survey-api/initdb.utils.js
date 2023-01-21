const mongoose = require("mongoose");

const MONGO_HOST = "localhost";
const MONGO_PORT = 27017;
const MONGO_DB_NAME = "surveydb";
const MONGO_URL = `mongodb://${MONGO_HOST}:${MONGO_PORT}/${MONGO_DB_NAME}`;
const connectOptions = {
  useNewUrlParser: true,
};

mongoose.connect(MONGO_URL, connectOptions);
const database = mongoose.connection;

module.exports = database;
