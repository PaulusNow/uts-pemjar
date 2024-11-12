"use client";

import Link from 'next/link';
import NavBar from './components/navbar';

export default function Home() {
  return (
    <>
      <NavBar />
      <div className="container text-center">
        <h1 className="mt-5">Selamat Datang di Sawerkuy</h1>
        <div className="d-flex justify-content-center mt-4">
          <Link href="/terima">
            <button className="btn btn-primary mx-2">Klik disini sebagai Penerima</button>
          </Link>
          <Link href="/kirim">
            <button className="btn btn-secondary mx-2">Klik disini sebagai Pengirim</button>
          </Link>
        </div>
      </div>
    </>
  );
}