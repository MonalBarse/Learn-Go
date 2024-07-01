const express = require("express");
const bodyParser = require("body-parser");

const app = express();

// Middleware functions
const validateContentType = (contentType) => {
  return (req, res, next) => {
    if (req.headers["content-type"] !== contentType) {
      return res.status(400).send({
        error: "Invalid Content-Type",
        message: `Expected ${contentType}`,
      });
    }
    next();
  };
};

const handleJSONParsingErrors = (err, req, res, next) => {
  if (err instanceof SyntaxError) {
    return res.status(400).send({
      error: "Invalid JSON",
      message: err.message,
    });
  }
  next();
};

// Middleware to parse application/x-www-form-urlencoded data globally
app.use(bodyParser.urlencoded({ extended: true }));

// Routes
app.get("/", (req, res) => {
  res.send("Welcome to the GET route!");
});

app.get("/get", (req, res) => {
  const name = req.query.name || "Unnamed visitor";
  res.send(`Hello, ${name}! (GET request)`);
});

app.post("/post", bodyParser.json(), (req, res) => {
  let dataReceived = req.body;
  res.status(200).send(dataReceived);
});

app.post(
  "/form",
  validateContentType("application/x-www-form-urlencoded"),
  bodyParser.urlencoded({ extended: true }),
  (req, res) => {
    let dataReceived = req.body;
    res.status(200).send(JSON.stringify(dataReceived));
  },
);

// Error handling middleware
app.use(handleJSONParsingErrors);

app.listen(3000, () => {
  console.log("Server listening on port 3000");
});
