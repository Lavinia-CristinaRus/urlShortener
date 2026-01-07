import { useEffect, useState } from "react";
import api from "../utils/api";
import { useNavigate } from "react-router-dom";

export default function MyUrls() {
  const [urls, setUrls] = useState([]);
  const navigate = useNavigate();
  const [isVisible, setVisible] = useState(false);
  const [url, setUrl] = useState("");
  const [error, setError] = useState(null);
  const [message, setMessage] = useState(null);

  useEffect(() => {
    api.get("/api/urls")
      .then(res => setUrls(res.data))
      .catch(err => console.error(err));
  }, [isVisible]);

  const submit = async () => {
    setError(null);
    setMessage(null);

    try {
      const response = await api.post("/api/generateUrl", {
        url
      });

      setMessage(response.data.message || "Generation done");
      setVisible(false)
    } catch (err) {
      if (err.response) {
        setError(err.response.data.error || "Generation failed");
      } else {
        setError("Cannot connect to server");
      }
    }
  };

  const signOut = () => {
    localStorage.removeItem("token");
    navigate("/signin");
  }

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <div onClick={signOut} style={styles.signOut}>
          <p>sign out</p>
        </div>
        <h2 style={styles.h2}>My Short Links</h2>
        <button
          style={styles.createButton}
          onClick={() => setVisible(true)}
        >
          + Create new short link
        </button>
      </div>

      {urls.length === 0 && (
        <p style={styles.empty}>You don’t have any short links yet.</p>
      )}

      {isVisible && (<div style={styles.generateUrl}>
        <input style={{width:"20%", marginRight:"150px", marginLeft:"120px", height:"25px"}} placeholder="Url" onChange={e => setUrl(e.target.value)} />
        <button
          style={styles.createButton}
          onClick={submit}
        >
          Generate short link
        </button>
        <button
          style={styles.cancelButton}
          onClick={() => setVisible(false)}
        >
          Cancel
        </button>
      </div>)}

      {error && <p style={{ color: "red" }}>{error}</p>}
      {message && <p style={{ color: "green" }}>{message}</p>}

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
              onClick={() => navigate(`/customizeurl/${url.Idurl}`)}
            >
              Customize
            </button>
            <button
              style={styles.edit}
              onClick={() => navigate(`/analytics/${url.Idurl}`)}
            >
              Analytics
            </button>
            <button
              style={styles.edit}
              onClick={() => navigate(`/generateQrCode/${url.Short_url}`)}
            >
              Generate QR
            </button>
            {/* /generateQrCode/:shortUrl */}
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
    width: "100%",
    display: "flex",
    alignItems: "center",
    marginBottom: "20px",
  },
  generateUrl: {
    width: "90%",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    marginBottom: "50px",
    marginTop: "50px",
    gap: "50px",
  },
  h2: {
    width: "30%",
    marginRight: "215px",
  },
  createButton: {
    padding: "8px 14px",
    borderRadius: "6px",
    border: "none",
    backgroundColor: "#4f46e5",
    color: "#fff",
    cursor: "pointer",
  },
  list: {
    display: "flex",
    flexDirection: "column",
    gap: "10px",
  },
  row: {
    display: "grid",
    gridTemplateColumns: "0.6fr 0.22fr 0.22fr auto auto auto",
    gap: "10px",
    alignItems: "center",
    padding: "8px 30px",
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
  }
};