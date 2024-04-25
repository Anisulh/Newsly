import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <nav className="navbar">
      <Link to="/" className="nav-link logo">Headlines</Link>
      <Link to="/categories" className="nav-link">
        Categories
      </Link>
      <Link to="/about" className="nav-link">
        About
      </Link>
      <Link to="/login" className="nav-link">
        Login
      </Link>
      <Link to="/sign-up" className="nav-link">
        Sign Up
      </Link>
    </nav>
  );
};

export default Navbar;
