import { useState, useEffect } from "react";
import axios from "axios";

export const useGetUrl = (url) => {
  const [data, setData] = useState(null);
  const [error, setError] = useState(false);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (!url) return;
    axios({
      method: "get",
      url: `${url}`,
      headers: {},
    })
      .then(function (response) {
        setData(response.data);
        setLoading(false);
      })
      .catch(function (error) {
        setError(error);
      });
  }, [url]);
  return { data, error, loading };
};
