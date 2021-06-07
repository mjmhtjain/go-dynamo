const fs = require('fs')
const ObjectsToCsv = require('objects-to-csv');
const dataSetLength = 10000 // length of dataset
const tableName = "UserDetail"

var generateRandomDataset = function () {
    let dataArr = []

    for (let i = 1; i <= dataSetLength; i++) {
        let item = {}
        item["id"] = stringTypeObject("" + i)
        item["name"] = stringTypeObject(randomString(5))
        item["email"] = stringTypeObject(randomString(5))

        dataArr.push(item)
    }

    exportCSV('./test.csv', dataArr)
}

var stringTypeObject = function (val) {
    return val;
}

var randomString = function (len) {
    var result = [];
    var characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    var charactersLength = characters.length;
    for (var i = 0; i < len; i++) {
        result.push(characters.charAt(Math.floor(Math.random() *
            charactersLength)));
    }
    return result.join('');
}

var exportData = function (filePath, data) {

    fs.writeFile(filePath, JSON.stringify(data, null, 4), function (err) {
        if (err == null) {
            console.log("Wrote successfully !!")
        } else {
            console.log("err: " + err)
        }
    })
}

var exportCSV = async function (filePath, data) {
    const csv = new ObjectsToCsv(data);

    // Save to file:
    await csv.toDisk(filePath); // './test.csv'

    // Return the CSV file as string:
    console.log(await csv.toString());
}

generateRandomDataset()