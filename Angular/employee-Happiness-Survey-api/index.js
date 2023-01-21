// Import dependencies
const express = require("express");
const database = require("./initdb.utils");

// Create app instance
const app = express();

// Define JSON as return type
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Configure headers
app.use((req, res, next) => {
  res.setHeader("Access-Control-Allow-Origin", "*");
  res.setHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE");
  res.setHeader("Access-Control-Allow-Headers", "Content-Type");
  res.setHeader("Access-Control-Allow-Credentials", true);
  next();
});

// DEFINE ROUTES
app.use("/hello", (req, res) => {
  res.status(200).json({ message: "Hello world !" });
});
app.use("/surveys", require("./survey.routes"));

// This route will handle all the requests that are
// not handled by any other route handler. In
// this hanlder we will redirect the user to
// an error page with NOT FOUND message and status
// code as 404 (HTTP status code for NOT found)
app.all("*", (req, res) => {
  res.status(404).json({ error: "End point not found" });
});

// Handle database error
database.on("error", (error) => {
  console.log("Connection error: --------------------------");
  console.log(error);
  console.log("--------------------------------------------");
});

// Start app after connecting to Database
database.once("connected", () => {
  console.log("Database Connected");

  const PORT = 3001;
  app.listen(PORT, () => console.log("Server ready at 3001"));
});
