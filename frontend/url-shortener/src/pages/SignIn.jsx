import { useState } from "react";
import api from "../utils/api";

export default function SignIn() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState(null);
  const [message, setMessage] = useState(null);

  const submit = async () => {
    setError(null);
    setMessage(null);

    try {
      const response = await api.post("/api/signin", {
        email,
        password,
      });

      setMessage(response.message || "Login successful");
    } catch (err) {
      if (err.response) {
        setError(err.response.data.error || "Login failed");
      } else {
        setError("Cannot connect to server");
      }
    }
  };

  return (
    <>
      <h2>Sign In</h2>
      <input placeholder="Email" onChange={e => setEmail(e.target.value)} />
      <input type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
      <button onClick={submit}>Sign In</button>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {message && <p style={{ color: "green" }}>{message}</p>}
    </>
  );
}