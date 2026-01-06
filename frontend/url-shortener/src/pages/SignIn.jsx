import { useState } from "react";
import api from "../utils/api";

export default function SignIn() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const submit = async () => {
    await api.post("/api/signin", { email, password });
    alert("Response from backend to be added here");
  };

  return (
    <>
      <h2>Sign In</h2>
      <input placeholder="Email" onChange={e => setEmail(e.target.value)} />
      <input type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
      <button onClick={submit}>Sign In</button>
    </>
  );
}