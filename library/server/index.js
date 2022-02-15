const express = require('express');
const app = express();
const Server = require("http").createServer(app);
var cors = require('cors');
const router = express.Router();

app.use(cors());
app.use(express.json());

console.log("started server ")


app.use(router);
Server.listen(3000);
