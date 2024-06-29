const express = require("express");
const app = express();

const bodyParser = require("body-parser");

app.use(bodyParser.urlencoded({ extended: true })); // for parsing application/x-www-form-urlencoded

app.get("/", (req, res) => {
  res.send("Welcome to the GET route!");
});

app.get("/get", (req, res) => {
  const name = req.query.name || "Unnamed visitor";
  res.send(`Hello, ${name}! (GET request)`);
});

app.post("/", (req, res) => {
  const name = req.body.name || "Unnamed visitor";
  res.send(`Hello, ${name}! (POST request)`);
});
// Accept any JSON data and return it back to the client
app.post("/post", (req, res) => {
  const recievedData = req.body;
  // We recieve raw JSON data and send it back as a response

  res.json(recievedData);
});
app.listen(3000, () => {
  console.log("Server listening on port 3000");
});
