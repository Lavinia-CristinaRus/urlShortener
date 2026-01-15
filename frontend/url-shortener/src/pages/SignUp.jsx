import { useState } from "react";
import api from "../utils/api";
import { useNavigate } from "react-router-dom";

export default function SignUp() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState(null);
  const [message, setMessage] = useState(null);
  const navigate = useNavigate();

  const submit = async () => {
    setError(null);
    setMessage(null);

    try {
      const response = await api.post("/api/signup", {
        email,
        password,
      });
      localStorage.setItem("token", response.data.token);

      setMessage(response.data.message || "New account created");
      navigate("/urls");
    } catch (err) {
      if (err.response) {
        setError(err.response.data.error || "New account could not be created");
      } else {
        setError("Cannot connect to server");
      }
    }
  };

  const signIn = () => {
    navigate("/signin");
  }

  return (
    <>
      <div  style={styles.container}>
        <h2>Sign Up</h2>
        <input placeholder="Email" onChange={e => setEmail(e.target.value)} />
        <input type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
        <button style={{ marginTop: "20px" }} onClick={submit}>Sign Up</button>
        {error && <p style={{ color: "red" }}>{error}</p>}
        {message && <p style={{ color: "green" }}>{message}</p>}
        <div onClick={signIn} style={styles.signIn}>
          <p> or sign in</p>
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
  signIn: {
    fontWeight: "bold",
    color: "#2563eb",
    width:"100%",
    alignItems:"center",
    justifyItems:"center",
    display: "grid"
  }
}