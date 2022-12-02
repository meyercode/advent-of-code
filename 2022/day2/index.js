const fs = require('fs')

const ROCK = 1
const PAPER = 2
const SCISSORS = 3
const WIN = 6
const DRAW = 3

const strategyMap = { 
    'A': {
        loss: SCISSORS,
        draw: ROCK,
        win: PAPER,
    },
    'B': {
        loss: ROCK,
        draw: PAPER,
        win: SCISSORS,
    },
    'C': {
        loss: PAPER,
        draw: SCISSORS,
        win: ROCK,
    },
}

try {
    const data = fs.readFileSync('./input', 'UTF-8')

    const lines = data.split(/\r?\n/)

    let score = 0

    lines.forEach(line => {
        const moves = line.split(" ")
        const oppMove = moves[0]
        const myMove = moves[1]

        /*
            * A == ROCK     
            * B == PAPER    
            * C == SCISSORS 
            *
            * X == loss
            * Y == draw
            * Z == win
        */

        switch (myMove) {
            case 'X':
                score += strategyMap[oppMove].loss
                break
            case 'Y':
                score += strategyMap[oppMove].draw
                score += DRAW
                break
            case 'Z':
                score += strategyMap[oppMove].win
                score += WIN
                break
        }

    })

    console.log(score)

} catch (err) {
    console.error(err)
}
