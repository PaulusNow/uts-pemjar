import Link from "next/link";

const Navbar = () => {
  return (
    <div>
      {/* Navbar with logo and text */}
      <nav className="navbar navbar-light bg-primary mb-4 justify-content-center">
        <Link href="/" className="navbar-brand text-white d-flex align-items-center">
          Sawerkuy
        </Link>
      </nav>
    </div>
  );
};

export default Navbar;