const fs = require('fs')

try {
    const data = fs.readFileSync('./input', 'UTF-8')

    const lines = data.split(/\r?\n/)

    lines.forEach(line => {
        console.log(line)
    })
} catch (err) {
    console.error(err)
}
