
const Schema = mongoose.Schema;
const schema = new Schema({
    name: {
        type: String,
        required: true,
        ref: 'bookname'
    }
    ,
    description: {
        type: String,
        ref: 'description'
    }
    ,
    author: {
        type: Schema.Types.ObjectID,
        required: true,
    }
},
    {
        timestamp: {
            required: true,
        }
    });

const Book = mongoose.model('Book', schema);
module.exports = Book;
