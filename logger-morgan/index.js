const express = require('express');
const morgan = require('morgan');
const app = express();
const bparser = require('body-parser')
const path = require('path');
const fsr = require('file-stream-rotator');
const port = 3000;

app.use(bparser.urlencoded({ extended: false }))
app.use(express.static(path.join(__dirname, 'css')));
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');
let logsinfo = fsr.getStream({ filename: "test.log", frequency: "1h", verbose: true });
app.use(morgan('combined', { stream: logsinfo }))

app.get('/', function (req, res) {
    res.render('index.ejs');
})

app.post('/print', function (req, res) {
    console.log(req.body)
    res.render('redirect.ejs', { fname: req.body.fname, lname: req.body.lname })
})

app.listen(port, () => {
    console.log("Startet Server")
});
