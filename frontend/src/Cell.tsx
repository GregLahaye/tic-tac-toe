export const Cell = (props: any) => {
  return (
    <button className="cell" onClick={props.onClick}>
      {props.value}
    </button>
  );
};
