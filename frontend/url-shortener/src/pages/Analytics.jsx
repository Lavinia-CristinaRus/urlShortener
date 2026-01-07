import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import api from "../utils/api";
import { Bar } from "react-chartjs-2";

import {
  Chart as ChartJS,
  BarElement,
  CategoryScale,
  LinearScale,
  Tooltip,
  Legend,
} from "chart.js";

ChartJS.register(
  BarElement,
  CategoryScale,
  LinearScale,
  Tooltip,
  Legend
);

export default function Analytics() {
  const { id } = useParams();
  const [data, setData] = useState(null);
  const [error, setError] = useState("");


  const clicksPerVisitorData = {
    labels: data?.clicks_per_visitor?.map((v, i) => `Visitor ${i + 1}`),
    datasets: [
      {
        label: "Clicks",
        data: data?.clicks_per_visitor?.map(v => v.clicks),
      },
    ],
  };

  const visitorsPerCityData = {
    labels: data?.visitors_per_city?.map(v => v.city || "Unknown"),
    datasets: [
      {
        label: "Visitors",
        data: data?.visitors_per_city?.map(v => v.visitors),
      },
    ],
  };

  useEffect(() => {
    api.get(`/api/urls/analytics/${id}`)
      .then(res => setData(res.data))
      .catch(err => {
        setError(err.response?.data?.error || "Failed to load analytics");
      }
    );
  }, [id]);

  return (
    <div style={{ padding: "40px", alignItems: "center",justifyItems:"center" }}>
      <h2 style={{marginBottom:"80px"}}>URL Analytics</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {(!data || (data?.total_clicks == null) || (data?.unique_visitors==null) || (data?.geo==null))? <p>Loading analytics...</p>:
      (<div style={{display: "grid", gridTemplateColumns: "0.4fr 0.6fr", gap: "10px", alignItems: "center",justifyItems:"center", width:"100%"}}><div>
        <p><strong>Total clicks:</strong> {data.total_clicks}</p>
        <p><strong>Unique visitors:</strong> {data.unique_visitors}</p>
      </div>


      {data.geo.length === 0 ? (
        <div style={{ width: "80%", alignItems: "center",justifyItems:"center" }}><h3>Geographic distribution</h3>
        <p>No click data yet.</p></div>
      ) : (
        <div style={{ width: "80%", alignItems: "center",justifyItems:"center" }}><h3>Geographic distribution</h3>
        <table border="1" cellPadding="10" style={{ width: "80%" }}>
          <thead>
            <tr>
              <th>Country</th>
              <th>City</th>
              <th>Clicks</th>
            </tr>
          </thead>
          <tbody>
            {data.geo.map((g, i) => (
              <tr key={i}>
                <td>{g.Country}</td>
                <td>{g.City}</td>
                <td>{g.Clicks}</td>
              </tr>
            ))}
          </tbody>
        </table></div>
      )}
      </div>)}
      <div style={{alignContent:"center", justifyContent:"center", height:"200px", width:"auto", marginTop: "80px"}}>{data?.clicks_per_visitor?.length > 0 && (
        <div style={{ marginBottom: "40px" }}>
          <h3>Clicks per visitor</h3>
          <Bar data={clicksPerVisitorData} />
        </div>
      )}</div>
      <div style={{alignContent:"center", justifyContent:"center", width:"auto", height:"200px", marginTop: "80px"}}>
      {data?.visitors_per_city?.length > 0 && (
        <div>
          <h3>Visitors per city</h3>
          <Bar data={visitorsPerCityData} />
        </div>
      )}</div>
    </div>
  );
}