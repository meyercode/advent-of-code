const fs = require('fs')

try {
    const data = fs.readFileSync('./input', 'UTF-8')

    const lines = data.split(/\r?\n/)

    let score = 0
    const X = 1
    const Y = 2
    const Z = 3
    const WIN = 6
    const DRAW = 3

    lines.forEach(line => {
        const moves = line.split(" ")
        const oppMove = moves[0]
        const myMove = moves[1]

        /*
            * A < Y     A == X
            * B < Z     B == Y
            * C < X     C == Z
        */

        switch (myMove) {
            case 'X':
                score += X
                if (oppMove == 'C')
                    score += WIN
                else if (oppMove == 'A')
                    score += DRAW
                break
            case 'Y':
                score += Y
                if (oppMove == 'A')
                    score += WIN
                else if (oppMove == 'B')
                    score += DRAW
                break
            case 'Z':
                score += Z
                if (oppMove == 'B')
                    score += WIN
                else if (oppMove == 'C')
                    score += DRAW
                break
        }

    })

    console.log(score)

} catch (err) {
    console.error(err)
}
