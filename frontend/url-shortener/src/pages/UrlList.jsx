import { useEffect, useState } from "react";
import api from "../utils/api";
import { useNavigate } from "react-router-dom";

export default function MyUrls() {
  const [urls, setUrls] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    api.get("/api/urls")
      .then(res => setUrls(res.data))
      .catch(err => console.error(err));
  }, []);

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <h2 style={styles.h2}>My Short Links</h2>
        <button
          style={styles.createButton}
          onClick={() => navigate("/generateUrl")}
        >
          + Create new short link
        </button>
      </div>

      {urls.length === 0 && (
        <p style={styles.empty}>You don’t have any short links yet.</p>
      )}

      <div style={styles.list}>
        {urls.map(url => (
          <div key={url.Idurl} style={styles.row}>
            <a
              href={url.Long_url}
              target="_blank"
              rel="noopener noreferrer"
              style={styles.original}
              title={url.Long_url}
            >
              {url.Long_url}
            </a>

            <a
              href={`http://localhost:4000/${url.Short_url}`}
              target="_blank"
              rel="noopener noreferrer"
              style={styles.short}
            >
              {`http://localhost:4000/${url.Short_url}`}
            </a>
            {url.Expires_at?<p style={styles.short}>{`Expires at: ${url.Expires_at}`}</p>:<p style={styles.short}>{`No expiration date set`}</p>}

            <button
              style={styles.edit}
              onClick={() => navigate(`/edit/${url.id}`)}
            >
              Customize
            </button>
            <button
              style={styles.edit}
              onClick={() => navigate(`/analytics/${url.id}`)}
            >
              Analytics
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}

const styles = {
  container: {
    width: "90%",
    margin: "40px auto",
    marginTop: "40px",
    padding: "20px 20px",
    fontFamily: "sans-serif",
    marginLeft: "40px",
    marginRight: "40px",
  },
  header: {
    width: "90%",
    display: "flex",
    alignItems: "center",
    marginBottom: "20px",
    marginLeft: "40px",
    marginRight: "40px",
  },
  h2: {
    marginRight: "100px",
  },
  createButton: {
    padding: "8px 14px",
    borderRadius: "6px",
    border: "none",
    backgroundColor: "#4f46e5",
    color: "#fff",
    cursor: "pointer",
    justifyContent: 'right',
  },
  list: {
    display: "flex",
    flexDirection: "column",
    gap: "10px",
  },
  row: {
    display: "grid",
    gridTemplateColumns: "0.5fr 0.22fr 0.22fr auto auto",
    gap: "85px",
    alignItems: "center",
    padding: "10px 14px",
    background: "#f9fafb",
    borderRadius: "8px",
    justifyContent: 'center',
  },
  original: {
    color: "#111827",
    textDecoration: "none",
  },
  short: {
    fontWeight: "bold",
    color: "#2563eb",
    textDecoration: "none",
  },
  edit: {
    padding: "6px 10px",
    marginLeft:"10px",
    borderRadius: "6px",
    border: "1px solid #d1d5db",
    background: "#fff",
    cursor: "pointer",
    backgroundColor: "#4f46e5",
    color: "#fff",
    fontWeight: "bold",
  },
  empty: {
    color: "#6b7280",
  },
};