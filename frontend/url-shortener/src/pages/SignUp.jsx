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

  return (
    <>
      <h2>Sign Up</h2>
      <input placeholder="Email" onChange={e => setEmail(e.target.value)} />
      <input type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
      <button onClick={submit}>Sign Up</button>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {message && <p style={{ color: "green" }}>{message}</p>}
    </>
  );
}