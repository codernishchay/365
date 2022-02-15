const dotEnv = require("dotenv");
dotEnv.config()

const mongoose = require('mongoose');
const uri = process.env.NODE_ENV === 'test' ? process.env.DATABASE_URL : process.env.DATABASE_URL;
mongoose.connect(uri, {
    useNewUrlParser: true,
    useUnifiedTopology: true,
    useFindAndModify: false,
    useCreateIndex: true
});
let db = mongoose.connection;

db.once('open', () => {
    console.log('Connection Successful');
});

db.on('error', () => {
    console.log('Error in mongodb connection');
});

module.exports = mongoose;