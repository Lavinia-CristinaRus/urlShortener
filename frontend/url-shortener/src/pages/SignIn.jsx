import { useState } from "react";
import api from "../utils/api";
import { useNavigate } from "react-router-dom";

export default function SignIn() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState(null);
  const [message, setMessage] = useState(null);
  const navigate = useNavigate();

  const submit = async () => {
    setError(null);
    setMessage(null);

    try {
      const response = await api.post("/api/signin", {
        email,
        password,
      });
      localStorage.setItem("token", response.data.token);

      setMessage(response.data.message || "Login done");
      navigate("/urls");
    } catch (err) {
      if (err.response) {
        setError(err.response.data.error || "Login failed");
      } else {
        setError("Cannot connect to server");
      }
    }
  };

  const signUp = () => {
    navigate("/signup");
  }

  return (
    <>
      <div  style={styles.container}>
        <h2>Sign In</h2>
        <input placeholder="Email" onChange={e => setEmail(e.target.value)} />
        <input type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
        <button style={{ marginTop: "20px" }} onClick={submit}>Sign In</button>
        {error && <p style={{ color: "red" }}>{error}</p>}
        {message && <p style={{ color: "green" }}>{message}</p>}
        <div onClick={signUp} style={styles.signUp}>
          <p> or sign up</p>
        </div>
      </div>
    </>
  );
}

const styles = {
  container: {
    width:"20%",
    height:"90%",
    marginLeft:"40%",
    marginTop:"5%",
    display: "flex",
    flexDirection: "column",
    gap: "20px",
  },
  signUp: {
    fontWeight: "bold",
    color: "#2563eb",
    width:"100%",
    alignItems:"center",
    justifyItems:"center",
    display: "grid"
  }
}
