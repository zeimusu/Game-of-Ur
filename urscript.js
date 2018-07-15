"use strict";

function roll() {
    return Math.floor(Math.random() * 2) +
        Math.floor(Math.random() * 2) +
        Math.floor(Math.random() * 2) +
        Math.floor(Math.random() * 2);
}

function getMove() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var j = this.responseText
            document.getElementById("game").innerHTML = this.responseText;
            drawBoard(j)
        }
    };
    var from = document.getElementById("movefrom").value;
    var spaces = document.getElementById("spaces").value;
    xhttp.open("GET", "move?from=" + from + "&spaces=" + spaces);
    xhttp.send();
}

function getComputerMove() {
    alert("computer move")
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
    [450, 50],
    [550, 50],
    [650, 50],
    [750, 50],
    [750, 150],
    [650, 150],
    [550, 150],
    [450, 150],
    [350, 150],
    [250, 150],
    [150, 150],
    [50, 150],
    [50, 50],
    [150, 50]
];
var blackpos = [
    [450, 250],
    [550, 250],
    [650, 250],
    [750, 250],
    [750, 150],
    [650, 150],
    [550, 150],
    [450, 150],
    [350, 150],
    [250, 150],
    [150, 150],
    [50, 150],
    [50, 250],
    [150, 250]
];
function drawBoard(j) {
    var board = JSON.parse(j);
    var c = document.getElementById('canvas1').getContext('2d');
    c.clearRect(0,0,800,300,false)
    drawStones("Black",board,c)
    drawStones("White",board,c)
}

function place(colour, x, c) {
    var radius = 40;
    c.beginPath();
    c.arc(x[0], x[1], radius, 0, Math.PI * 2);
    c.fillStyle = colour;
    c.fill();
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

