import { splitCamelCaseAndUppercase, isObject } from "../functions";
import { useGetUrl } from "../hooks/useGetUrl";

const DetailTable = ({ part, setPart }) => {
  // Get the data from the url.
  const { data, error, loading } = useGetUrl(`api/part?id=${part - 1}`);

  // Takes an object and returns a table of its properties and values.
  const JsonTableComponent = (json) => {
    if (isObject(json)) {
      let table = [];
      for (let key in json) {
        if (json.hasOwnProperty(key)) {
          table.push(
            <tr key={key}>
              <td style={{ backgroundColor: "#EAE7E6" }}>
                <b>{splitCamelCaseAndUppercase(key)}</b>
              </td>
              <td>{JsonTableComponent(json[key])}</td>
            </tr>
          );
        }
      }
      return (
        <table className="table" border="1">
          <tbody>{table}</tbody>
        </table>
      );
    }
    return json;
  };

  if (error) console.warn("Error: ", error);

  return (
    <>
      {loading ? (
        <div>Loading ...</div>
      ) : (
        <div>
          <h2>Service {part}</h2>
          <button onClick={() => setPart("")} className="button-1">
            Go back to index
          </button>
          <div className="tableContainer">
            <JsonTableComponent json={data} />
          </div>
          <button onClick={() => setPart("")} className="button-1">
            Go back to index
          </button>
        </div>
      )}
    </>
  );
};

export default DetailTable;
