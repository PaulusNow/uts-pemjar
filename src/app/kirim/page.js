"use client";

import { useState } from "react";
import NavBar from "../components/navbar";

// WebSocket
const socket = new WebSocket("ws://localhost:8080/ws");

export default function KirimForm() {
  const [sender, setSender] = useState("");
  const [amount, setAmount] = useState("");
  const [message, setMessage] = useState("");
  const [showToast, setShowToast] = useState(false); // state untuk kontrol toast

  const formatRupiah = (value) => {
    const numericValue = value.replace(/\D/g, "");
    return new Intl.NumberFormat("id-ID").format(numericValue);
  };

  const handleAmountChange = (e) => {
    const value = e.target.value;
    if (/^[0-9]*$/.test(value.replace(/[^0-9]/g, ""))) {
      const numericValue = value.replace(/\D/g, "");
      setAmount(formatRupiah(numericValue));
    }
  };

  const sendMessage = () => {
    const data = { sender, amount, message };
    socket.send(JSON.stringify(data)); // Kirim data melalui WebSocket
    setMessage(""); // Kosongkan pesan setelah kirim
    setSender(""); // Kosongkan nama pengirim
    setAmount(""); // Kosongkan jumlah uang
    setShowToast(true); // Menampilkan toast

    // Menghilangkan toast setelah 5 detik
    setTimeout(() => {
      setShowToast(false);
    }, 5000);
  };

  return (
    <>
      <NavBar />
      <div className="container mt-4 d-flex justify-content-center align-items-center" style={{ minHeight: '100vh' }}>
        <div className="row w-100 justify-content-center">
          <div className="col-8">
            <div className="p-3 rounded border">
              <div className="form-group mb-3">
                <label>Nama Pengirim:</label>
                <input
                  type="text"
                  className="form-control"
                  placeholder="Nama Pengirim"
                  value={sender}
                  onChange={(e) => setSender(e.target.value)}
                />
              </div>

              <div className="form-group mb-3">
                <label>Jumlah Uang:</label>
                <input
                  type="text"
                  className="form-control"
                  placeholder="Jumlah Uang"
                  value={amount}
                  onChange={handleAmountChange}
                />
              </div>

              <div className="form-group mb-3">
                <label>Pesan:</label>
                <textarea
                  className="form-control"
                  placeholder="Pesan"
                  value={message}
                  onChange={(e) => setMessage(e.target.value)}
                />
              </div>

              <button
                type="button"
                className="btn btn-primary"
                onClick={sendMessage}
                id="liveToastBtn"
              >
                Kirim
              </button>

              {/* Toast */}
              {showToast && (
                <div className="toast-container position-fixed top-0 end-0 p-3">
                  <div
                    id="liveToast"
                    className="toast show"
                    role="alert"
                    aria-live="assertive"
                    aria-atomic="true"
                  >
                    <div className="toast-header">
                      <strong className="me-auto">Info</strong>
                      <small>Just now</small>
                      <button
                        type="button"
                        className="btn-close"
                        data-bs-dismiss="toast"
                        aria-label="Close"
                      ></button>
                    </div>
                    <div className="toast-body">
                      Pesan berhasil dikirim!
                    </div>
                  </div>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </>
  );
}