import { useState } from "react";
import { Cell } from "./Cell";

export const Board = () => {
  const [board, setBoard] = useState([
    [" ", " ", " "],
    [" ", " ", " "],
    [" ", " ", " "],
  ]);

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

    const updated = await response.json();
    setBoard(updated);
  };

  return (
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
  );
};
