import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import SignIn from "../pages/SignIn";
import SignUp from "../pages/SignUp";
import UrlList from "../pages/UrlList";
import CustomizeUrl from "../pages/CustomizeUrl";
import QrCode from "../pages/QrCode";
import Analytics from "../pages/Analytics"

function PrivateRoute({ children }) {
  const token = localStorage.getItem("token");

  if (!token) {
    return <Navigate to="/signin" />;
  }

  return children;
}

export default function AppRouter() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navigate to="/signin" />} />
        <Route path="/signin" element={<SignIn />}/>
        <Route path="/signup" element={<SignUp />}/>
        <Route path="/urls" element={<PrivateRoute><UrlList /></PrivateRoute>}/>
        <Route path="/customizeurl/:id" element={<PrivateRoute><CustomizeUrl /></PrivateRoute>}/>
        <Route path="/generateQrCode/:shortUrl" element={<PrivateRoute><QrCode /></PrivateRoute>}/>
        <Route path="/analytics/:id" element={<PrivateRoute><Analytics /></PrivateRoute>}/>
      </Routes>
    </BrowserRouter>
  );
}