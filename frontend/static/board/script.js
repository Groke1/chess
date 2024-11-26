$(document).ready(function() {
    const board = $('#chessBoard');

    const initialPosition = {
        'a8': 'R', 'b8': 'N', 'c8': 'B', 'd8': 'Q', 'e8': 'K', 'f8': 'B', 'g8': 'N', 'h8': 'R',
        'a7': 'P', 'b7': 'P', 'c7': 'P', 'd7': 'P', 'e7': 'P', 'f7': 'P', 'g7': 'P', 'h7': 'P',
        'a2': 'p', 'b2': 'p', 'c2': 'p', 'd2': 'p', 'e2': 'p', 'f2': 'p', 'g2': 'p', 'h2': 'p',
        'a1': 'r', 'b1': 'n', 'c1': 'b', 'd1': 'q', 'e1': 'k', 'f1': 'b', 'g1': 'n', 'h1': 'r'
    };


    function chooseImage(piece) {
        switch (piece) {
            case "R": return "wR.png"
            case "N": return "wN.png"
            case "B": return "wB.png"
            case "K": return "wK.png"
            case "Q": return "wQ.png"
            case "P": return "wP.png"
            case "r": return "bR.png"
            case "n": return "bN.png"
            case "b": return "bB.png"
            case "k": return "bK.png"
            case "q": return "bQ.png"
            case "p": return "bP.png"
        }
    }

    function createBoard() {
        const files = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'];
        const ranks = ['8', '7', '6', '5', '4', '3', '2', '1'];

        ranks.forEach((rank, rankIndex) => {
            files.forEach((file, fileIndex) => {
                const square = $('<div></div>')
                    .addClass('square')
                    .addClass((rankIndex + fileIndex) % 2 === 0 ? 'white' : 'black')
                    .attr('data-position', file + rank);

                const position = file + rank;
                if (initialPosition[position]) {
                    const piece = $('<div></div>')
                        .addClass('piece')
                        .append(
                            $('<img>')
                                .attr('src', `img/chesspieces/${chooseImage(initialPosition[position])}`)
                                .addClass('piece-img')
                        );
                    square.append(piece);
                }

                board.append(square);
            });
        });
    }


    board.on('click', '.square', function() {
        $('.square').removeClass('selected');
        $('.square').removeClass('selected');
        $(this).addClass('selected');
    });

    createBoard();
});