"use strict";

function roll() {
    return Math.floor(Math.random() * 2) +
        Math.floor(Math.random() * 2) +
        Math.floor(Math.random() * 2) +
        Math.floor(Math.random() * 2);
}

function getMove(from) {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var j = this.responseText
            document.getElementById("game").innerHTML = this.responseText;
            drawBoard(j)
        }
    };
    //var from = document.getElementById("movefrom").value;
    //var spaces = document.getElementById("spaces").value;
    xhttp.open("GET", "move?from=" + from + "&spaces=" + r);
    xhttp.send();
}

function getComputerMove() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var j = this.responseText
            document.getElementById("game").innerHTML = this.responseText;
            drawBoard(j)
     
        }
    };
    var spaces = document.getElementById("cspaces").value;
    xhttp.open("GET", "computer?spaces=" + spaces);
    xhttp.send();
}

var k = '{"WhiteStones":7,"BlackStones":6,"Board":{"Black":[0,0,0,0,0,0,1,0,0,0,0,0,0,0],"White":[0,0,0,0,0,0,0,0,0,0,0,0,0,0]},"Pot":0}';
var whitepos = [
    [449, 55],
    [548, 55],
    [646, 55],
    [745, 55],
    [745, 150],
    [646, 150],
    [548, 150],
    [449, 150],
    [351, 150],
    [252, 150],
    [154, 150],
    [55, 150],
    [55, 55],
    [155, 55]
];
var blackpos = [
    [449, 245],
    [548, 245],
    [646, 245],
    [745, 245],
    [745, 150],
    [646, 150],
    [548, 150],
    [449, 150],
    [351, 150],
    [252, 150],
    [154, 150],
    [55, 150],
    [55, 245],
    [155, 245]
];

const canvas = document.getElementById('canvas1')

var gamestate = 0 // three gamestates
                  // 0 I'm waiting for you to roll the dice
// 1 I'm waiting for you to move
// 2 I'm busy doing something else

var r = 0 //current roll

function drawBoard(j) {
    var c = canvas.getContext('2d');
    var board = JSON.parse(j);
    c.clearRect(0,0,800,300,false)
    drawStones("Black",board,c)
    drawStones("White",board,c)
}

function place(colour, x, c) {
    var radius = 40;
    c.beginPath();
    c.arc(x[0], x[1], radius, 0, Math.PI * 2);
    c.fillStyle = colour;
    c.strokeStyle = "black";
    c.fill();
    c.stroke();
}

function drawStones(colour,board, c) {
    var colourboard = board["Board"][colour];
    for (var i = 0; i < colourboard.length; i++) {
        if (colourboard[i] === 1) {
            var pos = colour === "Black" ? blackpos[i] : whitepos[i]
            place(colour, pos, c, colour);
        }
    }
}

function squareToBlackPos(n) {
   return  [-1,-1,-1,-1,-1,-1,-1,-1,
            11,10, 9, 8, 7, 6, 5, 4,
            12,13,-1,-1, 0, 1, 2, 3,][n]
}

function do_move(e) {
    var rect = canvas.getBoundingClientRect();

    var square = Math.trunc((e.clientX-rect.left) / 98.57) + 8*Math.trunc((e.clientY-rect.top)/95)
    console.log("Canvas click " + e.clientX +","+ e.clientY+" : "+square )
    if (square >13) {square=-1}
    getMove(square)
    gamestate = 0
}

function do_roll() {
    r  = roll();
    var c =canvas.getContext('2d');
    c.clearRect(800,0,200,300);
    c.font = "96px Ariel";
    c.fillstyle = "black";
    c.textAlign = "center";
    c.fillText(r, 900, 200)

    gamestate = 1
}

canvas.addEventListener('click', (e) => {
    console.log(gamestate)
    if (gamestate == 2) {return} // I am busy ignore this click

    var rect = canvas.getBoundingClientRect();
    //I'm waiting for a roll
    if (gamestate == 0 && e.clientX-rect.left > 800) { 
        gamestate = 2; 
        do_roll()
        return   
    }
    //I'm waiting for a move
    if (e.clientX-rect.left > 800){return}
    gamestate = 2;
    do_move(e)
});
