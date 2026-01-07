import { useParams } from "react-router-dom";
import { useEffect, useState} from "react";

export default function QrCode() {
  const { shortUrl } = useParams();
  const [qrUrl, setQrUrl] = useState(null);
  useEffect(() => {
    const token = localStorage.getItem("token");
    fetch(`http://localhost:4000/api/generateQrCode/${shortUrl}`, {
        headers: {
        Authorization: `Bearer ${token}`,
        },
    })
        .then(res => res.blob())
        .then(blob => {
        setQrUrl(URL.createObjectURL(blob));
        });
    }, [shortUrl]);

  return (
    <div style={{width:"100%", height:"100%", alignItems:"center", justifyItems:"center", display: "flex",  flexDirection: "column",}}>
      <h2>QR Code</h2>
      <img
        src={qrUrl}
        alt="QR Code"
      />
    </div>
  );
}