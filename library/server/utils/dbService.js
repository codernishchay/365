
const DBservice = () => {
    /*
     * @description : create any mongoose document
     * @param  {obj} model : mongoose model
     * @param  {obj} data : {}
     * @return Promise
     */
    const createDocument = (model, data) => new Promise((resolve, reject) => {
        model.create(data, (err, result) => {
            if (err) reject(err);
            else resolve(result);
        });
    });


    /*
 * @description : update any existing mongoose document
 * @param  {obj} model : mongoose model
 * @param {ObjectId} id : mongoose document's _id
 * @param {obj} data : {}
 * @return Promise
 */
    const updateDocument = (model, id, data) => new Promise((resolve, reject) => {
        model.updateOne({ _id: id }, data, {
            runValidators: true,
            context: 'query',
        }, (err, result) => {
            if (err) reject(err);
            else resolve(result);
        });
    });


    /*
     * @description : get one single existing mongoose document
     * @param {obj} model : mongoose model 
     * @param {ObjectID} id : mongoose document's _id
     * @return Promise 
     */

    const getSingleDocumentById = (model, id) => new Promise((resolve, reject) => {
        model.findOne({ _id: id }, (err, result) => {
            if (err) reject(err);
            else resolve(result);
        })
    })

    const get

}