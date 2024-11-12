"use client";

import { useEffect, useState } from "react";
import NavBar from "../components/navbar";

const socket = new WebSocket("ws://localhost:8080/ws"); // Gunakan WebSocket biasa

export default function Client1() {
  const [messages, setMessages] = useState([]);

  useEffect(() => {
    socket.onmessage = (event) => {
      const data = JSON.parse(event.data); // Parse data yang diterima dari server
      setMessages((prevMessages) => [...prevMessages, data]);
    };
  }, []);

  return (
    <>
      <NavBar />
      <div className="container mt-4">
        <h1 className="text-center">Pesan Donasi</h1>
        <div className="row">
          {/* Menampilkan maksimal 3 kartu di desktop */}
          {messages.slice(0, 3).map((msg, index) => (
            <div key={index} className="col-12 col-md-6 col-lg-4 mb-4">
              <div
                className="card"
                style={{
                  backgroundColor: "#f1c40f", // Latar belakang kuning
                  color: "black", // Teks berwarna hitam
                  borderRadius: "8px", // Menambahkan border radius agar sudutnya lebih lembut
                  height: "200px", // Menyesuaikan tinggi kartu
                  display: "flex",
                  flexDirection: "column", // Mengatur konten agar terpusat secara vertikal
                  justifyContent: "center",
                  textAlign: "center", // Menyusun teks di tengah
                }}
              >
                <div className="card-body">
                  <h5 className="card-title">{msg.sender} Memberikan Donasi Sebesar Rp {msg.amount}</h5>
                  <p className="card-text">{msg.message}</p>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </>
  );
}