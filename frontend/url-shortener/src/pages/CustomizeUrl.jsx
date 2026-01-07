import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import api from "../utils/api";

export default function CustomizeUrl() {
  const { id } = useParams();
  const navigate = useNavigate();

  const [shortUrl, setShortUrl] = useState("");
  const [expiresAt, setExpiresAt] = useState("");
  const [error, setError] = useState(null);
  const [message, setMessage] = useState(null);

  useEffect(() => {
    api.get(`/api/urls/${id}`)
      .then(res => {
        setShortUrl(res.data.Short_url);
        setExpiresAt(res.data.Expires_at.split("T")[0] || "");
        localStorage.setItem("lastSavedUrl", res.data.Short_url);
        localStorage.setItem("lastSavedDate", res.data.Expires_at);
      })
      .catch(() => setError("Failed to load URL"));
  }, [id]);

  const cancelEditShortUrl = async () => {
    setMessage("");
    setError("");
    const urlData = localStorage.getItem("lastSavedUrl");
    setShortUrl(urlData);
  }

  const cancelEditDate = async () => {
    setMessage("");
    setError("");
    const urlData = localStorage.getItem("lastSavedDate");
    setExpiresAt(urlData.split("T")[0] || "");
  }

  const saveShortUrl = async () => {
    try {
      setMessage("");
      setError("");
      const customurl = shortUrl;
      const idurl = Number(id);
      await api.post(`/api/customizeUrl`, {
        idurl,
        customurl,
      });
      localStorage.setItem("lastSavedUrl", customurl);
      setMessage("URL updated");
    } catch (err) {
      setError(err.response?.data?.error || "Update failed");
    }
  }

  const saveDate = async () => {
    try {
      setMessage("");
      setError("");
      const expires_at = new Date(expiresAt).toISOString();
      const idurl = Number(id);
      await api.post(`/api/setUrlExpirationDate`, {
        idurl,
        expires_at,
      });
      localStorage.setItem("lastSavedDate", expiresAt);
      setMessage("Expiry date updated");
    } catch (err) {
      setError(err.response?.data?.error || "Update failed");
    }
  }

  const back = async () => {
    navigate("/urls");
  }

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <div onClick={back} style={styles.signOut}>
          <p>Back</p>
        </div>
        <h2 style={styles.h2}>Customize URL</h2>
      </div>

      {error && <p style={{color:"red"}}>{error}</p>}
      {message && <p style={{color:"green"}}>{message}</p>}
    <div style={styles.row}>
      <input
        placeholder="Custom short code"
        value={shortUrl}
        onChange={e => setShortUrl(e.target.value)}
        style={{marginRight:"80px"}}
      />
      <button
        style={styles.createButton}
        onClick={saveShortUrl}
      >
        Save
      </button>
      <button
        style={styles.cancelButton}
        onClick={cancelEditShortUrl}
      >
        Cancel
      </button>
      </div>

      <div style={styles.row}>
      <input
        type="date"
        value={expiresAt || ""}
        onChange={e => setExpiresAt(e.target.value)}
        style={{marginRight:"80px"}}
      />

      <button
        style={styles.createButton}
        onClick={saveDate}
      >
        Save
      </button>
      <button
        style={styles.cancelButton}
        onClick={cancelEditDate}
      >
        Cancel
      </button>
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
    width: "100%",
    display: "flex",
    alignItems: "center",
    marginBottom: "80px",
  },
  h2: {
    width: "30%",
    marginRight: "215px",
  },
  cancelButton: {
    padding: "8px 14px",
    borderRadius: "6px",
    border: "none",
    backgroundColor: "#942727ff",
    color: "#fff",
    cursor: "pointer",
    justifyContent:"right",
  },
  signOut: {
    fontWeight: "bold",
    color: "#2563eb",
    paddingLeft:"20px",
    paddingRight:"32%",
    width:"200px",
    alignItems:"center"
  },
  createButton: {
    padding: "8px 14px",
    borderRadius: "6px",
    border: "none",
    backgroundColor: "#4f46e5",
    color: "#fff",
    cursor: "pointer",
  },
  row: {
    display: "grid",
    gridTemplateColumns: "0.3fr 0.1fr 0.1fr",
    gap: "50px",
    alignItems: "center",
    padding: "10px 14px",
    borderRadius: "8px",
    justifyContent: 'center',
  },
};