import { useState } from "react";
import { Cell } from "./Cell";

interface Coordinate {
  x: number;
  y: number;
}

export const Board = () => {
  const [board, setBoard] = useState([
    [" ", " ", " "],
    [" ", " ", " "],
    [" ", " ", " "],
  ]);

  const [recommendation, setRecommendation] = useState<Coordinate>();
  const [message, setMessage] = useState("");

  const handleClick = async (row: number, col: number) => {
    board[row][col] = "X";
    setBoard([...board]);

    const response = await fetch("http://localhost:8080", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify(board),
    });

    const data = await response.json();
    setBoard(data.board);

    setRecommendation(data.recommendation);
    setMessage(data.state);
  };

  return (
    <div>
      <div>
        {board.map((cells, row) => (
          <div key={`${row}`} className="board-row">
            {cells.map((cell, col) => (
              <Cell
                key={`${row},${col}`}
                value={cell}
                onClick={() => handleClick(row, col)}
              ></Cell>
            ))}
          </div>
        ))}
      </div>

      {recommendation ? (
        <p>
          Best Move: ({recommendation.x}, {recommendation.y})
        </p>
      ) : null}

      <p>{message}</p>
    </div>
  );
};
